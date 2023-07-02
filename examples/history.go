package main

import (
	"fmt"
	"os"
	"time"

	"github.com/imyashkale/go-aliceblue-sdk/service"
)

func main() {

	var err error

	cf := service.Config{
		APIKey:   os.Getenv("API_KEY"),
		ClientId: os.Getenv("CLIENT_ID"),
	}

	ab := service.NewFromConfig(cf)

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
