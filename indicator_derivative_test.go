package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDerivativeIndicator(t *testing.T) {
	series := mockTimeSeries("1", "1", "2", "3", "5", "8", "13")
	indicator := DerivativeIndicator{
		Indicator: NewClosePriceIndicator(series),
	}

	t.Run("returns zero at index zero", func(t *testing.T) {
		assert.EqualValues(t, "0", indicator.Calculate(0).String())
	})

	t.Run("returns the derivative", func(t *testing.T) {
		assert.EqualValues(t, "0", indicator.Calculate(1).String())

		for i := 2; i <= series.LastIndex(); i++ {
			expected := series.ClosePrice(i - 2)

			assert.EqualValues(t, expected.String(), indicator.Calculate(i).String())
		}
	})
}
