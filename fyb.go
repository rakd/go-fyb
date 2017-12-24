package fyb

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

const (
	// APIBaseURLForSGD fybsg API endpoint
	APIBaseURLForSGD = "https://fybsg.com/api/SGD"
	// APIBaseURLForSEK fybse API endpoint
	APIBaseURLForSEK = "https://fybse.se/api/SEK"
	// APIBaseURLForTest fysg Test API endpoint
	APIBaseURLForTest = "https://private-a9161-fyb.apiary-proxy.com/api/SGD"
)

func init() {
	log.SetFlags(log.Lshortfile)

}

// New returns an instantiated fyb struct
func New(apibaseurl, apiKey, apiSecret string) *Fyb {
	client := NewClient(apibaseurl, apiKey, apiSecret)
	return &Fyb{client}
}

// NewWithCustomTimeout returns an instantiated fyb struct with custom timeout
func NewWithCustomTimeout(apibaseurl, apiKey, apiSecret string, timeout time.Duration) *Fyb {
	client := NewClientWithCustomTimeout(apibaseurl, apiKey, apiSecret, timeout)
	return &Fyb{client}
}

// Fyb represent a fyb client
type Fyb struct {
	client *Client
}

// GetOrderBook ..
func (b *Fyb) GetOrderBook() (orderbook OrderBook, err error) {
	r, err := b.client.do("GET", "orderbook.json", nil, false)
	if err != nil {
		//log.Print(err)
		return
	}
	js, err := simplejson.NewJson(r)
	if err != nil {
		return
	}

	//log.Print(js)
	asks, err := js.GetPath("asks").Array()
	if err != nil {
		return
	}
	for _, ask := range asks {

		item, e := NewPriceAmountFromInterface(ask)
		if e != nil {
			continue
		}
		orderbook.Asks = append(orderbook.Asks, item)

	}
	bids, err := js.GetPath("bids").Array()
	if err != nil {
		return
	}

	for _, bid := range bids {
		item, err := NewPriceAmountFromInterface(bid)
		if err != nil {
			continue
		}
		orderbook.Bids = append(orderbook.Bids, item)
	}

	return
}

// GetTicker ...
func (b *Fyb) GetTicker() (ticker Ticker, err error) {
	r, err := b.client.do("GET", "tickerdetailed.json", nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &ticker); err != nil {
		return
	}
	return
}

// GetTradeHistory ...
// tid ()= Trade ID) to begin trade history from.
// You should cache trade history and query only new trades by passing in last known trade id
func (b *Fyb) GetTradeHistory(tid int64) (trades Trades, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("trades.json?since=%d", tid), nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &trades); err != nil {
		return
	}
	return
}

// APITokenTest private API
func (b *Fyb) APITokenTest() (res TestResponse, err error) {
	r, err := b.client.do("POST", fmt.Sprintf("test"), nil, true)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &res); err != nil {
		kperr := KeyPermissionErrorResponse{}
		if err = json.Unmarshal(r, &kperr); err == nil {
			res.Error = 1
			res.Msg = kperr.Error
			err = fmt.Errorf(kperr.Error)
			return
		}
		log.Printf("body=%s", string(r))
		res.Error = 1
		err = fmt.Errorf("%s: body=%s", err.Error(), string(r))
		res.Msg = err.Error()
		return
	}
	return
}

// GetAccountInfo ..
func (b *Fyb) GetAccountInfo() (res AccountInfoResponse, err error) {
	r, err := b.client.do("POST", fmt.Sprintf("getaccinfo"), nil, true)
	if err != nil {
		return
	}
	//log.Printf("r:%s", string(r))

	if err = json.Unmarshal(r, &res); err != nil {
		kperr := KeyPermissionErrorResponse{}
		if err = json.Unmarshal(r, &kperr); err == nil {
			res.Error = 1
			res.Msg = kperr.Error
			err = fmt.Errorf(kperr.Error)
			return
		}

		log.Printf("body=%s", string(r))
		res.Error = 1
		err = fmt.Errorf("%s: body=%s", err.Error(), string(r))
		res.Msg = err.Error()

		return
	}
	// ok
	return
}

// GetPendingOrders ..
func (b *Fyb) GetPendingOrders() (res PendingOrderResponse, err error) {
	r, err := b.client.do("POST", fmt.Sprintf("getpendingorders"), nil, true)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &res); err != nil {
		kperr := KeyPermissionErrorResponse{}
		if err = json.Unmarshal(r, &kperr); err == nil {
			res.Error = 1
			res.Msg = kperr.Error
			err = fmt.Errorf(kperr.Error)
			return
		}
		log.Printf("body=%s", string(r))
		res.Error = 1
		err = fmt.Errorf("%s: body=%s", err.Error(), string(r))
		res.Msg = err.Error()
		return
	}
	return
}

