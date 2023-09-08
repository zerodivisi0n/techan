package techan

import "github.com/sdcoffey/big"

type modifiedMovingAverageIndicator struct {
	indicator   Indicator
	window      int
	resultCache resultCache
}

// NewMMAIndicator returns a derivative indciator which returns the modified moving average of the underlying
// indictator. An in-depth explanation can be found here:
// https://en.wikipedia.org/wiki/Moving_average#Modified_moving_average
func NewMMAIndicator(indicator Indicator, window int) Indicator {
	return &modifiedMovingAverageIndicator{
		indicator:   indicator,
		window:      window,
		resultCache: make([]*big.Decimal, 10000),
	}
}

func (mma *modifiedMovingAverageIndicator) Calculate(index int) big.Decimal {
	return mma.calculate(index, false)
}

func (mma *modifiedMovingAverageIndicator) calculate(index int, allowCache bool) big.Decimal {
	if cachedValue := returnIfCached(mma, index, allowCache, func(i int) big.Decimal {
		return NewSimpleMovingAverage(mma.indicator, mma.window).Calculate(i)
	}); cachedValue != nil {
		return *cachedValue
	}

	todayVal := mma.indicator.Calculate(index)
	lastVal := mma.calculate(index-1, true)

	result := lastVal.Add(big.NewDecimal(1.0 / float64(mma.window)).Mul(todayVal.Sub(lastVal)))
	if allowCache {
		cacheResult(mma, index, result)
	}

	return result
}

func (mma modifiedMovingAverageIndicator) cache() resultCache {
	return mma.resultCache
}

func (mma *modifiedMovingAverageIndicator) setCache(cache resultCache) {
	mma.resultCache = cache
}

func (mma modifiedMovingAverageIndicator) windowSize() int {
	return mma.window
}
