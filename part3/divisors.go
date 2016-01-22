package part3

import (
	"math"
)

func Divisors(product uint) []uint {
	lows := []uint{}

	fastDivision := float64(product)
	max := math.Sqrt(fastDivision)

	for i := 2.0; i < max; i++ {
		if math.Mod(fastDivision, i) == 0 {
			lows = append(lows, uint(i))
		}
	}

	highs := make([]uint, len(lows))
	for j, i := 0, len(lows)-1; i >= 0; i, j = i-1, j+1 {
		highs[j] = product / lows[i]
	}

	if math.Mod(fastDivision, max) == 0 {
		lows = append(lows, uint(max))
	}
	return append(lows, highs...)
}
