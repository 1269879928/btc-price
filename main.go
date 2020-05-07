package main

import (
	"fmt"
	"price/lastSymbol"
)

func main()  {
	// okex
	//res, err := new(lastSymbol.LastOkexPrice).Init("BTC-USDT").GetLastSymbolByOne()
	//fmt.Println(res, err)
	// 火币
	//res, err := new(lastSymbol.LastHuoBiPrice).Init("btcusdt").GetLastSymbolByOne()
	//fmt.Println(res, err)
	// Binance
	//res, err := new(lastSymbol.LastBinancePrice).Init("BTCUSDT").GetLastSymbolByOne()
	//fmt.Println(res, err)
	// zb
	res, err := new(lastSymbol.LastZbPrice).Init("btc_usdt").GetLastSymbolByOne()
	fmt.Println(res, err)

}