// GetOrderHistory ..
// limit int64, Number of Order History Items to return : Number
func (b *Fyb) GetOrderHistory(limit int64) (res OrderHistoryResponse, err error) {
	payload := map[string]string{}
	payload["limit"] = fmt.Sprintf("%d", limit)
	r, err := b.client.do("POST", fmt.Sprintf("getorderhistory"), payload, true)
	if err != nil {

		return
	}
	//log.Print(string(r))
	if err = json.Unmarshal(r, &res); err != nil {

		kperr := KeyPermissionErrorResponse{}
		if err = json.Unmarshal(r, &kperr); err == nil {
			res.Error = 1
			res.Msg = kperr.Error
			err = fmt.Errorf(kperr.Error)
			return
		}
		log.Printf("body=%s", string(r))
		res.Error = 1
		err = fmt.Errorf("%s: body=%s", err.Error(), string(r))
		res.Msg = err.Error()
		return
	}
	// ok
	return

}

// CancelPendingOrder ..
// orderNo, Ticket Number of Pending Order to cancel. :number
func (b *Fyb) CancelPendingOrder(orderNo int64) (res CancelPendingOrderResponse, err error) {
	payload := map[string]string{}
	payload["orderNo"] = fmt.Sprintf("%d", orderNo)

	r, err := b.client.do("POST", fmt.Sprintf("cancelpendingorder"), payload, true)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &res); err != nil {

		kperr := KeyPermissionErrorResponse{}
		if err = json.Unmarshal(r, &kperr); err == nil {
			res.Error = 1
			res.Msg = kperr.Error
			err = fmt.Errorf(kperr.Error)
			return
		}
		log.Printf("body=%s", string(r))
		res.Error = 1
		err = fmt.Errorf("%s: body=%s", err.Error(), string(r))
		res.Msg = err.Error()
		return
	}

	return
}

// PlaceOrder ..
// qty float64, Quantity of bitcoins, Number
// price , Price to place order at , Number
// type , Whether it is a buy or sell order. Must be either 'B' or 'S' only. ,Char
func (b *Fyb) PlaceOrder(orderType string, price, qty float64) (res PlaceOrderResponse, err error) {

	orderType = strings.ToUpper(orderType)
	payload := map[string]string{}
	if orderType == "SELL" || orderType == "S" {
		payload["type"] = "S"
	} else if orderType == "BUY" || orderType == "B" {
		payload["type"] = "B"
	} else {
		err = fmt.Errorf("orderType must be S or B")
		return
	}
	payload["price"] = fmt.Sprintf("%f", price)
	payload["qty"] = fmt.Sprintf("%f", qty)

	r, err := b.client.do("POST", fmt.Sprintf("placeorder"), payload, true)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &res); err != nil {
		kperr := KeyPermissionErrorResponse{}
		if err = json.Unmarshal(r, &kperr); err == nil {
			res.Error = 1
			res.Msg = kperr.Error
			err = fmt.Errorf(kperr.Error)
			return
		}
		log.Printf("body=%s", string(r))
		res.Error = 1
		err = fmt.Errorf("%s: body=%s", err.Error(), string(r))
		res.Msg = err.Error()
		return
	}
	return
}

// Withdraw ..
// amount Amount of bitcoins/dollars to withdraw: Number
// destination Bitcoin address to withdraw to, leave blank for XFERS: String
// type BTC/XFERS (XFERS only for FYB-SG) // Char
func (b *Fyb) Withdraw(amount float64, destination string, destinationType string) (res WithdrawResponse, err error) {
	destinationType = strings.ToUpper(destinationType)
	payload := map[string]string{}
	payload["destination"] = strings.Trim(destination, "\r\n ")
	payload["amount"] = fmt.Sprintf("%f", amount)

	if destinationType == "BTC" {
		payload["type"] = "BTC"
	} else if destinationType == "XFERS" {
		payload["type"] = "XFERS"
	} else {
		err = fmt.Errorf("destinationType must be BTC or XFERS")
		return
	}

	r, err := b.client.do("POST", fmt.Sprintf("withdraw"), nil, true)
	if err != nil {
		return
	}
	//log.Printf("body=%s", string(r))
	if err = json.Unmarshal(r, &res); err != nil {

		kperr := KeyPermissionErrorResponse{}
		if err = json.Unmarshal(r, &kperr); err == nil {
			res.Error = 1
			res.Msg = kperr.Error
			err = fmt.Errorf(kperr.Error)
			return
		}
		log.Printf("body=%s", string(r))

		res.Error = 1
		err = fmt.Errorf("%s: body=%s", err.Error(), string(r))
		res.Msg = err.Error()

		return
	}
	return
}
