package service_test

import (
	"github.com/imyashkale/go-aliceblue-sdk/service"
	"testing"
)

var cf = service.Config{
	BaseURL:  "https://ant.aliceblueonline.com/rest/AliceBlueAPIService/api",
	APIKey:   "il70vdNS6FuYlNJQF4bP3LcCgXS8bepJTv0vvVttuLdl73zQ3lGmjBzAuIoTvU9qqGg0dukXglS5NhvaX1PvmT4rIsVerON2sHsSRaFgrdqcVVImHfqjucicdsG1KsJI",
	ClientId: "551165",
}

func TestGetEncryptionKey(t *testing.T) {
	ab := service.NewFromConfig(cf)

	var err error
	var enc service.EncryptionResponse
	if enc, err = ab.GetAPIEncKey(); err != nil {
		t.Error(err)
	}

	if enc.EncKey == "" {
		t.Error("encryption key not returned")
	}

}
