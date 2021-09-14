// wotoMath Project
// Copyright (C) 2021 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMath

import (
	"fmt"
	"math"
)

//---------------------------------------------------------

func (v *Vector2) IsLarger(vect Vector2) bool {
	return v.X > vect.X || v.Y > vect.Y
}

func (v *Vector2) IsLargerXY(vect Vector2) bool {
	return v.X > vect.X && v.Y > vect.Y
}

func (v *Vector2) IsBiggerX(vect Vector2) bool {
	return v.X > vect.X
}

func (v *Vector2) IsBiggerY(vect Vector2) bool {
	return v.Y > vect.Y
}

func (v *Vector2) IsLargerOrEqual(vect Vector2) bool {
	return v.X >= vect.X || v.Y >= vect.Y
}

func (v *Vector2) IsLargerOrEqualXY(vect Vector2) bool {
	return v.X >= vect.X && v.Y >= vect.Y
}

func (v *Vector2) IsBiggerOrEqualX(vect Vector2) bool {
	return v.X >= vect.X
}

func (v *Vector2) IsBiggerOrEqualY(vect Vector2) bool {
	return v.Y >= vect.Y
}

func (v *Vector2) IsSmaller(vect Vector2) bool {
	return v.X < vect.X || v.Y < vect.Y
}

func (v *Vector2) IsSmallerXY(vect Vector2) bool {
	return v.X < vect.X && v.Y < vect.Y
}

func (v *Vector2) IsSmallerX(vect Vector2) bool {
	return v.X < vect.X
}

func (v *Vector2) IsSmallerY(vect Vector2) bool {
	return v.Y < vect.Y
}

func (v *Vector2) IsSmallerOrEqual(vect Vector2) bool {
	return v.X <= vect.X || v.Y <= vect.Y
}

func (v *Vector2) IsSmallerOrEqualXY(vect Vector2) bool {
	return v.X <= vect.X && v.Y <= vect.Y
}

func (v *Vector2) IsSmallerOrEqualX(vect Vector2) bool {
	return v.X <= vect.X
}

func (v *Vector2) IsSmallerOrEqualY(vect Vector2) bool {
	return v.Y <= vect.Y
}

func (v *Vector2) Add(other *Vector2) *Vector2 {
	if other == nil {
		return v
	}

	return NewVector2(v.X+other.X, v.Y+other.Y)
}

func (v *Vector2) AddThis(other *Vector2) *Vector2 {
	if other == nil {
		return v
	}

	v.X += other.X
	v.Y += other.Y
	return v
}

func (v *Vector2) Sub(other *Vector2) *Vector2 {
	if other == nil {
		return v
	}

	return NewVector2(v.X-other.X, v.Y-other.Y)
}

func (v *Vector2) SubThis(other *Vector2) *Vector2 {
	if other == nil {
		return v
	}

	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v *Vector2) Mul(other *Vector2) *Vector2 {
	if other == nil {
		return v
	}

	return NewVector2(v.X*other.X, v.Y*other.Y)
}

func (v *Vector2) MulThis(other *Vector2) *Vector2 {
	if other == nil {
		return v
	}

	v.X *= other.X
	v.Y *= other.Y
	return v
}

func (v *Vector2) Div(other *Vector2) *Vector2 {
	if other == nil {
		return v
	}

	if other.X == 0 && other.Y == 0 {
		return NewVector2(v.X, v.Y)
	} else if other.X == 0 {
		return NewVector2(v.X, v.Y/other.Y)
	} else if other.Y == 0 {
		return NewVector2(v.X/other.X, v.Y)
	}

	return NewVector2(v.X/other.X, v.Y/other.Y)
}

func (v *Vector2) DivThis(other *Vector2) *Vector2 {
	if other == nil {
		return v
	}

	if other.X != 0 {
		v.X /= other.X
	}
	if other.Y != 0 {
		v.Y /= other.Y
	}

	return v
}

func (v *Vector2) Copy() *Vector2 {
	return NewVector2(v.X, v.Y)
}

