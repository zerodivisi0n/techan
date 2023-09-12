package techan

import (
	"fmt"

	"github.com/sdcoffey/big"
)

// TimeSeries represents type for candle time series
type TimeSeries interface {
	OpenPrice(index int) big.Decimal
	ClosePrice(index int) big.Decimal
	HighPrice(index int) big.Decimal
	LowPrice(index int) big.Decimal
	Volume(index int) big.Decimal
	LastIndex() int
	Key() string
}

// BaseTimeSeries implements TimeSeries using in-memory array of candles
type BaseTimeSeries struct {
	Candles []*Candle
}

var _ TimeSeries = (*BaseTimeSeries)(nil)

// NewBaseTimeSeries returns a new, empty, TimeSeries
func NewBaseTimeSeries() (t *BaseTimeSeries) {
	t = new(BaseTimeSeries)
	t.Candles = make([]*Candle, 0)

	return t
}

// AddCandle adds the given candle to this TimeSeries if it is not nil and after the last candle in this timeseries.
// If the candle is added, AddCandle will return true, otherwise it will return false.
func (ts *BaseTimeSeries) AddCandle(candle *Candle) bool {
	if candle == nil {
		panic(fmt.Errorf("error adding Candle: candle cannot be nil"))
	}

	var lastCandle *Candle
	if lastIdx := ts.LastIndex(); lastIdx >= 0 {
		lastCandle = ts.Candles[lastIdx]
	}
	if lastCandle == nil || candle.Period.Since(lastCandle.Period) >= 0 {
		ts.Candles = append(ts.Candles, candle)
		return true
	}

	return false
}

// OpenPrice returns open price for given index
func (ts *BaseTimeSeries) OpenPrice(index int) big.Decimal {
	return ts.Candles[index].OpenPrice
}

// ClosePrice returns close price for given index
func (ts *BaseTimeSeries) ClosePrice(index int) big.Decimal {
	return ts.Candles[index].ClosePrice
}

// HighPrice returns high price for given index
func (ts *BaseTimeSeries) HighPrice(index int) big.Decimal {
	return ts.Candles[index].MaxPrice
}

// LowPrice returns low price for given index
func (ts *BaseTimeSeries) LowPrice(index int) big.Decimal {
	return ts.Candles[index].MinPrice
}

// Volume returns volume for given index
func (ts *BaseTimeSeries) Volume(index int) big.Decimal {
	return ts.Candles[index].Volume
}

// LastIndex will return the index of the last candle in this series
func (ts *BaseTimeSeries) LastIndex() int {
	return len(ts.Candles) - 1
}

// Key returns description of this timeseries (e.g. candle interval)
func (ts *BaseTimeSeries) Key() string {
	if len(ts.Candles) == 0 {
		return "unknown"
	}
	return ts.Candles[0].Period.End.Sub(ts.Candles[0].Period.Start).String()
}
