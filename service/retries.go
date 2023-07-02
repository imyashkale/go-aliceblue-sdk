package service

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

func (a *AliceBlue) retryUnAuthorized(r *resty.Response, err error) bool {
	return r.StatusCode() == http.StatusUnauthorized
}
