package spark

import (
	"testing"
)

func TestLine(t *testing.T) {
	for _, c := range []struct {
		name   string
		ys     []float64
		length int
	}{
		{
			"empty",
			[]float64{},
			0,
		},
		// Kicking tyres
		{
			"regular",
			[]float64{1, 3, 2},
			3,
		},
		{
			"single element",
			[]float64{1},
			1,
		},
		{
			"constant series",
			[]float64{3, 3, 3},
			3,
		},
	} {
		line := Line(c.ys)
		length := len([]rune(line))
		if length != c.length {
			t.Errorf("%s: expected sparkline of length %d, "+
				"got %q of length %d",
				c.name, c.length, line, length)
		}
	}
}
