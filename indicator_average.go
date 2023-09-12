package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type averageIndicator struct {
	indicator Indicator
	window    int
}

// NewAverageGainsIndicator Returns a new average gains indicator, which returns the average gains
// in the given window based on the given indicator.
func NewAverageGainsIndicator(indicator Indicator, window int) Indicator {
	return averageIndicator{
		indicator: NewCumulativeGainsIndicator(indicator, window),
		window:    window,
	}
}

// NewAverageLossesIndicator Returns a new average losses indicator, which returns the average losses
// in the given window based on the given indicator.
func NewAverageLossesIndicator(indicator Indicator, window int) Indicator {
	return averageIndicator{
		indicator: NewCumulativeLossesIndicator(indicator, window),
		window:    window,
	}
}

func (ai averageIndicator) Calculate(index int) big.Decimal {
	return ai.indicator.Calculate(index).Div(big.NewDecimal(float64(Min(index+1, ai.window))))
}

func (ai averageIndicator) LastIndex() int {
	return ai.indicator.LastIndex()
}

func (ai averageIndicator) Key() string {
	return fmt.Sprintf("avg(%d):%s", ai.window, ai.indicator.Key())
}
