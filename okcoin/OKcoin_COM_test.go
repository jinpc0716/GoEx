package okcoin

import (
	"net/http"
	"testing"

	"github.com/leek-box/GoEx"
)

var okcom = NewCOM(http.DefaultClient, "", "")

func TestOKCoinCOM_API_GetTicker(t *testing.T) {
	ticker, _ := okcom.GetTicker(goex.BTC_CNY)
	t.Log(ticker)
}
