package wotoMath

import "math"

func NewVector2(x float64, y float64) *Vector2 {
	return &Vector2{X: x, Y: y}
}

func Dot(ihs *Vector2, rhs *Vector2) float64 {
	return ihs.X*rhs.X + ihs.Y*rhs.Y
}

func Lerp(a *Vector2, b *Vector2, t float64) *Vector2 {
	return NewVector2(
		a.X+(b.X-a.X)*t,
		a.Y+(b.Y-a.Y)*t,
	)
}

func Distance(a *Vector2, b *Vector2) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func Reflect(ihs *Vector2, rhs *Vector2) *Vector2 {
	factor := -2.0 * Dot(ihs, rhs)
	return NewVector2(
		factor*ihs.X+rhs.X,
		factor*ihs.Y+rhs.Y,
	)
}
