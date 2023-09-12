package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type commidityChannelIndexIndicator struct {
	typicalPrice    Indicator
	typicalPriceSma Indicator
	meanDeviation   Indicator
	window          int
}

// NewCCIIndicator Returns a new Commodity Channel Index Indicator
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:commodity_channel_index_cci
func NewCCIIndicator(ts TimeSeries, window int) Indicator {
	typicalPrice := NewTypicalPriceIndicator(ts)
	typicalPriceSma := NewSimpleMovingAverage(typicalPrice, window)
	meanDeviation := NewMeanDeviationIndicator(NewClosePriceIndicator(ts), window)
	return commidityChannelIndexIndicator{
		typicalPrice:    typicalPrice,
		typicalPriceSma: typicalPriceSma,
		meanDeviation:   meanDeviation,
		window:          window,
	}
}

func (ccii commidityChannelIndexIndicator) Calculate(index int) big.Decimal {
	return ccii.typicalPrice.Calculate(index).Sub(ccii.typicalPriceSma.Calculate(index)).Div(ccii.meanDeviation.Calculate(index).Mul(big.NewFromString("0.015")))
}

func (ccii commidityChannelIndexIndicator) LastIndex() int {
	return ccii.typicalPrice.LastIndex()
}

func (ccii commidityChannelIndexIndicator) Key() string {
	return fmt.Sprintf("ccii(%d):%s", ccii.window, ccii.typicalPrice.Key())
}
