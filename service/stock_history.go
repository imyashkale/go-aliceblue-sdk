package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/imyashkale/go-aliceblue-sdk/options"

	"github.com/go-resty/resty/v2"
)

type Stock struct {
	Volume float64 `json:"volume"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Time   string  `json:"time"`
	Close  float64 `json:"close"`
	Open   float64 `json:"open"`
}

type StockHistoryResponse struct {
	Stat    string      `json:"stat"`
	Result  []Stock     `json:"result"`
	Message interface{} `json:"message"`
}

type StockHistoryInput struct {
	Token      string    `json:"token"`
	Resolution string    `json:"resolution"`
	From       time.Time `json:"from"`
	To         time.Time `json:"to"`
	Exchange   string    `json:"exchange"`
}

func (a *AliceBlue) GetStockHistory(params StockHistoryInput) (StockHistoryResponse, error) {
	var err error
	var sh StockHistoryResponse

	client := resty.New()
	client.SetRetryCount(3)

	if params.Resolution == "" {
		params.Resolution = options.Day
	}

	if params.Exchange == "" {
		params.Exchange = options.NSE
	}

	pm := map[string]any{
		"token":      params.Token,
		"resolution": params.Resolution,
		"from":       params.From.UnixMilli(),
		"to":         params.To.UnixMilli(),
		"exchange":   params.Exchange,
	}

	var rsp *resty.Response
	if rsp, err = client.R().EnableTrace().SetAuthToken(a.token).AddRetryCondition(a.retryUnAuthorized).SetBody(pm).Post(a.endpoints.GetStockHistory); err != nil {
		return StockHistoryResponse{}, err
	}

	if rsp.StatusCode() != http.StatusOK {
		return StockHistoryResponse{}, fmt.Errorf("GetStockHistory code: %d body: %s", rsp.StatusCode(), rsp.String())
	}

	if err = json.Unmarshal(rsp.Body(), &sh); err != nil {
		return StockHistoryResponse{}, err
	}

	return sh, err
}
