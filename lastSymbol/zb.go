package lastSymbol

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/parnurzeal/gorequest"
)

type LastZbPrice struct {
	Symbol string
}
// symbol string 币对名称 如：btc_usdt
func (tx *LastZbPrice) Init(symbol string) PriceInterface {
	return &LastZbPrice{Symbol: symbol}
}


func (tx *LastZbPrice)GetLastSymbolByOne() (result interface{}, err error)  {
	uri := "http://api.zb.live/data/v1/ticker?market=" + tx.Symbol
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
