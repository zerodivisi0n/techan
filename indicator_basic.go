package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type volumeIndicator struct {
	TimeSeries
}

// NewVolumeIndicator returns an indicator which returns the volume of a candle for a given index
func NewVolumeIndicator(series TimeSeries) Indicator {
	return volumeIndicator{series}
}

func (vi volumeIndicator) Calculate(index int) big.Decimal {
	return vi.Volume(index)
}

func (vi volumeIndicator) Key() string {
	return fmt.Sprintf("volume(%s)", vi.TimeSeries.Key())
}

type closePriceIndicator struct {
	TimeSeries
}

// NewClosePriceIndicator returns an Indicator which returns the close price of a candle for a given index
func NewClosePriceIndicator(series TimeSeries) Indicator {
	return closePriceIndicator{series}
}

func (cpi closePriceIndicator) Calculate(index int) big.Decimal {
	return cpi.ClosePrice(index)
}

func (cpi closePriceIndicator) Key() string {
	return fmt.Sprintf("close(%s)", cpi.TimeSeries.Key())
}

type highPriceIndicator struct {
	TimeSeries
}

// NewHighPriceIndicator returns an Indicator which returns the high price of a candle for a given index
func NewHighPriceIndicator(series TimeSeries) Indicator {
	return highPriceIndicator{
		series,
	}
}

func (hpi highPriceIndicator) Calculate(index int) big.Decimal {
	return hpi.HighPrice(index)
}

func (hpi highPriceIndicator) Key() string {
	return fmt.Sprintf("high(%s)", hpi.TimeSeries.Key())
}

type lowPriceIndicator struct {
	TimeSeries
}

// NewLowPriceIndicator returns an Indicator which returns the low price of a candle for a given index
func NewLowPriceIndicator(series TimeSeries) Indicator {
	return lowPriceIndicator{
		series,
	}
}

func (lpi lowPriceIndicator) Calculate(index int) big.Decimal {
	return lpi.LowPrice(index)
}

func (lpi lowPriceIndicator) Key() string {
	return fmt.Sprintf("low(%s)", lpi.TimeSeries.Key())
}

type openPriceIndicator struct {
	TimeSeries
}

// NewOpenPriceIndicator returns an Indicator which returns the open price of a candle for a given index
func NewOpenPriceIndicator(series TimeSeries) Indicator {
	return openPriceIndicator{
		series,
	}
}

func (opi openPriceIndicator) Calculate(index int) big.Decimal {
	return opi.OpenPrice(index)
}

func (opi openPriceIndicator) Key() string {
	return fmt.Sprintf("open(%s)", opi.TimeSeries.Key())
}

type typicalPriceIndicator struct {
	TimeSeries
}

// NewTypicalPriceIndicator returns an Indicator which returns the typical price of a candle for a given index.
// The typical price is an average of the high, low, and close prices for a given candle.
func NewTypicalPriceIndicator(series TimeSeries) Indicator {
	return typicalPriceIndicator{series}
}

func (tpi typicalPriceIndicator) Calculate(index int) big.Decimal {
	numerator := tpi.HighPrice(index).Add(tpi.LowPrice(index)).Add(tpi.ClosePrice(index))
	return numerator.Div(big.NewFromString("3"))
}

func (tpi typicalPriceIndicator) Key() string {
	return fmt.Sprintf("typical(%s)", tpi.TimeSeries.Key())
}