func (v *Vector2) Set(x float64, y float64) *Vector2 {
	v.X = x
	v.Y = y
	return v
}

func (v *Vector2) SetX(x float64) *Vector2 {
	v.X = x
	return v
}

func (v *Vector2) SetY(y float64) *Vector2 {
	v.Y = y
	return v
}

func (v *Vector2) AddScalar(scalar float64) *Vector2 {
	return NewVector2(v.X+scalar, v.Y+scalar)
}

func (v *Vector2) AddScalarThis(scalar float64) *Vector2 {
	v.X += scalar
	v.Y += scalar
	return v
}

func (v *Vector2) AddScalars(x, y float64) *Vector2 {
	return NewVector2(v.X+x, v.Y+y)
}

func (v *Vector2) AddScalarsThis(x, y float64) *Vector2 {
	v.X += x
	v.Y += y
	return v
}

func (v *Vector2) SubScalar(scalar float64) *Vector2 {
	return NewVector2(v.X-scalar, v.Y-scalar)
}

func (v *Vector2) SubScalarThis(scalar float64) *Vector2 {
	v.X -= scalar
	v.Y -= scalar
	return v
}

func (v *Vector2) SubScalars(x, y float64) *Vector2 {
	return NewVector2(v.X-x, v.Y-y)
}

func (v *Vector2) SubScalarsThis(x, y float64) *Vector2 {
	v.X -= x
	v.Y -= y
	return v
}

func (v *Vector2) MulScalar(scalar float64) *Vector2 {
	return NewVector2(v.X*scalar, v.Y*scalar)
}

func (v *Vector2) MulScalarThis(scalar float64) *Vector2 {
	v.X *= scalar
	v.Y *= scalar
	return v
}

func (v *Vector2) MulScalars(x, y float64) *Vector2 {
	return NewVector2(v.X-x, v.Y-y)
}

func (v *Vector2) MulScalarsThis(x, y float64) *Vector2 {
	v.X -= x
	v.Y -= y
	return v
}

func (v *Vector2) DivScalar(scalar float64) *Vector2 {
	if scalar == 0 {
		return v
	}

	return NewVector2(v.X*scalar, v.Y*scalar)
}

func (v *Vector2) DivScalarThis(scalar float64) *Vector2 {
	if scalar == 0 {
		return v
	}

	v.X /= scalar
	v.Y /= scalar
	return v
}

func (v *Vector2) DivScalars(x, y float64) *Vector2 {
	if x == 0 && y == 0 {
		return v
	} else if x == 0 {
		return NewVector2(v.X, v.Y/y)
	} else if y == 0 {
		return NewVector2(v.X/x, v.Y)
	}

	return NewVector2(v.X/x, v.Y/y)
}

func (v *Vector2) DivScalarsThis(x, y float64) *Vector2 {
	if x != 0 {
		v.X /= x
	}
	if y != 0 {
		v.Y /= y
	}

	return v
}

