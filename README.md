# Unofficial Go AliceBlue SDK
This project, is a personal endeavor developed to address a specific personal use case. It is important to note that this project is not an official SDK or a project sponsored by any company or organization. As such, it may not encompass all the features provided by the official API. [ Official API Docs](https://v2api.aliceblueonline.com/introduction)

# Disclaimer: Use this project at your own risk.

### Installation
```
go get github.com/imyashkale/go-aliceblue-sdk@v1.0.1
```

## Example

# Fetch Historical Data 
````
package main

import (
	"fmt"
	"github.com/imyashkale/go-aliceblue-sdk/service"
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

````
