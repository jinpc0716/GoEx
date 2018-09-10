package btcmarkets

import (
	"net/http"
	"testing"

	"github.com/leek-box/GoEx"
)

var btcm = New(http.DefaultClient, "", "")

func TestBtcm_GetTicker(t *testing.T) {
	AUD := goex.NewCurrency("AUD", "")
	BTC_AUD := goex.NewCurrencyPair(goex.BTC, AUD)
	ticker, err := btcm.GetTicker(BTC_AUD)
	t.Log("err=>", err)
	t.Log("ticker=>", ticker)
}
