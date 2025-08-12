package utilities

import "math"

func RadiansToDegrees(rad float64) float64 {
	return rad * (180 / math.Pi)
}

func DegreesToRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}
