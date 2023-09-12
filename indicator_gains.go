package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type gainLossIndicator struct {
	indicator   Indicator
	coefficient big.Decimal
}

// NewGainIndicator returns a derivative indicator that returns the gains in the underlying indicator in the last bar,
// if any. If the delta is negative, zero is returned
func NewGainIndicator(indicator Indicator) Indicator {
	return gainLossIndicator{
		indicator:   indicator,
		coefficient: big.ONE,
	}
}

// NewLossIndicator returns a derivative indicator that returns the losses in the underlying indicator in the last bar,
// if any. If the delta is positive, zero is returned
func NewLossIndicator(indicator Indicator) Indicator {
	return gainLossIndicator{
		indicator:   indicator,
		coefficient: big.ONE.Neg(),
	}
}

func (gli gainLossIndicator) Calculate(index int) big.Decimal {
	if index == 0 {
		return big.ZERO
	}

	curValue := gli.indicator.Calculate(index)
	prevValue := gli.indicator.Calculate(index - 1)
	// gain: curValue < prevValue
	// loss: curValue > prevValue
	if curValue.LT(prevValue) == gli.coefficient.GT(big.ZERO) {
		return big.ZERO
	}

	return curValue.Sub(prevValue).Mul(gli.coefficient)
}

func (gli gainLossIndicator) LastIndex() int {
	return gli.indicator.LastIndex()
}

func (gli gainLossIndicator) Key() string {
	return fmt.Sprintf("gli(%v):%s", gli.coefficient, gli.indicator.Key())
}

type cumulativeIndicator struct {
	indicator Indicator
	window    int
	mult      big.Decimal
}

// NewCumulativeGainsIndicator returns a derivative indicator which returns all gains made in a base indicator for a given
// window.
func NewCumulativeGainsIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		indicator: indicator,
		window:    window,
		mult:      big.ONE,
	}
}

// NewCumulativeLossesIndicator returns a derivative indicator which returns all losses in a base indicator for a given
// window.
func NewCumulativeLossesIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		indicator: indicator,
		window:    window,
		mult:      big.ONE.Neg(),
	}
}

func (ci cumulativeIndicator) Calculate(index int) big.Decimal {
	total := big.NewDecimal(0.0)

	for i := Max(1, index-(ci.window-1)); i <= index; i++ {
		diff := ci.indicator.Calculate(i).Sub(ci.indicator.Calculate(i - 1))
		if diff.Mul(ci.mult).GT(big.ZERO) {
			total = total.Add(diff.Abs())
		}
	}

	return total
}

func (ci cumulativeIndicator) LastIndex() int {
	return ci.indicator.LastIndex()
}

func (ci cumulativeIndicator) Key() string {
	return fmt.Sprintf("cum(%d,%v):%s", ci.window, ci.mult, ci.indicator.Key())
}

type percentChangeIndicator struct {
	indicator Indicator
}

// NewPercentChangeIndicator returns a derivative indicator which returns the percent change (positive or negative)
// made in a base indicator up until the given indicator
func NewPercentChangeIndicator(indicator Indicator) Indicator {
	return percentChangeIndicator{indicator}
}

func (pgi percentChangeIndicator) Calculate(index int) big.Decimal {
	if index == 0 {
		return big.ZERO
	}

	cp := pgi.indicator.Calculate(index)
	cplast := pgi.indicator.Calculate(index - 1)
	return cp.Div(cplast).Sub(big.ONE)
}

func (pgi percentChangeIndicator) LastIndex() int {
	return pgi.indicator.LastIndex()
}

func (pgi percentChangeIndicator) Key() string {
	return fmt.Sprintf("pgi:%s", pgi.indicator.Key())
}
