package techan

import "testing"

func TestIndicatorTimeSeries(t *testing.T) {
	indicatorTimeseries := NewIndicatorTimeSeries(mockedTimeSeries, func(indicator Indicator) Indicator {
		return NewSimpleMovingAverage(indicator, 3)
	})

	expectedOpenCloseVolumeValues := []float64{
		0,
		0,
		64.09,
		63.75,
		63.67,
		63.49,
		63.55,
		63.65,
		63.57,
		63.39,
		62.55,
		62.07,
	}

	indicatorEquals(t, expectedOpenCloseVolumeValues, NewOpenPriceIndicator(indicatorTimeseries))
	indicatorEquals(t, expectedOpenCloseVolumeValues, NewClosePriceIndicator(indicatorTimeseries))
	indicatorEquals(t, expectedOpenCloseVolumeValues, NewVolumeIndicator(indicatorTimeseries))
	indicatorEquals(t, []float64{0, 0, 65.09, 64.75, 64.67, 64.49, 64.55, 64.65, 64.57, 64.39, 63.55, 63.07},
		NewHighPriceIndicator(indicatorTimeseries))
	indicatorEquals(t, []float64{0, 0, 63.09, 62.75, 62.67, 62.49, 62.55, 62.65, 62.57, 62.39, 61.55, 61.07},
		NewLowPriceIndicator(indicatorTimeseries))
}
