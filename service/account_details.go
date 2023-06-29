package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type AccountResponse struct {
	AccountStatus string   `json:"accountStatus"`
	DpType        string   `json:"dpType"`
	AccountID     string   `json:"accountId"`
	SBrokerName   string   `json:"sBrokerName"`
	Product       []string `json:"product"`
	AccountName   string   `json:"accountName"`
	CellAddr      string   `json:"cellAddr"`
	EmailAddr     string   `json:"emailAddr"`
	ExchEnabled   string   `json:"exchEnabled"`
}

func (a AliceBlue) GetAccountDetails() (AccountResponse, error) {
	var err error

	client := resty.New()

	var ar AccountResponse

	var rsp *resty.Response
	if rsp, err = client.R().SetAuthToken(a.token).
		Get(a.endpoints.GetAccountDetails); err != nil {
		return AccountResponse{}, err
	}

	if rsp.StatusCode() != http.StatusOK {
		return AccountResponse{}, fmt.Errorf("GetAccountDetails code: %d body: %s", rsp.StatusCode(), rsp.String())
	}

	if err = json.Unmarshal(rsp.Body(), &ar); err != nil {
		return AccountResponse{}, err
	}

	return ar, err
}
