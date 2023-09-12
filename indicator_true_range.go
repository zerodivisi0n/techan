package techan

import "github.com/sdcoffey/big"

type trueRangeIndicator struct {
	series TimeSeries
}

// NewTrueRangeIndicator returns a base indicator
// which calculates the true range at the current point in time for a series
// https://www.investopedia.com/terms/a/atr.asp
func NewTrueRangeIndicator(series TimeSeries) Indicator {
	return trueRangeIndicator{
		series: series,
	}
}

func (tri trueRangeIndicator) Calculate(index int) big.Decimal {
	if index-1 < 0 {
		return big.ZERO
	}

	previousClose := tri.series.ClosePrice(index - 1)

	trueHigh := big.MaxSlice(tri.series.HighPrice(index), previousClose)
	trueLow := big.MinSlice(tri.series.LowPrice(index), previousClose)

	return trueHigh.Sub(trueLow)
}

func (tri trueRangeIndicator) LastIndex() int {
	return tri.series.LastIndex()
}
