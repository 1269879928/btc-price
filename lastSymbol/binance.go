package lastSymbol

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/parnurzeal/gorequest"
)

type LastBinancePrice struct {
	Symbol string
}
// Symbol string 币对名称
func (tx *LastBinancePrice) Init(symbol string) PriceInterface {
	return &LastBinancePrice{Symbol: symbol}
}


func (tx *LastBinancePrice)GetLastSymbolByOne() (result interface{}, err error)  {
	uri := "https://api.binance.com/api/v3/ticker/price?symbol=" + tx.Symbol
	request := gorequest.New()
	_, body, errs := request.Get(uri).End()

	if errs != nil {
		err = fmt.Errorf("get huobi lasted price failed, err:%v\n", errs[0])
		return
	}
	result, err = gabs.ParseJSON([]byte(body))
	if err != nil {
		err = fmt.Errorf("get huobi lasted price failed, err:%v\n", err.Error())
		return
	}
	return
}
