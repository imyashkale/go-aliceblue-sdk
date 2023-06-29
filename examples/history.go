package main

import (
	"fmt"
	"time"

	"github.com/imyashkale/go-aliceblue-sdk/service"
)

func main() {
	
	var err error
	var ab *service.AliceBlue

	if ab, err = service.NewFromEnv(); err != nil {
		panic(err)
	}

	if err = ab.Connect(); err != nil {
		panic(err)
	}

	ITC_TOKEN := "1660"

	params := service.StockHistoryInput{
		Token: ITC_TOKEN,
		From:  time.Now().AddDate(0, 0, -1),
		To:    time.Now(),
	}

	var sh service.StockHistoryResponse
	if sh, err = ab.GetStockHistory(params); err != nil {
		panic(err)
	}

	fmt.Println("no of stocks resolutions retrived ", len(sh.Result))

}
