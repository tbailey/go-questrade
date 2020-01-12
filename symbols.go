package questrade

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Symbol struct {
	Symbol string
	SymbolId uint64
	PrevDayClosePrice float64
	HighPrice52 float64
	LowPrice52 float64
	AverageVol3Months uint64
	AverageVol20Days uint64
	OutstandingShares uint64
	EPS float64
	PE float64
	Dividend float64
	Yield float64
	ExDate time.Time
	MarketCap uint64
	OptionType string
	OptionDurationType string
	OptionRoot string
	OptionContractDeliverables OptionOrderDeliverable
	OptionExerciseType string
	ListingExchange string
	Description string
	SecurityType string
	OptionExpiryDate time.Time
	DividendDate time.Time
	OptionStrikePrice uint64
	IsTradable bool
	IsQuotable bool
	HasOptions bool
	MinTicks []MinTickData
	IndustrySector string
	IndustryGroup string
	IndustrySubGroup string
}

type Symbols struct {
	Symbols []Symbol
}

type MinTickData struct {
	Pivot uint64
	MintTick uint64
}

type OptionOrderDeliverable struct {
	Underlyings []UnderlyingMultiplierPair
	CashInLieu uint64
}

type UnderlyingMultiplierPair struct {
	Multiplier uint64
	UnderlyingSymbol string
	UnderlyingSymbolId string
}

func (c *Client) getSymbols(parameters url.Values) (*Symbols, error) {
	u := fmt.Sprintf("symbols?%v", parameters.Encode())
	fmt.Println(u)

	var symbols Symbols
	err := c.NewRequest("GET", u, nil, &symbols)
	if err != nil {
		return nil, err
	}

	return &symbols, nil
}

func (c *Client) SymbolById(symbolId int) (*Symbols, error) {
	u := fmt.Sprintf("symbols/%v", symbolId)
	fmt.Println(u)

	var symbols Symbols
	err := c.NewRequest("GET", u, nil, &symbols)
	if err != nil {
		return nil, err
	}

	return &symbols, nil
}

func (c *Client) SymbolsById(symbolIds []int) ([]Symbol, error) {
	parameters := url.Values{}
	parameters.Add("ids", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(symbolIds)), ","), "[]"))

	symbols, err := c.getSymbols(parameters)
	if err != nil {
		return nil, err
	}

	return symbols.Symbols, nil
}

func (c *Client) SymbolsByName(symbolNames []string) ([]Symbol, error) {
	parameters := url.Values{}
	parameters.Add("names", strings.Join(symbolNames, ","))

	symbols, err := c.getSymbols(parameters)
	if err != nil {
		return nil, err
	}

	return symbols.Symbols, nil
}