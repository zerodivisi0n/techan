package techan

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
		cache[ind.Key()] = ind
		return ind
	}
}
