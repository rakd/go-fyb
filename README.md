go-fyb
==========

go-fyb is an implementation of the fybsg.com & fybse.se API in Golang.

Based off of https://github.com/rakd/go-fyb/

## Import
```
import "github.com/rakd/go-fyb"
```

## Usage
~~~ go
package main

import (
	"fmt"
	"github.com/rakd/go-fyb"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// fyb client
	client := fyb.New(fyb.APIBaseURLForSGD, API_KEY, API_SECRET)

	// Get tickers
	tickers, err := client.GetTickers()
	fmt.Println(err, tickers)
}
~~~


## Stay tuned

- [Follow me on Twitter](https://twitter.com/kaz_lavender)

## Donate

- BTC: 1Ah8sarQ4w9FnsCs8LoG6JuYiFHmrAAy6F
