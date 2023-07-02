package service

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func (a *AliceBlue) retryUnAuthorized(r *resty.Response, err error) bool {
	if r.StatusCode() == http.StatusUnauthorized {
		a.Connect()
		fmt.Println("request retrying due to unauthorized error")
		return true
	}
	return false
}
