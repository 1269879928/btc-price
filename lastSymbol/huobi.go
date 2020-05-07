package lastSymbol

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/parnurzeal/gorequest"
)

type LastHuoBiPrice struct {
	Symbol string
}
// symbol string 币对名称 如：btcusdt
func (tx *LastHuoBiPrice) Init(symbol string) PriceInterface {
	return &LastHuoBiPrice{Symbol: symbol}
}


func (tx *LastHuoBiPrice)GetLastSymbolByOne() (result interface{}, err error)  {
	//result, err := newOKExClient().GetSpotInstrumentTicker(tx.Symbol)
	//uri := "https://" + config.Host + "/market/detail/merged?symbol=" + tx.Symbol
	uri := "https://api.huobi.me/market/detail/merged?symbol=" + tx.Symbol
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