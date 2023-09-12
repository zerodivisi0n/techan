package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type meanDeviationIndicator struct {
	indicator     Indicator
	movingAverage Indicator
	window        int
}

// NewMeanDeviationIndicator returns a derivative Indicator which returns the mean deviation of a base indicator
// in a given window. Mean deviation is an average of all values on the base indicator from the mean of that indicator.
func NewMeanDeviationIndicator(indicator Indicator, window int) Indicator {
	return meanDeviationIndicator{
		indicator:     indicator,
		movingAverage: NewSimpleMovingAverage(indicator, window),
		window:        window,
	}
}

func (mdi meanDeviationIndicator) Calculate(index int) big.Decimal {
	if index < mdi.window-1 {
		return big.ZERO
	}

	average := mdi.movingAverage.Calculate(index)
	start := Max(0, index-mdi.window+1)
	absoluteDeviations := big.NewDecimal(0)

	for i := start; i <= index; i++ {
		absoluteDeviations = absoluteDeviations.Add(average.Sub(mdi.indicator.Calculate(i)).Abs())
	}

	return absoluteDeviations.Div(big.NewDecimal(float64(Min(mdi.window, index-start+1))))
}

func (mdi meanDeviationIndicator) LastIndex() int {
	return mdi.indicator.LastIndex()
}

func (mdi meanDeviationIndicator) Key() string {
	return fmt.Sprintf("mdi(%d):%s", mdi.window, mdi.indicator.Key())
}
