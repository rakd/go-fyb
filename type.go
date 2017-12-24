package fyb

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// Ticker ..
type Ticker struct {
	Ask  decimal.Decimal `json:"ask"`
	Bid  decimal.Decimal `json:"bid"`
	Last decimal.Decimal `json:"last"`
	Vol  decimal.Decimal `json:"vol"`
}

// Trade ..
type Trade struct {
	Amount decimal.Decimal `json:"amount"`
	Date   int64           `json:"date"`
	Price  decimal.Decimal `json:"price"`
	TID    int64           `json:"tid"`
}

// Trades sortable trade array
type Trades []Trade

// OrderBook ..
type OrderBook struct {
	Asks []PriceAmount `json:"asks"`
	Bids []PriceAmount `json:"bids"`
}

// PriceAmount ..
type PriceAmount struct {
	Price  decimal.Decimal `json:"price"`
	Amount decimal.Decimal `json:"amount"`
}

// NewPriceAmountFromInterface ..
func NewPriceAmountFromInterface(i interface{}) (PriceAmount, error) {
	var item PriceAmount
	s := fmt.Sprintf("%v", i)

	if len(s) < 5 {
		return item, fmt.Errorf("wrong interface for PriceAmount")
	}

	if !strings.HasPrefix(s, "[") {
		return item, fmt.Errorf("wrong interface for PriceAmount")
	}
	if !strings.HasSuffix(s, "]") {
		return item, fmt.Errorf("wrong interface for PriceAmount")
	}

	s = s[1 : len(s)-2]
	decimals := strings.Split(s, " ")
	if len(decimals) != 2 {
		return item, fmt.Errorf("wrong interface for PriceAmount")
	}
	price, err := decimal.NewFromString(decimals[0])
	if err != nil {
		return item, err
	}
	item.Price = price
	amount, err := decimal.NewFromString(decimals[1])
	if err != nil {
		return item, err
	}
	item.Amount = amount

	return item, nil

}

//

// TestResponse ..
type TestResponse struct {
	Error int64  `json:"error"`
	Msg   string `json:"msg"`
}

// AccountInfoResponse ...
/* {
  "accNo": 1234,
  "btcBal": "23.00000000",
  "btcDeposit": "1FkrHkVAFg5Jn3s2njdnWFcbizMYbb423W",
  "email": "goondoo@hotmail.com",
  "error": 0,
  "sgdBal": "57.50"
} */
type AccountInfoResponse struct {
	AccNo      int64           `json:"accNo"`
	BtcBal     decimal.Decimal `json:"btcBal"`
	BtcDeposit string          `json:"btcDeposit"`
	Email      string          `json:"email"`
	Error      int64           `json:"error"`
	SgdBal     decimal.Decimal `json:"sgdBal"`
	Msg        string          `json:"msg"` // for error handling
}

// PendingOrderResponse ...
/*
{
  "error": 0,
  "orders": [
    {
      "date": 1387099682,
      "price": "5.00",
      "qty": "0.99000000",
      "ticket": 6,
      "type": "S"
    },
    {
      "date": 1386932631,
      "price": "2.00",
      "qty": "0.99000000",
      "ticket": 5,
      "type": "B"
    },
    {
      "date": 1386099367,
      "price": "1.00",
      "qty": "1.00000000",
      "ticket": 4,
      "type": "B"
    }
  ]
}*/
type PendingOrderResponse struct {
	Error  int64  `json:"error"`
	Msg    string `json:"msg"` // for error handling.
	Orders []struct {
		Date int64 `json:"date"`

		Price string `json:"price"`
		Qty   string `json:"qty"`

		Type   string `json:"type"`
		Ticket int64  `json:"ticket"`
	} `json:"orders"`
}

// OrderHistoryResponse ...
/* {
  "error": 0,
  "orders": [
    {
      "date_created": 1387971414,
      "date_executed": 1387971414,
      "price": "S$3.00",
      "qty": "2.00000000BTC",
      "status": "A",
      "ticket": 11,
      "type": "B"
    },
    {
      "date_created": 1387971314,
      "date_executed": 1387971414,
      "price": "S$3.00",
      "qty": "2.00000000BTC",
      "status": "F",
      "ticket": 6,
      "type": "S"
    },
    {
      "date_created": 1387971414,
      "date_executed": 1387971414,
      "price": "S$5.00",
      "qty": "1.00000000BTC",
      "status": "A",
      "ticket": 12,
      "type": "B"
    },
    {
      "date_created": 1387971306,
      "date_executed": 1387971414,
      "price": "S$5.00",
      "qty": "1.00000000BTC",
      "status": "F",
      "ticket": 5,
      "type": "S"
    },
    {
      "date_created": 1387971398,
      "date_executed": 1387971398,
      "price": "S$2.50",
      "qty": "1.00000000BTC",
      "status": "F",
      "ticket": 10,
      "type": "B"
    },
    {
      "date_created": 1387971335,
      "date_executed": 1387971398,
      "price": "S$2.50",
      "qty": "1.00000000BTC",
      "status": "F",
      "ticket": 8,
      "type": "S"
    }
  ]
}*/
type OrderHistoryResponse struct {
	Error  int64  `json:"error"`
	Msg    string `json:"msg"` // for error handling
	Orders []struct {
		DateCreated  int64  `json:"date_created"`
		DateExecuted int64  `json:"date_executed"`
		Price        string `json:"price"`
		Qty          string `json:"qty"`
		//Status int64 `json:"status"` // it's wrong, erro occured
		//Type   int64 `json:"type"`// it's wrong, erro occured
		Ticket int64 `json:"ticket"`
	} `json:"orders"`
}

// PlaceOrderResponse ..
/*
{
  "error": 0,
  "msg": "",
  "pending_oid": "28"
}*/
type PlaceOrderResponse struct {
	Error      int64  `json:"error"`
	Msg        string `json:"msg"`
	PendingOID string `json:"pending_oid"`
}

// CancelPendingOrderResponse ...
/*

{
  "error": 0
}
*/
type CancelPendingOrderResponse struct {
	Error int64  `json:"error"`
	Msg   string `json:"msg,omitempty"`
}

// WithdrawResponse ..
/*

{
  "error": 0,
  "msg": "11750"
}*/
type WithdrawResponse struct {
	Error int64  `json:"error"`
	Msg   string `json:"msg"`
}

// KeyPermissionErrorResponse ...
type KeyPermissionErrorResponse struct {
	Error string `json:"error"`
}
