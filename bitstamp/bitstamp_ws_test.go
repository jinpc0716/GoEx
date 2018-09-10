package bitstamp

import (
	"log"
	"testing"
	"time"

	"github.com/leek-box/GoEx"
)

func TestBitstamp_GetDepthWithWs(t *testing.T) {
	btmp.GetDepthWithWs(goex.BCH_USD, func(depth *goex.Depth) {
		log.Println(depth)
	})
	btmp.GetDepthWithWs(goex.LTC_USD, func(depth *goex.Depth) {
		log.Println(depth)
	})
	time.Sleep(1 * time.Minute)
}
