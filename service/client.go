package service

import (
	"errors"
	"os"

	"github.com/imyashkale/go-aliceblue-sdk/endpoints"

	"github.com/joho/godotenv"
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

func NewFromEnv() (*AliceBlue, error) {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		return &AliceBlue{}, err
	}

	baseURL := os.Getenv("BASE_URL")
	return &AliceBlue{
		apiKey:    os.Getenv("API_KEY"),
		clientId:  os.Getenv("CLIENT_ID"),
		baseURL:   baseURL,
		endpoints: endpoints.New(baseURL),
	}, nil
}

func (a *AliceBlue) Connect() error {

	if a.apiKey == "" || a.clientId == "" {
		return errors.New("apikey or client id not provided")
	}

	var err error
	var enc EncryptionResponse
	if enc, err = a.getAPIEncKey(); err != nil {
		return err
	}
	a.encKey = enc.EncKey

	var session SessionResponse
	if session, err = a.getUserSID(); err != nil {
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
