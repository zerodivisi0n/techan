package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type averageTrueRangeIndicator struct {
	indicator Indicator
	window    int
}

// NewAverageTrueRangeIndicator returns a base indicator that calculates the average true range of the
// underlying over a window
// https://www.investopedia.com/terms/a/atr.asp
func NewAverageTrueRangeIndicator(series TimeSeries, window int) Indicator {
	return averageTrueRangeIndicator{
		indicator: NewTrueRangeIndicator(series),
		window:    window,
	}
}

func (atr averageTrueRangeIndicator) Calculate(index int) big.Decimal {
	if index < atr.window {
		return big.ZERO
	}

	sum := big.ZERO

	for i := index; i > index-atr.window; i-- {
		sum = sum.Add(atr.indicator.Calculate(i))
	}

	return sum.Div(big.NewFromInt(atr.window))
}

func (atr averageTrueRangeIndicator) LastIndex() int {
	return atr.indicator.LastIndex()
}

func (atr averageTrueRangeIndicator) Key() string {
	return fmt.Sprintf("atr(%d):%s", atr.window, atr.indicator.Key())
}
