package endpoints

import "fmt"

type Endpoints struct {
	GetEncryptionKey  string
	GetUserSID        string
	GetAccountDetails string
	GetStockHistory   string
	GetLimits         string
	Base              string
	ContractMaster   string
}

func New(base string) Endpoints {
	if base == "" {
		base = "https://ant.aliceblueonline.com/rest/AliceBlueAPIService/api"
	}
	return Endpoints{
		GetEncryptionKey:  fmt.Sprintf("%s/customer/getAPIEncpkey", base),
		GetUserSID:        fmt.Sprintf("%s/customer/getUserSID", base),
		GetAccountDetails: fmt.Sprintf("%s/customer/accountDetails", base),
		GetStockHistory:   fmt.Sprintf("%s/chart/history", base),
		GetLimits:         fmt.Sprintf("%s/limits/getRmsLimits", base),
		ContractMaster:   "https://v2api.aliceblueonline.com/restpy/contract_master?exch=NSE",
	}
}
