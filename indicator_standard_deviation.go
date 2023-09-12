package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

// NewStandardDeviationIndicator calculates the standard deviation of a base indicator.
// See https://www.investopedia.com/terms/s/standarddeviation.asp
func NewStandardDeviationIndicator(ind Indicator) Indicator {
	return standardDeviationIndicator{
		indicator: NewVarianceIndicator(ind),
	}
}

type standardDeviationIndicator struct {
	indicator Indicator
}

// Calculate returns the standard deviation of a base indicator
func (sdi standardDeviationIndicator) Calculate(index int) big.Decimal {
	return sdi.indicator.Calculate(index).Sqrt()
}

func (sdi standardDeviationIndicator) LastIndex() int {
	return sdi.indicator.LastIndex()
}

func (sdi standardDeviationIndicator) Key() string {
	return fmt.Sprintf("std:%s", sdi.indicator.Key())
}
