package main

import (
	"fmt"
	"imyashkale/go-aliceblue-sdk/service"
	"time"
)

func main() {

	var err error
	var svc *service.AliceBlue
	if svc, err = service.NewFromEnv(); err != nil {
		panic(err)
	}

	if err = svc.Connect(); err != nil {
		panic(err)
	}

	var h service.StockHistoryResponse
	if h, err = svc.GetStockHistory(service.StockHistoryInput{
		Token: "1594",
		From:  time.Now().AddDate(0, 0, -5),
		To:    time.Now(),
	}); err != nil {
		panic(err)
	}

	for _, st := range h.Result {
		fmt.Println(st.Time, st.High)
	}
}
