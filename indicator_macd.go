package techan

// NewMACDIndicator returns a derivative Indicator which returns the difference between two EMAIndicators with long and
// short windows. It's useful for gauging the strength of price movements. A more in-depth explanation can be found here:
// http://www.investopedia.com/terms/m/macd.asp
func NewMACDIndicator(baseIndicator Indicator, shortwindow, longwindow int) Indicator {
	return NewMACDIndicatorWithProxy(DefaultProxy, baseIndicator, shortwindow, longwindow)
}

func NewMACDIndicatorWithProxy(proxy IndicatorProxy, baseIndicator Indicator, shortwindow, longwindow int) Indicator {
	return NewDifferenceIndicatorWithProxy(proxy,
		proxy(NewEMAIndicator(baseIndicator, shortwindow)),
		proxy(NewEMAIndicator(baseIndicator, longwindow)),
	)
}

// NewMACDHistogramIndicator returns a derivative Indicator based on the MACDIndicator, the result of which is
// the macd indicator minus it's signalLinewindow EMA. A more in-depth explanation can be found here:
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:macd-histogram
func NewMACDHistogramIndicator(macdIdicator Indicator, signalLinewindow int) Indicator {
	return NewMACDHistogramIndicatorWithProxy(DefaultProxy, macdIdicator, signalLinewindow)
}

func NewMACDHistogramIndicatorWithProxy(proxy IndicatorProxy, macdIdicator Indicator, signalLinewindow int) Indicator {
	return NewDifferenceIndicatorWithProxy(proxy, macdIdicator, proxy(NewEMAIndicator(macdIdicator, signalLinewindow)))
}
