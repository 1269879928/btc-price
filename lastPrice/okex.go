package lastPrice

import (
	"fmt"
	"price/okexSDK"
)

type LastPrice struct {
	InstrumentId string
}
func (tx * LastPrice)GetLastPrice() (result interface{})  {
	result, err := newOKExClient().GetSpotInstrumentTicker(tx.InstrumentId)
	if err != nil {
		fmt.Println("err:", err)
	}
	return
}

func newOKExClient() *okexSDK.Client {
	var config okexSDK.Config
	config.Endpoint = "https://www.okex.me/"
	config.ApiKey = "e0988b51-98ed-4842-8466-0160151bc2ca"
	config.SecretKey = "11FCF8EC99ECDCDD2D02D8CCE5296843"
	config.Passphrase = "uuu123iii"
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = okexSDK.ENGLISH

	client := okexSDK.NewClient(config)
	return client
}