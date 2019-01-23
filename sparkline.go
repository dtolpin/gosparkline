// Sparklines
package spark // import "bitbucket.org/dtolpin/sparkline"

import (
	"math"
)

var levels = []rune("▁▂▃▄▅▆▇█")

// Line generates a sparkline string from a slice of float64.
func Line(ys []float64) string {
	if len(ys) == 0 {
		return ""
	}
	min := math.Inf(1)
	max := math.Inf(-1)
	for _, y := range ys {
		if y < min {
			min = y
		}
		if y > max {
			max = y
		}
	}
	if max == min {
		max = min + 1
	}
	line := make([]rune, len(ys))
	for i := range ys {
		j := int(math.Floor(
			(float64(len(levels)) - 1.) *
				(ys[i] - min) / (max - min)))
		line[i] = levels[j]
	}
	return string(line)
}
