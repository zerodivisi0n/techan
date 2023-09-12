package techan

import (
	"github.com/sdcoffey/big"
)

// IndicatorProxy is a wrapper type to add some additional logic to indicators
type IndicatorProxy func(Indicator) Indicator

// DefaultProxy is a default proxy for indicators without any logic
var DefaultProxy = func(ind Indicator) Indicator { return ind }

// NewCachedIndicatorProxy returns indicator proxy with cache for indicators
// If indicator already exists in cache it is used
func NewCachedIndicatorProxy() IndicatorProxy {
	cache := make(map[string]Indicator)

	return func(ind Indicator) Indicator {
		if cached, ok := cache[ind.Key()]; ok {
			return cached
		}
		if _, ok := ind.(cachedIndicator); !ok {
			ind = newCachedIndicatorWrapper(ind)
		}
		cache[ind.Key()] = ind
		return ind
	}
}

type cachedIndicatorWrapper struct {
	indicator   Indicator
	resultCache resultCache
}

func newCachedIndicatorWrapper(indicator Indicator) Indicator {
	return &cachedIndicatorWrapper{
		indicator:   indicator,
		resultCache: make([]*big.Decimal, 1000),
	}
}

func (ciw *cachedIndicatorWrapper) Calculate(index int) big.Decimal {
	if cachedValue := returnIfCached(ciw, index, nil); cachedValue != nil {
		return *cachedValue
	}

	result := ciw.indicator.Calculate(index)
	if index != ciw.indicator.LastIndex() {
		cacheResult(ciw, index, result)
	}

	return result
}

func (ciw cachedIndicatorWrapper) LastIndex() int {
	return ciw.indicator.LastIndex()
}

func (ciw cachedIndicatorWrapper) Key() string {
	return ciw.indicator.Key()
}

func (ciw cachedIndicatorWrapper) cache() resultCache { return ciw.resultCache }

func (ciw *cachedIndicatorWrapper) setCache(newCache resultCache) {
	ciw.resultCache = newCache
}

func (ciw cachedIndicatorWrapper) windowSize() int { return 0 }
