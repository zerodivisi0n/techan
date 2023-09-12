package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type differenceIndicator struct {
	minuend    Indicator
	subtrahend Indicator
}

// NewDifferenceIndicator returns an indicator which returns the difference between one indicator (minuend) and a second
// indicator (subtrahend).
func NewDifferenceIndicator(minuend, subtrahend Indicator) Indicator {
	return NewDifferenceIndicatorWithProxy(DefaultProxy, minuend, subtrahend)
}

func NewDifferenceIndicatorWithProxy(proxy IndicatorProxy, minuend, subtrahend Indicator) Indicator {
	return proxy(differenceIndicator{
		minuend:    minuend,
		subtrahend: subtrahend,
	})
}

func (di differenceIndicator) Calculate(index int) big.Decimal {
	return di.minuend.Calculate(index).Sub(di.subtrahend.Calculate(index))
}

func (di differenceIndicator) LastIndex() int {
	return di.minuend.LastIndex()
}

func (di differenceIndicator) Key() string {
	return fmt.Sprintf("difference(%s,%s)",
		di.minuend.Key(), di.subtrahend.Key())
}
