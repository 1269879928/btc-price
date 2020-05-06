package trades

import (
	"fmt"
	"price/okexSDK"
)

func TestOKExServerTime() {
	//serverTime, err := NewOKExClient().GetSpotInstrumentTicker("BTC-USDT")
	serverTime, err := NewOKExClient().GetAccountWallet()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("OKEx's server time: ", serverTime)
}

func NewOKExClient() *okexSDK.Client {
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