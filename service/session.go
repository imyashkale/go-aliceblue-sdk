package service

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type SessionResponse struct {
	Stat      string `json:"stat"`
	SessionID string `json:"sessionID"`
	Token     string `json:"token"`
}

func (a *AliceBlue) getUserSID() (SessionResponse, error) {
	var session SessionResponse
	var err error

	client := resty.New()

	var ud string
	if ud, err = a.generateUserData(); err != nil {
		return SessionResponse{}, err
	}

	rb := map[string]any{
		"userId":   a.clientId,
		"userData": ud,
	}

	var rsp *resty.Response
	if rsp, err = client.R().EnableTrace().SetBody(rb).Post(a.endpoints.GetUserSID); err != nil {
		return SessionResponse{}, err
	}

	if rsp.StatusCode() != http.StatusOK {
		return SessionResponse{}, fmt.Errorf("GetUserSID code: %d body: %s", rsp.StatusCode(), rsp.String())
	}

	if err = json.Unmarshal(rsp.Body(), &session); err != nil {
		return SessionResponse{}, err
	}

	if session.SessionID == "" {
		return SessionResponse{}, fmt.Errorf("GetAPIEncKey session id not found in response code: %d body: %s", rsp.StatusCode(), rsp.String())
	}

	session.Token = fmt.Sprintf("%s %s", a.clientId, session.SessionID)
	return session, err
}

func (a *AliceBlue) generateUserData() (string, error) {
	var err error

	// This request_token, along with a checkSum (USER ID + API_KEY + Encryption Key)
	// is posted to the token API to obtain an access_token,
	// which is then used for signing all subsequent requests.

	dt := strings.Builder{}
	dt.WriteString(a.clientId)
	dt.WriteString(a.apiKey)
	dt.WriteString(a.encKey)

	t := dt.String()

	h := sha256.New()
	h.Write([]byte(t))

	b := fmt.Sprintf("%x", h.Sum(nil))

	return string(b), err
}
