package techan

import (
	"fmt"
	"math"

	"github.com/sdcoffey/big"
)

type relativeStrengthIndexIndicator struct {
	rsIndicator Indicator
	oneHundred  big.Decimal
}

// NewRelativeStrengthIndexIndicator returns a derivative Indicator which returns the relative strength index of the base indicator
// in a given time frame. A more in-depth explanation of relative strength index can be found here:
// https://www.investopedia.com/terms/r/rsi.asp
func NewRelativeStrengthIndexIndicator(indicator Indicator, timeframe int) Indicator {
	return NewRelativeStrengthIndexIndicatorWithProxy(DefaultProxy, indicator, timeframe)
}

func NewRelativeStrengthIndexIndicatorWithProxy(proxy IndicatorProxy, indicator Indicator, timeframe int) Indicator {
	return proxy(relativeStrengthIndexIndicator{
		rsIndicator: NewRelativeStrengthIndicatorWithProxy(proxy, indicator, timeframe),
		oneHundred:  big.NewFromString("100"),
	})
}

func (rsi relativeStrengthIndexIndicator) Calculate(index int) big.Decimal {
	relativeStrength := rsi.rsIndicator.Calculate(index)

	return rsi.oneHundred.Sub(rsi.oneHundred.Div(big.ONE.Add(relativeStrength)))
}

func (rsi relativeStrengthIndexIndicator) LastIndex() int {
	return rsi.rsIndicator.LastIndex()
}

func (rsi relativeStrengthIndexIndicator) Key() string {
	return fmt.Sprintf("rsii:%s", rsi.rsIndicator.Key())
}

type relativeStrengthIndicator struct {
	avgGain Indicator
	avgLoss Indicator
	window  int
}

// NewRelativeStrengthIndicator returns a derivative Indicator which returns the relative strength of the base indicator
// in a given time frame. Relative strength is the average again of up periods during the time frame divided by the
// average loss of down period during the same time frame
func NewRelativeStrengthIndicator(indicator Indicator, timeframe int) Indicator {
	return NewRelativeStrengthIndicatorWithProxy(DefaultProxy, indicator, timeframe)
}

func NewRelativeStrengthIndicatorWithProxy(proxy IndicatorProxy, indicator Indicator, timeframe int) Indicator {
	return proxy(relativeStrengthIndicator{
		avgGain: proxy(NewMMAIndicator(proxy(NewGainIndicator(indicator)), timeframe)),
		avgLoss: proxy(NewMMAIndicator(proxy(NewLossIndicator(indicator)), timeframe)),
		window:  timeframe,
	})
}

func (rs relativeStrengthIndicator) Calculate(index int) big.Decimal {
	if index < rs.window-1 {
		return big.ZERO
	}

	avgGain := rs.avgGain.Calculate(index)
	avgLoss := rs.avgLoss.Calculate(index)

	if avgLoss.IsZero() {
		return big.NewDecimal(math.Inf(1))
	}

	return avgGain.Div(avgLoss)
}

func (rs relativeStrengthIndicator) LastIndex() int {
	return rs.avgGain.LastIndex()
}

func (rs relativeStrengthIndicator) Key() string {
	return fmt.Sprintf("rsi(%d):%s", rs.window, rs.avgGain.Key())
}
