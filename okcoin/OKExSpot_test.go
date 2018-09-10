package okcoin

import (
	"net/http"
	"testing"

	"github.com/leek-box/GoEx"
	"github.com/stretchr/testify/assert"
)

var okexSpot = NewOKExSpot(http.DefaultClient, "", "")

func TestOKExSpot_GetTicker(t *testing.T) {
	ticker, err := okexSpot.GetTicker(goex.ETC_BTC)
	assert.Nil(t, err)
	t.Log(ticker)
}

func TestOKExSpot_GetDepth(t *testing.T) {
	dep, err := okexSpot.GetDepth(2, goex.ETC_BTC)
	assert.Nil(t, err)
	t.Log(dep)
}

func TestOKExSpot_GetKlineRecords(t *testing.T) {
	klines, err := okexSpot.GetKlineRecords(goex.LTC_BTC, goex.KLINE_PERIOD_1MIN, 1000, -1)
	t.Log(err, klines)
}
