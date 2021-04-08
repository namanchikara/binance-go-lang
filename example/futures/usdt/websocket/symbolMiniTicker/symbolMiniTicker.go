package main

import (
	"fmt"
	"github.com/dirname/Binance/futures/usdt/websocket/market"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
)

func main() {
	client := futuresusdt.NewUSDTFuturesSymbolMiniTickerWebsocketClient("btcusdt@miniTicker")
	client.SetHandler(func() {
		client.Subscribe(123, "btcusdt@miniTicker", "ltcusdt@miniTicker")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.SymbolMiniTickerResponse:
			logger.Info("SymbolMiniTicker: %v", response.(futuresusdt.SymbolMiniTickerResponse))
		case futuresusdt.SymbolMiniTickerCombinedResponse:
			logger.Info("SymbolMiniTickerCombinedResponse: %v", response.(futuresusdt.SymbolMiniTickerCombinedResponse))
		case model.WebsocketCommonResponse:
			logger.Info("Websocket Response: %v", response.(model.WebsocketCommonResponse))
		case model.WebsocketErrorResponse:
			logger.Info("Websocket Error Response: %v", response.(model.WebsocketErrorResponse))
		default:
			logger.Info("Unknown Response: %v", response)
		}
	})
	client.Connect(true)
	fmt.Scanln()

	client.Unsubscribe(123, "btcusdt@miniTicker", "ltcusdt@miniTicker")
	client.Close()
	logger.Info("Client closed")
}
