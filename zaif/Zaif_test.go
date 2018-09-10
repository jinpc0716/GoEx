package zaif

import (
	"net/http"
	"testing"

	"github.com/leek-box/GoEx"
	"github.com/stretchr/testify/assert"
)

var api = New(http.DefaultClient, "", "")
var BTC_JPY = goex.NewCurrencyPair(goex.BTC, goex.NewCurrency("JPY", ""))

func TestZaif_GetTicker(t *testing.T) {
	ticker, err := api.GetTicker(BTC_JPY)
	assert.Empty(t, err)
	t.Log(ticker)
}

func TestZaif_GetDepth(t *testing.T) {
	depth, err := api.GetDepth(4, BTC_JPY)
	assert.Empty(t, err)
	t.Log(depth)
}
