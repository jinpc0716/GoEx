package c_cex

import (
	"net/http"
	"testing"

	"github.com/leek-box/GoEx"
)

var ccex = New(http.DefaultClient, "", "")

func TestCcex_GetTicker(t *testing.T) {
	ticker, err := ccex.GetTicker(goex.BTC_USD)
	t.Log("err=>", err)
	t.Log("ticker=>", ticker)
}
