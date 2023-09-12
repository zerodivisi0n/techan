package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

type keltnerChannelIndicator struct {
	ema    Indicator
	atr    Indicator
	mul    big.Decimal
	window int
}

// NewKeltnerChannelUpperIndicator returns Indicator which represents Keltner Channel Upper Band for the given window.
// https://www.investopedia.com/terms/k/keltnerchannel.asp
func NewKeltnerChannelUpperIndicator(series TimeSeries, window int) Indicator {
	return keltnerChannelIndicator{
		atr:    NewAverageTrueRangeIndicator(series, window),
		ema:    NewEMAIndicator(NewClosePriceIndicator(series), window),
		mul:    big.ONE,
		window: window,
	}
}

// NewKeltnerChannelLowerIndicator returns Indicator which represents Keltner Channel Lower Band for the given window.
// https://www.investopedia.com/terms/k/keltnerchannel.asp
func NewKeltnerChannelLowerIndicator(series TimeSeries, window int) Indicator {
	return keltnerChannelIndicator{
		atr:    NewAverageTrueRangeIndicator(series, window),
		ema:    NewEMAIndicator(NewClosePriceIndicator(series), window),
		mul:    big.ONE.Neg(),
		window: window,
	}
}

func (kci keltnerChannelIndicator) Calculate(index int) big.Decimal {
	if index <= kci.window-1 {
		return big.ZERO
	}

	coefficient := big.NewFromInt(2).Mul(kci.mul)

	return kci.ema.Calculate(index).Add(kci.atr.Calculate(index).Mul(coefficient))
}

func (kci keltnerChannelIndicator) LastIndex() int {
	return kci.atr.LastIndex()
}

func (kci keltnerChannelIndicator) Key() string {
	return fmt.Sprintf("kci(%d,%v):%s", kci.window, kci.mul, kci.atr.Key())
}
