package simple

import (
	"math"
)

// TakePart - gives the nearest whole greater number from total by the percentage.
func TakePart(percentage float64, total int64) (part int64) {
	if percentage < 0 {
		percentage = 0
	}
	if percentage > 100 {
		percentage = 100
	}
	return int64(math.Ceil(percentage * float64(total) / 100))
}
