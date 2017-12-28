package fyb

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
func TestPlaceOrder(t *testing.T) {
	token := os.Getenv("FYBSG_KEY")
	secret := os.Getenv("FYBSG_SECRET")
	api := New(APIBaseURLForTest, token, secret)
	ret, body, err := api.PlaceOrder("BUY", 1.2, 1.1)
	require.NoError(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("ret.Error:%d", ret.Error)
	log.Printf("ret.PendingOID:%s", ret.PendingOID)
	log.Printf("ret.Msg:%s", ret.Msg)

	ret2, body, err2 := api.PlaceOrder("SELL", 999999.99, 0.01)
	require.NoError(t, err2, nil)
	log.Printf("body:%s", string(body))
	log.Printf("ret2.Error:%d", ret2.Error)
	log.Printf("ret2.PendingOID:%s", ret2.PendingOID)
	log.Printf("ret2.Msg:%s", ret2.Msg)

	return
}
*/
func TestPrivateAPITest(t *testing.T) {

	token := os.Getenv("FYBSG_KEY")
	secret := os.Getenv("FYBSG_SECRET")
	api := New(APIBaseURLForTest, token, secret)
	ret, body, err := api.APITokenTest()
	log.Print(err)

	log.Printf("body:%s", string(body))
	require.NoError(t, err, nil)
	require.Equal(t, "success", ret.Msg, nil)
	return
}

func TestCancelPendingOrdersFail(t *testing.T) {
	log.Print("TestCancelPendingOrdersFail")
	api := New(APIBaseURLForTest, "wrongtoken", "wrongsecret")
	ret, body, err := api.GetPendingOrders()

	require.Error(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("err:%v", err)
	log.Printf("ret.Msg:%s", ret.Msg)

	return
}

func TestWithdrawFail(t *testing.T) {
	log.Print("TestWithdrawFail")
	api := New(APIBaseURLForTest, "wrongtoken", "wrongsecret")
	ret, body, err := api.Withdraw(0.01, "aaa", "BTC")
	require.Error(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("err:%v", err)
	log.Printf("ret.Msg:%s", ret.Msg)

	return
}
func TestPlaceOrderFail(t *testing.T) {
	log.Print("TestPlaceOrderFail")
	api := New(APIBaseURLForTest, "wrongtoken", "wrongsecret")
	ret, body, err := api.PlaceOrder("BUY", 1.2, 1.1)

	require.Error(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("err:%v", err)
	log.Printf("ret.Msg:%s", ret.Msg)

	return
}

func TestGetPendingOrdersFail(t *testing.T) {
	log.Print("TestGetPendingOrdersFail")
	api := New(APIBaseURLForTest, "wrongtoken", "wrongsecret")

	ret, body, err := api.GetPendingOrders()

	require.Error(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("err:%v", err)
	log.Printf("ret.Msg:%s", ret.Msg)

	return
}
func TestGetAccountInfoFail(t *testing.T) {

	api := New(APIBaseURLForTest, "wrongtoken", "wrongsecret")
	ret, body, err := api.GetAccountInfo()
	require.Error(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("err:%v", err)
	log.Printf("ret.Msg:%s", ret.Msg)
	return
}

func TestPrivateAPIFail(t *testing.T) {

	api := New(APIBaseURLForTest, "wrongtoken", "wrongsecret")
	ret, body, err := api.APITokenTest()
	require.Error(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("err:%v", err)
	log.Printf("ret.Msg:%s", ret.Msg)
	return
}

func TestGetOrderHistoryFail(t *testing.T) {
	log.Printf("TestGetOrderHistoryFail")

	api := New(APIBaseURLForTest, "wrongtoken", "wrongsecret")
	ret, body, err := api.GetOrderHistory(5)
	require.Error(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("err:%v", err)
	log.Printf("ret.Msg:%s", ret.Msg)

	return
}

func TestOrderBook(t *testing.T) {

	api := New(APIBaseURLForTest, "", "")
	orderbook, body, err := api.GetOrderBook()

	require.NoError(t, err, nil)
	log.Printf("body:%s", string(body))

	for _, ask := range orderbook.Asks {
		log.Printf("ask: price=%v, amount=%v", ask.Price, ask.Amount)
	}
	for _, bid := range orderbook.Bids {
		log.Printf("bid: price=%v, amount=%v", bid.Price, bid.Amount)
	}

	return
}

func TestGetTicker(t *testing.T) {

	api := New(APIBaseURLForTest, "", "")

	ticker, body, err := api.GetTicker()

	require.NoError(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("ticker.Ask:%v", ticker.Ask)
	log.Printf("ticker.Bid:%v", ticker.Bid)
	log.Printf("ticker.Last:%v", ticker.Last)
	log.Printf("ticker.Vol:%v", ticker.Vol)

	return
}

func TestGetTradeHistoryTestTrades(t *testing.T) {

	api := New(APIBaseURLForTest, "", "")
	tradeHistory, body, err := api.GetTradeHistory(2218610)
	require.NoError(t, err, nil)
	log.Printf("body:%s", string(body))
	for _, trade := range tradeHistory {
		log.Printf("trade.Date:%d", trade.Date)
		log.Printf("trade.TID:%d", trade.TID)
		log.Printf("trade.Amount:%v", trade.Amount)
		log.Printf("trade.Price:%v", trade.Price)
	}

	return
}

func TestGetAccountInfo(t *testing.T) {

	token := os.Getenv("FYBSG_KEY")
	secret := os.Getenv("FYBSG_SECRET")
	api := New(APIBaseURLForTest, token, secret)
	ret, body, err := api.GetAccountInfo()
	require.NoError(t, err, nil)
	log.Printf("body:%s", string(body))
	//require.Equal(t, "success", ret.Msg, nil)
	log.Printf("ret.AccNo:%v", ret.AccNo)
	log.Printf("ret.BtcBal:%v", ret.BtcBal)
	log.Printf("ret.BtcDeposit:%v", ret.BtcDeposit)
	log.Printf("ret.Email:%v", ret.Email)
	log.Printf("ret.Error:%v", ret.Error)
	log.Printf("ret.SgdBal:%v", ret.SgdBal)
	return
}

func TestGetOrderHistory(t *testing.T) {
	log.Printf("TestGetOrderHistory")
	token := os.Getenv("FYBSG_KEY")
	secret := os.Getenv("FYBSG_SECRET")
	api := New(APIBaseURLForTest, token, secret)
	ret, body, err := api.GetOrderHistory(5)

	require.NoError(t, err, nil)
	log.Printf("body:%s", string(body))

	for _, order := range ret.Orders {
		log.Printf("order.DateExecuted:%d", order.DateExecuted)
		log.Printf("order.DateCreated:%d", order.DateCreated)
		log.Printf("order.Price:%s", order.Price)
		log.Printf("order.Qty:%s", order.Qty)
	}

	return
}

func TestGetPendingOrders(t *testing.T) {

	token := os.Getenv("FYBSG_KEY")
	secret := os.Getenv("FYBSG_SECRET")
	api := New(APIBaseURLForTest, token, secret)
	ret, body, err := api.GetPendingOrders()
	require.NoError(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("ret.Error:%d", ret.Error)
	for _, order := range ret.Orders {
		log.Printf("======")
		log.Printf("order.Date:%v", order.Date)
		log.Printf("order.Price:%v", order.Price)
		log.Printf("order.Qty:%v", order.Qty)
		log.Printf("order.Ticket:%d", order.Ticket)
		log.Printf("order.Type:%s", order.Type)

	}

	return
}

func TestCancelPendingOrders(t *testing.T) {
	token := os.Getenv("FYBSG_KEY")
	secret := os.Getenv("FYBSG_SECRET")
	api := New(APIBaseURLForTest, token, secret)
	ret, body, err := api.GetPendingOrders()
	require.NoError(t, err, nil)
	log.Printf("body:%s", string(body))
	log.Printf("ret.Error:%d", ret.Error)
	for _, order := range ret.Orders {
		log.Printf("======")
		log.Printf("order.Date:%v", order.Date)
		log.Printf("order.Price:%v", order.Price)
		log.Printf("order.Qty:%v", order.Qty)
		log.Printf("order.Ticket:%d", order.Ticket)
		log.Printf("order.Type:%s", order.Type)
		log.Printf("CancelPendingOrder(%d)", order.Ticket)

		ret2, body, err2 := api.CancelPendingOrder(order.Ticket)
		log.Printf("body:%s", string(body))

		require.NoError(t, err2, nil)
		log.Printf("ret2.Error:%d", ret2.Error)
	}

	return
}

/*
func TestWithdraw(t *testing.T) {
	log.Printf("TestWithdraw")
	token := os.Getenv("FYBSG_KEY")
	secret := os.Getenv("FYBSG_SECRET")
	api := New(APIBaseURLForTest, token, secret)
	ret, body, err := api.Withdraw(0.01, "aaa", "BTC")
	log.Printf("err=%v", err)
	log.Printf("body:%s",string(body))
	require.NoError(t, err, nil)
	log.Printf("ret:%v", ret)
	log.Printf("ret.Error:%v", ret.Error)
	//log.Printf("ret.Msg:%s", ret.Msg)

	return
}

*/
