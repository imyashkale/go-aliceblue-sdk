package service

import (
	"imyashkale/go-aliceblue-sdk/endpoints"
)

type AliceBlue struct {
	apiKey    string
	sessionId string
	encKey    string
	baseURL   string
	clientId  string
	token     string
	endpoints endpoints.Endpoints
}

type Config struct {
	APIKey   string
	ClientId string
	BaseURL  string
}

// NewFromConfig
func NewFromConfig(cf Config) *AliceBlue {
	return &AliceBlue{
		baseURL:   cf.BaseURL,
		clientId:  cf.ClientId,
		apiKey:    cf.APIKey,
		endpoints: endpoints.New(cf.BaseURL),
	}
}

func (a *AliceBlue) Connect() error {
	var err error

	var enc EncryptionResponse
	if enc, err = a.GetAPIEncKey(); err != nil {
		return err
	}
	a.encKey = enc.EncKey

	var session SessionResponse
	if session, err = a.GetUserSID(); err != nil {
		return err
	}

	a.sessionId = session.SessionID
	a.token = session.Token
	return err
}

func (a AliceBlue) Session() string {
	return a.sessionId
}

func (a AliceBlue) Token() string {
	return a.token
}
