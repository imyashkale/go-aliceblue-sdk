package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type EncryptionResponse struct {
	UserID    string      `json:"userId"`
	UserData  interface{} `json:"userData"`
	EncKey    string      `json:"encKey"`
	Apikey    interface{} `json:"apikey"`
	Stat      string      `json:"stat"`
	Emsg      string      `json:"emsg"`
	LoginType interface{} `json:"loginType"`
	Version   interface{} `json:"version"`
	FcmToken  interface{} `json:"fcmToken"`
	Imei      interface{} `json:"imei"`
	Login     bool        `json:"login"`
}

func (a *AliceBlue) getAPIEncKey() (EncryptionResponse, error) {
	var err error
	var encryption EncryptionResponse

	client := resty.New()

	rb := map[string]any{
		"userId": a.clientId,
	}

	var rsp *resty.Response
	if rsp, err = client.R().EnableTrace().SetBody(rb).Post(a.endpoints.GetEncryptionKey); err != nil {
		return EncryptionResponse{}, err
	}

	if rsp.StatusCode() != http.StatusOK {
		return EncryptionResponse{}, fmt.Errorf("GetAPIEncKey code: %d body: %s", rsp.StatusCode(), rsp.String())
	}

	if err = json.Unmarshal(rsp.Body(), &encryption); err != nil {
		return EncryptionResponse{}, fmt.Errorf("GetAPIEncKey unmarshaling failed body: %s", rsp.String())
	}

	if encryption.EncKey == "" {
		return EncryptionResponse{}, fmt.Errorf("GetAPIEncKey response doesn't contain encryption key body: %s", rsp.String())
	}

	return encryption, err
}
