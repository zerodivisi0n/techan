package techan

import "github.com/sdcoffey/big"

type indicatorTimeSeries struct {
	TimeSeries
	openPrice  Indicator
	closePrice Indicator
	highPrice  Indicator
	lowPrice   Indicator
	volume     Indicator
}

var _ TimeSeries = indicatorTimeSeries{}

// NewIndicatorTimeSeries returns a new time series with given indicator
func NewIndicatorTimeSeries(ts TimeSeries, indicatorCtor func(indicator Indicator) Indicator) TimeSeries {
	return indicatorTimeSeries{
		TimeSeries: ts,
		openPrice:  indicatorCtor(NewOpenPriceIndicator(ts)),
		closePrice: indicatorCtor(NewClosePriceIndicator(ts)),
		highPrice:  indicatorCtor(NewHighPriceIndicator(ts)),
		lowPrice:   indicatorCtor(NewLowPriceIndicator(ts)),
		volume:     indicatorCtor(NewVolumeIndicator(ts)),
	}
}

func (its indicatorTimeSeries) OpenPrice(index int) big.Decimal {
	return its.openPrice.Calculate(index)
}

func (its indicatorTimeSeries) ClosePrice(index int) big.Decimal {
	return its.closePrice.Calculate(index)
}

func (its indicatorTimeSeries) HighPrice(index int) big.Decimal {
	return its.highPrice.Calculate(index)
}

func (its indicatorTimeSeries) LowPrice(index int) big.Decimal {
	return its.lowPrice.Calculate(index)
}

func (its indicatorTimeSeries) Volume(index int) big.Decimal {
	return its.volume.Calculate(index)
}
