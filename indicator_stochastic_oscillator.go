package techan

import (
	"fmt"
	"math"

	"github.com/sdcoffey/big"
)

type kIndicator struct {
	closePrice Indicator
	minValue   Indicator
	maxValue   Indicator
	window     int
}

// NewFastStochasticIndicator returns a derivative Indicator which returns the fast stochastic indicator (%K) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewFastStochasticIndicator(series TimeSeries, timeframe int) Indicator {
	return NewFastStochasticIndicatorWithProxy(DefaultProxy, series, timeframe)
}

func NewFastStochasticIndicatorWithProxy(proxy IndicatorProxy, series TimeSeries, timeframe int) Indicator {
	return proxy(kIndicator{
		closePrice: NewClosePriceIndicator(series),
		minValue:   proxy(NewMinimumValueIndicator(NewLowPriceIndicator(series), timeframe)),
		maxValue:   proxy(NewMaximumValueIndicator(NewHighPriceIndicator(series), timeframe)),
		window:     timeframe,
	})
}

func (k kIndicator) Calculate(index int) big.Decimal {
	closeVal := k.closePrice.Calculate(index)
	minVal := k.minValue.Calculate(index)
	maxVal := k.maxValue.Calculate(index)

	if minVal.EQ(maxVal) {
		return big.NewDecimal(math.Inf(1))
	}

	return closeVal.Sub(minVal).Div(maxVal.Sub(minVal)).Mul(big.NewDecimal(100))
}

func (k kIndicator) LastIndex() int {
	return k.closePrice.LastIndex()
}

func (k kIndicator) Key() string {
	return fmt.Sprintf("stochk(%d):%s", k.window, k.closePrice.Key())
}

type dIndicator struct {
	kSma   Indicator
	window int
}

// NewSlowStochasticIndicator returns a derivative Indicator which returns the slow stochastic indicator (%D) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewSlowStochasticIndicator(k Indicator, window int) Indicator {
	return NewSlowStochasticIndicatorWithProxy(DefaultProxy, k, window)
}

func NewSlowStochasticIndicatorWithProxy(proxy IndicatorProxy, k Indicator, window int) Indicator {
	return proxy(dIndicator{
		kSma:   NewSimpleMovingAverage(k, window),
		window: window,
	})
}

func (d dIndicator) Calculate(index int) big.Decimal {
	return d.kSma.Calculate(index)
}

func (d dIndicator) LastIndex() int {
	return d.kSma.LastIndex()
}

func (d dIndicator) Key() string {
	return fmt.Sprintf("stochd(%d):%s", d.window, d.kSma.Key())
}
