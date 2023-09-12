package techan

import (
	"fmt"
	"math"

	"github.com/sdcoffey/big"
)

type constantIndicator float64

// NewConstantIndicator returns an indicator which always returns the same value for any index. It's useful when combined
// with other, fluxuating indicators to determine when an indicator has crossed a threshold.
func NewConstantIndicator(constant float64) Indicator {
	return constantIndicator(constant)
}

func (ci constantIndicator) Calculate(index int) big.Decimal {
	return big.NewDecimal(float64(ci))
}

func (ci constantIndicator) LastIndex() int {
	return math.MaxInt
}

func (ci constantIndicator) Key() string {
	return fmt.Sprintf("constant(%f)", float64(ci))
}
