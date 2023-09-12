package techan

import (
	"fmt"
	"strings"

	"github.com/sdcoffey/big"
)

type fixedIndicator []float64

// NewFixedIndicator returns an indicator with a fixed set of values that are returned when an index is passed in
func NewFixedIndicator(vals ...float64) Indicator {
	return fixedIndicator(vals)
}

func (fi fixedIndicator) Calculate(index int) big.Decimal {
	return big.NewDecimal(fi[index])
}

func (fi fixedIndicator) LastIndex() int {
	return len(fi) - 1
}

func (fi fixedIndicator) Key() string {
	vals := make([]string, 0, len(fi))
	for _, f := range fi {
		vals = append(vals, fmt.Sprintf("%f", f))
	}
	return fmt.Sprintf("fixed(%d:%s)", len(fi), strings.Join(vals, ","))
}