func (v *Vector2) Distance(other *Vector2) float64 {
	if other == nil {
		return 0
	}

	dx := v.X - other.X
	dy := v.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (v *Vector2) Dot(other *Vector2) float64 {
	if other == nil {
		return 0
	}

	return v.X*other.X + v.Y*other.Y
}

func (v *Vector2) Lerp(other *Vector2, t float64) *Vector2 {
	if other == nil {
		return v
	}

	return NewVector2(
		v.X+(other.X-v.X)*t,
		v.Y+(other.Y-v.Y)*t,
	)
}

func (v *Vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector2) Normalize() *Vector2 {
	m := v.Magnitude()

	if m > Epsilon {
		return v.DivScalar(m)
	} else {
		return v.Copy()
	}
}

func (v *Vector2) Reflect(other *Vector2) *Vector2 {
	if other == nil {
		return v
	}

	factor := -2.0 * v.Dot(other)
	return NewVector2(
		factor*v.X+other.X,
		factor*v.Y+other.Y,
	)
}

func (v *Vector2) Equals(other *Vector2) bool {
	if other == nil {
		return false
	}

	return EqualFloats64(v.X, other.X) && EqualFloats64(v.Y, other.Y)
}

func (v *Vector2) EqualsX(other *Vector2) bool {
	if other == nil {
		return false
	}

	return EqualFloats64(v.X, other.X)
}

func (v *Vector2) EqualsScalarX(x float64) bool {
	return EqualFloats64(v.X, x)
}

func (v *Vector2) EqualsY(other *Vector2) bool {
	if other == nil {
		return false
	}

	return EqualFloats64(v.Y, other.Y)
}

func (v *Vector2) EqualsScalarY(y float64) bool {
	return EqualFloats64(v.Y, y)
}

func (v *Vector2) Serialize() string {
	//TODO
	return fmt.Sprintf("Vector2(%f, %f)", v.X, v.Y)
}

func (v *Vector2) ToString() string {
	return fmt.Sprintf("Vector2(%f, %f)", v.X, v.Y)
}

//---------------------------------------------------------

// Contains method will check and see if `rect` argument is in its
// region or not (even if it's only 0.1 point, it will return true).
func (r *Rectangle) Contains(rect Rectangle) bool {

	return false
}

// Encloses method will check and see if
func (r *Rectangle) Encloses(rect Rectangle) bool {

	return false
}

func (r *Rectangle) IsLarger(rect Rectangle) bool {
	return false
}

func (r *Rectangle) IsLargerOrEqual(rect Rectangle) bool {
	return false
}

func (r *Rectangle) IsSmaller(rect Rectangle) bool {
	return false
}

func (r *Rectangle) IsSmallerOrEqual(rect Rectangle) bool {
	return false
}

//---------------------------------------------------------

func (v *Vector2Int) IsLarger(vect Vector2Int) bool {
	return v.X > vect.X || v.Y > vect.Y
}

func (v *Vector2Int) IsLargerXY(vect Vector2Int) bool {
	return v.X > vect.X && v.Y > vect.Y
}

func (v *Vector2Int) IsBiggerX(vect Vector2Int) bool {
	return v.X > vect.X
}

func (v *Vector2Int) IsBiggerY(vect Vector2Int) bool {
	return v.Y > vect.Y
}

func (v *Vector2Int) IsLargerOrEqual(vect Vector2Int) bool {
	return v.X >= vect.X || v.Y >= vect.Y
}

func (v *Vector2Int) IsLargerOrEqualXY(vect Vector2Int) bool {
	return v.X >= vect.X && v.Y >= vect.Y
}

func (v *Vector2Int) IsBiggerOrEqualX(vect Vector2Int) bool {
	return v.X >= vect.X
}

func (v *Vector2Int) IsBiggerOrEqualY(vect Vector2Int) bool {
	return v.Y >= vect.Y
}

func (v *Vector2Int) IsSmaller(vect Vector2Int) bool {
	return v.X < vect.X || v.Y < vect.Y
}

func (v *Vector2Int) IsSmallerXY(vect Vector2Int) bool {
	return v.X < vect.X && v.Y < vect.Y
}

func (v *Vector2Int) IsSmallerX(vect Vector2Int) bool {
	return v.X < vect.X
}

func (v *Vector2Int) IsSmallerY(vect Vector2Int) bool {
	return v.Y < vect.Y
}

func (v *Vector2Int) IsSmallerOrEqual(vect Vector2Int) bool {
	return v.X <= vect.X || v.Y <= vect.Y
}

func (v *Vector2Int) IsSmallerOrEqualXY(vect Vector2Int) bool {
	return v.X <= vect.X && v.Y <= vect.Y
}

func (v *Vector2Int) IsSmallerOrEqualX(vect Vector2Int) bool {
	return v.X <= vect.X
}

func (v *Vector2Int) IsSmallerOrEqualY(vect Vector2Int) bool {
	return v.Y <= vect.Y
}

func (v *Vector2Int) Add(other *Vector2Int) *Vector2Int {
	if other == nil {
		return v
	}

	return NewVector2Int(v.X+other.X, v.Y+other.Y)
}

func (v *Vector2Int) AddThis(other *Vector2Int) *Vector2Int {
	if other == nil {
		return v
	}

	v.X += other.X
	v.Y += other.Y
	return v
}

func (v *Vector2Int) Sub(other *Vector2Int) *Vector2Int {
	if other == nil {
		return v
	}

	return NewVector2Int(v.X-other.X, v.Y-other.Y)
}

func (v *Vector2Int) SubThis(other *Vector2Int) *Vector2Int {
	if other == nil {
		return v
	}

	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v *Vector2Int) Mul(other *Vector2Int) *Vector2Int {
	if other == nil {
		return v
	}

	return NewVector2Int(v.X*other.X, v.Y*other.Y)
}

func (v *Vector2Int) MulThis(other *Vector2Int) *Vector2Int {
	if other == nil {
		return v
	}

	v.X *= other.X
	v.Y *= other.Y
	return v
}

func (v *Vector2Int) Div(other *Vector2Int) *Vector2Int {
	if other == nil {
		return v
	}

	if other.X == 0 && other.Y == 0 {
		return NewVector2Int(v.X, v.Y)
	} else if other.X == 0 {
		return NewVector2Int(v.X, v.Y/other.Y)
	} else if other.Y == 0 {
		return NewVector2Int(v.X/other.X, v.Y)
	}

	return NewVector2Int(v.X/other.X, v.Y/other.Y)
}

func (v *Vector2Int) DivThis(other *Vector2Int) *Vector2Int {
	if other == nil {
		return v
	}

	if other.X != 0 {
		v.X /= other.X
	}
	if other.Y != 0 {
		v.Y /= other.Y
	}

	return v
}

func (v *Vector2Int) Copy() *Vector2Int {
	return NewVector2Int(v.X, v.Y)
}

func (v *Vector2Int) Set(x int, y int) *Vector2Int {
	v.X = x
	v.Y = y
	return v
}

func (v *Vector2Int) SetX(x int) *Vector2Int {
	v.X = x
	return v
}

func (v *Vector2Int) SetY(y int) *Vector2Int {
	v.Y = y
	return v
}

func (v *Vector2Int) AddScalar(scalar int) *Vector2Int {
	return NewVector2Int(v.X+scalar, v.Y+scalar)
}

func (v *Vector2Int) AddScalarThis(scalar int) *Vector2Int {
	v.X += scalar
	v.Y += scalar
	return v
}

func (v *Vector2Int) AddScalars(x, y int) *Vector2Int {
	return NewVector2Int(v.X+x, v.Y+y)
}

func (v *Vector2Int) AddScalarsThis(x, y int) *Vector2Int {
	v.X += x
	v.Y += y
	return v
}

func (v *Vector2Int) SubScalar(scalar int) *Vector2Int {
	return NewVector2Int(v.X-scalar, v.Y-scalar)
}

func (v *Vector2Int) SubScalarThis(scalar int) *Vector2Int {
	v.X -= scalar
	v.Y -= scalar
	return v
}

func (v *Vector2Int) SubScalars(x, y int) *Vector2Int {
	return NewVector2Int(v.X-x, v.Y-y)
}

func (v *Vector2Int) SubScalarsThis(x, y int) *Vector2Int {
	v.X -= x
	v.Y -= y
	return v
}

func (v *Vector2Int) MulScalar(scalar int) *Vector2Int {
	return NewVector2Int(v.X*scalar, v.Y*scalar)
}

func (v *Vector2Int) MulScalarThis(scalar int) *Vector2Int {
	v.X *= scalar
	v.Y *= scalar
	return v
}

func (v *Vector2Int) MulScalars(x, y int) *Vector2Int {
	return NewVector2Int(v.X-x, v.Y-y)
}

func (v *Vector2Int) MulScalarsThis(x, y int) *Vector2Int {
	v.X -= x
	v.Y -= y
	return v
}

func (v *Vector2Int) DivScalar(scalar int) *Vector2Int {
	if scalar == 0 {
		return v
	}

	return NewVector2Int(v.X*scalar, v.Y*scalar)
}

func (v *Vector2Int) DivScalarThis(scalar int) *Vector2Int {
	if scalar == 0 {
		return v
	}

	v.X /= scalar
	v.Y /= scalar
	return v
}

func (v *Vector2Int) DivScalars(x, y int) *Vector2Int {
	if x == 0 && y == 0 {
		return v
	} else if x == 0 {
		return NewVector2Int(v.X, v.Y/y)
	} else if y == 0 {
		return NewVector2Int(v.X/x, v.Y)
	}

	return NewVector2Int(v.X/x, v.Y/y)
}

func (v *Vector2Int) DivScalarsThis(x, y int) *Vector2Int {
	if x != 0 {
		v.X /= x
	}
	if y != 0 {
		v.Y /= y
	}

	return v
}

func (v *Vector2Int) Distance(other *Vector2Int) int {
	if other == nil {
		return 0
	}

	dx := v.X - other.X
	dy := v.Y - other.Y
	return int(math.Sqrt(float64(dx*dx + dy*dy)))
}

func (v *Vector2Int) Dot(other *Vector2Int) int {
	if other == nil {
		return 0
	}

	return v.X*other.X + v.Y*other.Y
}

func (v *Vector2Int) Lerp(other *Vector2Int, t int) *Vector2Int {
	if other == nil {
		return v
	}

	return NewVector2Int(
		v.X+(other.X-v.X)*t,
		v.Y+(other.Y-v.Y)*t,
	)
}

func (v *Vector2Int) Magnitude() int {
	return int(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v *Vector2Int) Normalize() *Vector2Int {
	m := v.Magnitude()

	if float64(m) > Epsilon {
		return v.DivScalar(m)
	} else {
		return v.Copy()
	}
}

func (v *Vector2Int) Reflect(other *Vector2Int) *Vector2Int {
	if other == nil {
		return v
	}

	factor := -2.0 * v.Dot(other)
	return NewVector2Int(
		factor*v.X+other.X,
		factor*v.Y+other.Y,
	)
}

func (v *Vector2Int) Equals(other *Vector2Int) bool {
	if other == nil {
		return false
	}

	return v.X == other.X && v.Y == other.Y
}

func (v *Vector2Int) EqualsX(other *Vector2Int) bool {
	if other == nil {
		return false
	}

	return v.X == other.X
}

func (v *Vector2Int) EqualsScalarX(x int) bool {
	return v.X == x
}

func (v *Vector2Int) EqualsY(other *Vector2Int) bool {
	if other == nil {
		return false
	}

	return v.Y == other.Y
}

func (v *Vector2Int) EqualsScalarY(y int) bool {
	return v.Y == y
	//return EqualFloats64(v.Y, y)
}

func (v *Vector2Int) Serialize() string {
	//TODO
	return fmt.Sprintf("Vector2Int(%d, %d)", v.X, v.Y)
}

func (v *Vector2Int) ToString() string {
	return fmt.Sprintf("Vector2Int(%d, %d)", v.X, v.Y)
}

//---------------------------------------------------------

// Contains method will check and see if `rect` argument is in its
// region or not (even if it's only 0.1 point, it will return true).
func (r *RectangleInt) Contains(rect RectangleInt) bool {

	return false
}

// Encloses method will check and see if
func (r *RectangleInt) Encloses(rect RectangleInt) bool {

	return false
}

func (r *RectangleInt) IsLarger(rect RectangleInt) bool {
	return false
}

func (r *RectangleInt) IsLargerOrEqual(rect RectangleInt) bool {
	return false
}

func (r *RectangleInt) IsSmaller(rect RectangleInt) bool {
	return false
}

func (r *RectangleInt) IsSmallerOrEqual(rect RectangleInt) bool {
	return false
}

//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
