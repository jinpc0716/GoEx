package aex

import (
	"net/http"
	"testing"

	"github.com/leek-box/GoEx"
)

var acx = New(http.DefaultClient, "", "", "")

func TestAex_GetTicker(t *testing.T) {
	ticker, err := acx.GetTicker(goex.ETH_BTC)
	t.Log("err=>", err)
	t.Log("ticker=>", ticker)
}
