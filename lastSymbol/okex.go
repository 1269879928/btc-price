package lastSymbol

import (
	"fmt"
	"price/okexSDK"
)

type LastOkexPrice struct {
	InstrumentId string
}
// instrumentId string 币对名称
func (tx *LastOkexPrice) Init(instrumentId string) PriceInterface {
	return &LastOkexPrice{InstrumentId: instrumentId}
}

// 获取最新的交易行情详情
func (tx *LastOkexPrice)GetLastSymbolByOne() (result interface{}, err error)  {
	result, err = newOKExClient().GetSpotInstrumentTicker(tx.InstrumentId)
	if err != nil {
		err = fmt.Errorf("get okex lasted price failed, err:%v\n", err.Error())
		return
	}
	return
}

func newOKExClient() *okexSDK.Client {
	var config okexSDK.Config
	config.Endpoint = "https://www.okex.me/"
	config.ApiKey = "xxxx"
	config.SecretKey = "xxx"
	config.Passphrase = "xxxx"
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = okexSDK.ENGLISH

	client := okexSDK.NewClient(config)
	return client
}