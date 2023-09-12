package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type smaIndicator struct {
	indicator Indicator
	window    int
}

// NewSimpleMovingAverage returns a derivative Indicator which returns the average of the current value and preceding
// values in the given windowSize.
func NewSimpleMovingAverage(indicator Indicator, window int) Indicator {
	return smaIndicator{indicator, window}
}

func (sma smaIndicator) Calculate(index int) big.Decimal {
	if index < sma.window-1 {
		return big.ZERO
	}

	sum := big.ZERO
	for i := index; i > index-sma.window; i-- {
		sum = sum.Add(sma.indicator.Calculate(i))
	}

	result := sum.Div(big.NewFromInt(sma.window))

	return result
}

func (sma smaIndicator) LastIndex() int {
	return sma.indicator.LastIndex()
}

func (sma smaIndicator) Key() string {
	return fmt.Sprintf("sma:%s", sma.indicator.Key())
}
