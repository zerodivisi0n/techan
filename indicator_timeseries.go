package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type indicatorTimeSeries struct {
	ts         TimeSeries
	openPrice  Indicator
	closePrice Indicator
	highPrice  Indicator
	lowPrice   Indicator
	volume     Indicator
}

var _ TimeSeries = indicatorTimeSeries{}

// NewIndicatorTimeSeries returns a new time series with given indicator
func NewIndicatorTimeSeries(ts TimeSeries, indicatorCtor func(indicator Indicator) Indicator) TimeSeries {
	return NewIndicatorTimeSeriesWithProxy(DefaultProxy, ts, indicatorCtor)
}

func NewIndicatorTimeSeriesWithProxy(proxy IndicatorProxy, ts TimeSeries, indicatorCtor func(indicator Indicator) Indicator) TimeSeries {
	return indicatorTimeSeries{
		ts:         ts,
		openPrice:  proxy(indicatorCtor(NewOpenPriceIndicator(ts))),
		closePrice: proxy(indicatorCtor(NewClosePriceIndicator(ts))),
		highPrice:  proxy(indicatorCtor(NewHighPriceIndicator(ts))),
		lowPrice:   proxy(indicatorCtor(NewLowPriceIndicator(ts))),
		volume:     proxy(indicatorCtor(NewVolumeIndicator(ts))),
	}
}

// NewTimeSeriesFromIndicator returns a time series with values from an indicator
func NewTimeSeriesFromIndicator(ts TimeSeries, indicator Indicator) TimeSeries {
	return indicatorTimeSeries{
		ts:         ts,
		openPrice:  indicator,
		closePrice: indicator,
		highPrice:  indicator,
		lowPrice:   indicator,
		volume:     indicator,
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

func (its indicatorTimeSeries) LastIndex() int {
	return its.ts.LastIndex()
}

func (its indicatorTimeSeries) Key() string {
	return fmt.Sprintf("timeseries(%s):%s",
		its.closePrice.Key(), its.ts.Key())
}
