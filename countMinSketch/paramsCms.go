package cms

import "math"

func CalculateMC(epsilon float64) uint {
	return uint(math.Ceil(math.E / epsilon))
}

func CalculateKC(delta float64) uint {
	return uint(math.Ceil(math.Log(math.E / delta)))
}
