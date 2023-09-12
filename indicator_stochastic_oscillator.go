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
	return kIndicator{
		closePrice: NewClosePriceIndicator(series),
		minValue:   NewMinimumValueIndicator(NewLowPriceIndicator(series), timeframe),
		maxValue:   NewMaximumValueIndicator(NewHighPriceIndicator(series), timeframe),
		window:     timeframe,
	}
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
	k      Indicator
	window int
}

// NewSlowStochasticIndicator returns a derivative Indicator which returns the slow stochastic indicator (%D) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewSlowStochasticIndicator(k Indicator, window int) Indicator {
	return dIndicator{k, window}
}

func (d dIndicator) Calculate(index int) big.Decimal {
	return NewSimpleMovingAverage(d.k, d.window).Calculate(index)
}

func (d dIndicator) LastIndex() int {
	return d.k.LastIndex()
}

func (d dIndicator) Key() string {
	return fmt.Sprintf("stochd(%d):%s", d.window, d.k.Key())
}
