package lastSymbol

type PriceInterface interface {
	Init(instrumentId string) PriceInterface
	GetLastSymbolByOne() (result interface{}, err error)
}
