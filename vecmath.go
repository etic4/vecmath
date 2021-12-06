package vecmath

import (
	"fmt"
	"math"

	"github.com/etic4/vecmath/maths"
)

//Vec2 Représente un vecteur 2D
type Vec2 struct {
	X float64
	Y float64
}

//NewVector crée un nouveau vecteur
func NewVector(x, y float64) Vec2 {
	return Vec2{x, y}
}

//FromPolar retourne un vecteur à partir de coordonées polaires
// param angle: angle en radians
// param length: longueur
func FromPolar(angle float64, length float64) Vec2 {
	x := math.Cos(angle) * length
	y := -math.Sin(angle) * length // y = 0 en haut à gauche

	return NewVector(x, y)
}

//Clone retourne une copie du vecteur
//Comme le vecteur est passé par valeur, 'v' est déjà une copie
func (v Vec2) Clone() Vec2 {
	return v
}

//Equals vérifie si les deux vecteur sont égaux
func (v Vec2) Equals(other Vec2) bool {
	return v.X == other.X && v.Y == other.Y
}

//Angle retourne l'angle en radians entre ce vecteur
// et l'axe des x
func (v Vec2) Angle() float64 {
	angle := math.Atan2(v.Y, v.X)
	if angle < 0 {
		angle += 2 * math.Pi
	}
	return -angle //négatif à cause de la direction de y
}

//Add additionne deux vecteurs
func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v.X + v2.X, v.Y + v2.Y}
}

//AddScalar ajoute 's' à chaque composante
func (v Vec2) AddScalar(s float64) Vec2 {
	return Vec2{v.X + s, v.Y + s}
}

//Sub soustrait deux vecteurs
func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v.X - v2.X, v.Y - v2.Y}
}

//SubScalar soustrait s à chaque composante
func (v Vec2) SubScalar(s float64) Vec2 {
	return Vec2{v.X - s, v.Y - s}
}

//DotProduct produit scalaire de deux vecteurs
func (v Vec2) DotProduct(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

//Mult multiplie le vecteur par un scalaire
func (v Vec2) Mult(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

//Div divise le vecteur par un scalaire
func (v Vec2) Div(s float64) Vec2 {
	return Vec2{v.X / s, v.Y / s}
}

//Neg négation d'un vecteur (180°)
func (v Vec2) Neg() Vec2 {
	return Vec2{-1 * v.X, -1 * v.Y}
}

//Length norme du vecteur
func (v Vec2) Length() float64 {

	return math.Sqrt(v.DotProduct(v))
}

//Normalize retourne vecteur unitaire (normalisé)
func (v Vec2) Normalize() Vec2 {
	length := v.Length()
	if length == 0 {
		return v
	}
	return v.Div(length)
}

//Distance distance entre deux vecteurs
func (v Vec2) Distance(v2 Vec2) float64 {
	return v2.Sub(v).Length()
}

//DistanceSq pour la facilité
func (v Vec2) DistanceSq(v2 Vec2) float64 {
	dx := v2.X - v.X
	dy := v2.Y - v.Y

	return dx*dx + dy*dy
}

//SetMag Définit la taille du vecteur
func (v Vec2) SetMag(lenght float64) Vec2 {
	v = v.Normalize()
	return v.Mult(lenght)
}

//Lerp interpolation linéaire entre ce vecteur et l'autre
func (v Vec2) Lerp(other Vec2, t float64) Vec2 {
	nwX := v.X + t*(other.X-v.X)
	nwY := v.Y + t*(other.Y-v.Y)

	return Vec2{X: nwX, Y: nwY}
}

//ZERO vecteur (0, 0)
func ZERO() Vec2 {
	return NewVector(0, 0)
}

//UP vecteur (0, -1)
func UP() Vec2 {
	return NewVector(0, -1)
}

//DOWN vecteur (0, 1)
func DOWN() Vec2 {
	return NewVector(0, 1)
}

//LEFT vecteur (-1, 0)
func LEFT() Vec2 {
	return NewVector(-1, 0)
}

//RIGHT vecteur (1, 0)
func RIGHT() Vec2 {
	return NewVector(1, 0)
}

//Round arrondis les deux composantes du vecteur à la nième decimale
func (v Vec2) Round(decimals int) Vec2 {
	prop := math.Pow10(decimals)
	return Vec2{math.Round(v.X*prop) / prop, math.Round(v.Y*prop) / prop}
}

//LimitMag limite la longueur du vecteur
func (v Vec2) LimitMag(min float64, max float64) Vec2 {
	clamped := maths.Clamp(v.Length(), min, max)
	return v.SetMag(clamped)
}

//String Retourne imprimable
func (v Vec2) String() string {
	return fmt.Sprintf("%v,%v", v.X, v.Y)
}
