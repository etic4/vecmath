package maths

import (
	"math"
)

//Rad retourne la valeur en radian de l'angle passé en degrés
func Rad(deg float64) float64 {
	return deg * math.Pi / 180
}

//Deg retourne la valeur en degrés de l'angle passé en radians
func Deg(rad float64) float64 {
	return rad * 180 / math.Pi
}

//Clamp restreint 'val' entre min et max (float64)
func Clamp(val float64, min float64, max float64) float64 {
	return math.Max(min, math.Min(val, max))
}

//Lerp interpolation linéaire entre first et second selon t
func Lerp(first float64, second float64, t float64) float64 {
	return first + t*(second-first)
}
