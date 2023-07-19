package math

import (
	"math"
)

func cal(b1, b2, a2, a3, gap float64) []float64 {
	alpha := math.Atan(b2 / a2)
	a1 := b1 / math.Tan(alpha)
	a := a1 + a2 + a3
	b := a * math.Tan(alpha)
	c := math.Sqrt(a*a + b*b)

	dif := gap*math.Tan(alpha) + gap/math.Tan(alpha)

	return []float64{c, dif}
}
