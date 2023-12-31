package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Limit struct {
	Symbol                 string `json:"symbol"`
	CncMarginUsed          string `json:"cncMarginUsed"`
	Spanmargin             string `json:"spanmargin"`
	BranchAdhoc            string `json:"branchAdhoc"`
	AdhocMargin            string `json:"adhocMargin"`
	Payoutamount           string `json:"payoutamount"`
	CdsSpreadBenefit       string `json:"cdsSpreadBenefit"`
	Adhocscripmargin       string `json:"adhocscripmargin"`
	Exposuremargin         string `json:"exposuremargin"`
	Scripbasketmargin      string `json:"scripbasketmargin"`
	Credits                string `json:"credits"`
	Segment                string `json:"segment"`
	Net                    string `json:"net"`
	Turnover               string `json:"turnover"`
	Grossexposurevalue     string `json:"grossexposurevalue"`
	MfssAmountUsed         string `json:"mfssAmountUsed"`
	RealizedMtomPrsnt      string `json:"realizedMtomPrsnt"`
	Product                string `json:"product"`
	Stat                   string `json:"stat"`
	CncSellCrditPrsnt      string `json:"cncSellCrditPrsnt"`
	Debits                 string `json:"debits"`
	Varmargin              string `json:"varmargin"`
	Multiplier             string `json:"multiplier"`
	Elm                    string `json:"elm"`
	Mfamount               string `json:"mfamount"`
	Cashmarginavailable    string `json:"cashmarginavailable"`
	BrokeragePrsnt         string `json:"brokeragePrsnt"`
	CncRealizedMtomPrsnt   string `json:"cncRealizedMtomPrsnt"`
	NotionalCash           string `json:"notionalCash"`
	Directcollateralvalue  string `json:"directcollateralvalue"`
	CncBrokeragePrsnt      string `json:"cncBrokeragePrsnt"`
	Valueindelivery        string `json:"valueindelivery"`
	NfoSpreadBenefit       string `json:"nfoSpreadBenefit"`
	Losslimit              string `json:"losslimit"`
	Subtotal               string `json:"subtotal"`
	RmsPayInAmnt           string `json:"rmsPayInAmnt"`
	UnrealizedMtomPrsnt    string `json:"unrealizedMtomPrsnt"`
	CoverOrderMarginPrsnt  string `json:"coverOrderMarginPrsnt"`
	Exchange               string `json:"exchange"`
	Category               string `json:"category"`
	Collateralvalue        string `json:"collateralvalue"`
	RmsIpoAmnt             string `json:"rmsIpoAmnt"`
	CncUnrealizedMtomPrsnt string `json:"cncUnrealizedMtomPrsnt"`
	PremiumPrsnt           string `json:"premiumPrsnt"`
}

func (a AliceBlue) GetLimits() ([]Limit, error) {
	var err error

	client := resty.New()

	var limits []Limit

	var rsp *resty.Response
	if rsp, err = client.R().SetAuthToken(a.token).
		Get(a.endpoints.GetLimits); err != nil {
		return limits, err
	}

	if rsp.StatusCode() != http.StatusOK {
		return limits, fmt.Errorf("GetLimits code: %d body: %s", rsp.StatusCode(), rsp.String())
	}

	if err = json.Unmarshal(rsp.Body(), &limits); err != nil {
		return limits, err
	}

	return limits, err
}
