package service

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type NSEContractResponse struct {
	Contracts []Contract `json:"NSE"`
}

type BSEContractResponse struct {
	Contracts []Contract `json:"BSE"`
}

type Contract struct {
	Symbol string `json:"symbol"`
	Token  string `json:"token"`
}

func (a AliceBlue) GetNSEContracts() (NSEContractResponse, error) {
	var err error
	var contracts NSEContractResponse

	client := resty.New()

	var rsp *resty.Response
	if rsp, err = client.R().EnableTrace().SetQueryParam("exch", "NSE").Get(a.endpoints.ContractMaster); err != nil {
		return contracts, err
	}

	if err = json.Unmarshal(rsp.Body(), &contracts); err != nil {
		return contracts, err
	}
	return contracts, err
}

func (a AliceBlue) GetBSEContracts() (BSEContractResponse, error) {
	var err error
	var contracts BSEContractResponse

	client := resty.New()

	var rsp *resty.Response
	if rsp, err = client.R().EnableTrace().SetQueryParam("exch", "BSE").Get(a.endpoints.ContractMaster); err != nil {
		return contracts, err
	}

	if err = json.Unmarshal(rsp.Body(), &contracts); err != nil {
		return contracts, err
	}
	return contracts, err
}
