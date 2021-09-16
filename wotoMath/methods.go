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

func (v *Vector2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v *Vector2) IsZeroX() bool {
	return v.X == 0
}

func (v *Vector2) IsZeroY() bool {
	return v.Y == 0
}

func (v *Vector2) IsZeroXOrY() bool {
	return v.Y == 0 || v.X == 0
}

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

func (r *Rectangle) IsEmpty() bool {
	return r.Size.IsZero()
}

func (r *Rectangle) Left() float64 {
	return r.Location.X
}

func (r *Rectangle) Top() float64 {
	return r.Location.X
}

func (r *Rectangle) Right() float64 {
	return r.Location.X + r.Size.X
}

func (r *Rectangle) Bottom() float64 {
	return r.Location.Y + r.Size.Y
}

func (r *Rectangle) BottomLeft() Vector2 {
	return Vector2{X: r.Left(), Y: r.Bottom()}
}

func (r *Rectangle) BottomRight() Vector2 {
	return Vector2{X: r.Right(), Y: r.Bottom()}
}

func (r *Rectangle) TopRight() Vector2 {
	return Vector2{X: r.Right(), Y: r.Top()}
}

func (r *Rectangle) TopLeft() Vector2 {
	return Vector2{X: r.Left(), Y: r.Top()}
}

func (r *Rectangle) Width() float64 {
	return r.Size.X
}

func (r *Rectangle) Height() float64 {
	return r.Size.Y
}

// Contains method will check and see if `rect` argument is in its
// region or not (even if it's only 0.1 point, it will return true).
func (r *Rectangle) Contains(rect Rectangle) bool {
	return r.ContainsPoint(rect.TopLeft()) ||
		r.ContainsPoint(rect.TopRight()) ||
		r.ContainsPoint(rect.BottomLeft()) ||
		r.ContainsPoint(rect.BottomRight())
}

func (r *Rectangle) ContainsPoint(point Vector2) bool {
	xb := point.X > r.Left() && point.X < r.Right()
	yb := point.Y > r.Top() && point.Y < r.Bottom()

	return xb && yb
}

// Encloses method will check and see if the current rectangle
// completely closes the `rect` argument or not.
// it should contains all four major points of the argument,
// which are `top-left`, `top-right`, `bottom-left` and `bottom-right`.
// if this rectangle doesn't contain even "one of the them", this
// method will return false.
func (r *Rectangle) Encloses(rect Rectangle) bool {
	return r.ContainsPoint(rect.TopLeft()) &&
		r.ContainsPoint(rect.TopRight()) &&
		r.ContainsPoint(rect.BottomLeft()) &&
		r.ContainsPoint(rect.BottomRight())
}

func (r *Rectangle) IsLarger(rect Rectangle) bool {
	return r.Size.IsLargerXY(rect.Size)
}

func (r *Rectangle) IsLargerOrEqual(rect Rectangle) bool {
	return r.Size.IsLargerOrEqualXY(rect.Size)
}

func (r *Rectangle) IsSmaller(rect Rectangle) bool {
	return r.Size.IsSmallerXY(rect.Size)
}

func (r *Rectangle) IsSmallerOrEqual(rect Rectangle) bool {
	return r.Size.IsSmallerOrEqualXY(rect.Size)
}

//---------------------------------------------------------

func (v *Vector2Int) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v *Vector2Int) IsZeroX() bool {
	return v.X == 0
}

func (v *Vector2Int) IsZeroY() bool {
	return v.Y == 0
}

func (v *Vector2Int) IsZeroXOrY() bool {
	return v.Y == 0 || v.X == 0
}

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
}

func (v *Vector2Int) Serialize() string {
	//TODO
	return fmt.Sprintf("Vector2Int(%d, %d)", v.X, v.Y)
}

func (v *Vector2Int) ToString() string {
	return fmt.Sprintf("Vector2Int(%d, %d)", v.X, v.Y)
}

//---------------------------------------------------------

func (r *RectangleInt) IsEmpty() bool {
	return r.Size.IsZero()
}

func (r *RectangleInt) Left() int {
	return r.Location.X
}

func (r *RectangleInt) Top() int {
	return r.Location.X
}

func (r *RectangleInt) Right() int {
	return r.Location.X + r.Size.X
}

func (r *RectangleInt) Bottom() int {
	return r.Location.Y + r.Size.Y
}

func (r *RectangleInt) BottomLeft() Vector2Int {
	return Vector2Int{X: r.Left(), Y: r.Bottom()}
}

func (r *RectangleInt) BottomRight() Vector2Int {
	return Vector2Int{X: r.Right(), Y: r.Bottom()}
}

func (r *RectangleInt) TopRight() Vector2Int {
	return Vector2Int{X: r.Right(), Y: r.Top()}
}

func (r *RectangleInt) TopLeft() Vector2Int {
	return Vector2Int{X: r.Left(), Y: r.Top()}
}

func (r *RectangleInt) Width() int {
	return r.Size.X
}

func (r *RectangleInt) Height() int {
	return r.Size.Y
}

// Contains method will check and see if `rect` argument is in its
// region or not (even if it's only 0.1 point, it will return true).
func (r *RectangleInt) Contains(rect RectangleInt) bool {
	return r.ContainsPoint(rect.TopLeft()) ||
		r.ContainsPoint(rect.TopRight()) ||
		r.ContainsPoint(rect.BottomLeft()) ||
		r.ContainsPoint(rect.BottomRight())
}

func (r *RectangleInt) ContainsPoint(point Vector2Int) bool {
	xb := point.X > r.Left() && point.X < r.Right()
	yb := point.Y > r.Top() && point.Y < r.Bottom()

	return xb && yb
}

// Encloses method will check and see if the current rectangle
// completely closes the `rect` argument or not.
// it should contains all four major points of the argument,
// which are `top-left`, `top-right`, `bottom-left` and `bottom-right`.
// if this rectangle doesn't contain even "one of the them", this
// method will return false.
func (r *RectangleInt) Encloses(rect RectangleInt) bool {
	return r.ContainsPoint(rect.TopLeft()) &&
		r.ContainsPoint(rect.TopRight()) &&
		r.ContainsPoint(rect.BottomLeft()) &&
		r.ContainsPoint(rect.BottomRight())
}

func (r *RectangleInt) IsLarger(rect RectangleInt) bool {
	return r.Size.IsLargerXY(rect.Size)
}

func (r *RectangleInt) IsLargerOrEqual(rect RectangleInt) bool {
	return r.Size.IsLargerOrEqualXY(rect.Size)
}

func (r *RectangleInt) IsSmaller(rect RectangleInt) bool {
	return r.Size.IsSmallerXY(rect.Size)
}

func (r *RectangleInt) IsSmallerOrEqual(rect RectangleInt) bool {
	return r.Size.IsSmallerOrEqualXY(rect.Size)
}

//---------------------------------------------------------

func (v *Vector2UInt) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v *Vector2UInt) IsZeroX() bool {
	return v.X == 0
}

func (v *Vector2UInt) IsZeroY() bool {
	return v.Y == 0
}

func (v *Vector2UInt) IsZeroXOrY() bool {
	return v.Y == 0 || v.X == 0
}

func (v *Vector2UInt) IsLarger(vect Vector2UInt) bool {
	return v.X > vect.X || v.Y > vect.Y
}

func (v *Vector2UInt) IsLargerXY(vect Vector2UInt) bool {
	return v.X > vect.X && v.Y > vect.Y
}

func (v *Vector2UInt) IsBiggerX(vect Vector2UInt) bool {
	return v.X > vect.X
}

func (v *Vector2UInt) IsBiggerY(vect Vector2UInt) bool {
	return v.Y > vect.Y
}

func (v *Vector2UInt) IsLargerOrEqual(vect Vector2UInt) bool {
	return v.X >= vect.X || v.Y >= vect.Y
}

func (v *Vector2UInt) IsLargerOrEqualXY(vect Vector2UInt) bool {
	return v.X >= vect.X && v.Y >= vect.Y
}

func (v *Vector2UInt) IsBiggerOrEqualX(vect Vector2UInt) bool {
	return v.X >= vect.X
}

func (v *Vector2UInt) IsBiggerOrEqualY(vect Vector2UInt) bool {
	return v.Y >= vect.Y
}

func (v *Vector2UInt) IsSmaller(vect Vector2UInt) bool {
	return v.X < vect.X || v.Y < vect.Y
}

func (v *Vector2UInt) IsSmallerXY(vect Vector2UInt) bool {
	return v.X < vect.X && v.Y < vect.Y
}

func (v *Vector2UInt) IsSmallerX(vect Vector2UInt) bool {
	return v.X < vect.X
}

func (v *Vector2UInt) IsSmallerY(vect Vector2UInt) bool {
	return v.Y < vect.Y
}

func (v *Vector2UInt) IsSmallerOrEqual(vect Vector2UInt) bool {
	return v.X <= vect.X || v.Y <= vect.Y
}

func (v *Vector2UInt) IsSmallerOrEqualXY(vect Vector2UInt) bool {
	return v.X <= vect.X && v.Y <= vect.Y
}

func (v *Vector2UInt) IsSmallerOrEqualX(vect Vector2UInt) bool {
	return v.X <= vect.X
}

func (v *Vector2UInt) IsSmallerOrEqualY(vect Vector2UInt) bool {
	return v.Y <= vect.Y
}

func (v *Vector2UInt) Add(other *Vector2UInt) *Vector2UInt {
	if other == nil {
		return v
	}

	return NewVector2UInt(v.X+other.X, v.Y+other.Y)
}

func (v *Vector2UInt) AddThis(other *Vector2UInt) *Vector2UInt {
	if other == nil {
		return v
	}

	v.X += other.X
	v.Y += other.Y
	return v
}

func (v *Vector2UInt) Sub(other *Vector2UInt) *Vector2UInt {
	if other == nil {
		return v
	}

	return NewVector2UInt(v.X-other.X, v.Y-other.Y)
}

func (v *Vector2UInt) SubThis(other *Vector2UInt) *Vector2UInt {
	if other == nil {
		return v
	}

	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v *Vector2UInt) Mul(other *Vector2UInt) *Vector2UInt {
	if other == nil {
		return v
	}

	return NewVector2UInt(v.X*other.X, v.Y*other.Y)
}

func (v *Vector2UInt) MulThis(other *Vector2UInt) *Vector2UInt {
	if other == nil {
		return v
	}

	v.X *= other.X
	v.Y *= other.Y
	return v
}

func (v *Vector2UInt) Div(other *Vector2UInt) *Vector2UInt {
	if other == nil {
		return v
	}

	if other.X == 0 && other.Y == 0 {
		return NewVector2UInt(v.X, v.Y)
	} else if other.X == 0 {
		return NewVector2UInt(v.X, v.Y/other.Y)
	} else if other.Y == 0 {
		return NewVector2UInt(v.X/other.X, v.Y)
	}

	return NewVector2UInt(v.X/other.X, v.Y/other.Y)
}

func (v *Vector2UInt) DivThis(other *Vector2UInt) *Vector2UInt {
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

func (v *Vector2UInt) Copy() *Vector2UInt {
	return NewVector2UInt(v.X, v.Y)
}

func (v *Vector2UInt) Set(x uint, y uint) *Vector2UInt {
	v.X = x
	v.Y = y
	return v
}

func (v *Vector2UInt) SetX(x uint) *Vector2UInt {
	v.X = x
	return v
}

func (v *Vector2UInt) SetY(y uint) *Vector2UInt {
	v.Y = y
	return v
}

func (v *Vector2UInt) AddScalar(scalar uint) *Vector2UInt {
	return NewVector2UInt(v.X+scalar, v.Y+scalar)
}

func (v *Vector2UInt) AddScalarThis(scalar uint) *Vector2UInt {
	v.X += scalar
	v.Y += scalar
	return v
}

func (v *Vector2UInt) AddScalars(x, y uint) *Vector2UInt {
	return NewVector2UInt(v.X+x, v.Y+y)
}

func (v *Vector2UInt) AddScalarsThis(x, y uint) *Vector2UInt {
	v.X += x
	v.Y += y
	return v
}

func (v *Vector2UInt) SubScalar(scalar uint) *Vector2UInt {
	return NewVector2UInt(v.X-scalar, v.Y-scalar)
}

func (v *Vector2UInt) SubScalarThis(scalar uint) *Vector2UInt {
	v.X -= scalar
	v.Y -= scalar
	return v
}

func (v *Vector2UInt) SubScalars(x, y uint) *Vector2UInt {
	return NewVector2UInt(v.X-x, v.Y-y)
}

func (v *Vector2UInt) SubScalarsThis(x, y uint) *Vector2UInt {
	v.X -= x
	v.Y -= y
	return v
}

func (v *Vector2UInt) MulScalar(scalar uint) *Vector2UInt {
	return NewVector2UInt(v.X*scalar, v.Y*scalar)
}

func (v *Vector2UInt) MulScalarThis(scalar uint) *Vector2UInt {
	v.X *= scalar
	v.Y *= scalar
	return v
}

func (v *Vector2UInt) MulScalars(x, y uint) *Vector2UInt {
	return NewVector2UInt(v.X-x, v.Y-y)
}

func (v *Vector2UInt) MulScalarsThis(x, y uint) *Vector2UInt {
	v.X -= x
	v.Y -= y
	return v
}

func (v *Vector2UInt) DivScalar(scalar uint) *Vector2UInt {
	if scalar == 0 {
		return v
	}

	return NewVector2UInt(v.X*scalar, v.Y*scalar)
}

func (v *Vector2UInt) DivScalarThis(scalar uint) *Vector2UInt {
	if scalar == 0 {
		return v
	}

	v.X /= scalar
	v.Y /= scalar
	return v
}

func (v *Vector2UInt) DivScalars(x, y uint) *Vector2UInt {
	if x == 0 && y == 0 {
		return v
	} else if x == 0 {
		return NewVector2UInt(v.X, v.Y/y)
	} else if y == 0 {
		return NewVector2UInt(v.X/x, v.Y)
	}

	return NewVector2UInt(v.X/x, v.Y/y)
}

func (v *Vector2UInt) DivScalarsThis(x, y uint) *Vector2UInt {
	if x != 0 {
		v.X /= x
	}
	if y != 0 {
		v.Y /= y
	}

	return v
}

func (v *Vector2UInt) Distance(other *Vector2UInt) uint {
	if other == nil {
		return 0
	}

	dx := v.X - other.X
	dy := v.Y - other.Y
	return uint(math.Sqrt(float64(dx*dx + dy*dy)))
}

func (v *Vector2UInt) Dot(other *Vector2UInt) uint {
	if other == nil {
		return 0
	}

	return v.X*other.X + v.Y*other.Y
}

func (v *Vector2UInt) Lerp(other *Vector2UInt, t uint) *Vector2UInt {
	if other == nil {
		return v
	}

	return NewVector2UInt(
		v.X+(other.X-v.X)*t,
		v.Y+(other.Y-v.Y)*t,
	)
}

func (v *Vector2UInt) Magnitude() uint {
	return uint(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v *Vector2UInt) Normalize() *Vector2UInt {
	m := v.Magnitude()

	if float64(m) > Epsilon {
		return v.DivScalar(m)
	} else {
		return v.Copy()
	}
}

func (v *Vector2UInt) Equals(other *Vector2UInt) bool {
	if other == nil {
		return false
	}

	return v.X == other.X && v.Y == other.Y
}

func (v *Vector2UInt) EqualsX(other *Vector2UInt) bool {
	if other == nil {
		return false
	}

	return v.X == other.X
}

func (v *Vector2UInt) EqualsScalarX(x uint) bool {
	return v.X == x
}

func (v *Vector2UInt) EqualsY(other *Vector2UInt) bool {
	if other == nil {
		return false
	}

	return v.Y == other.Y
}

func (v *Vector2UInt) EqualsScalarY(y uint) bool {
	return v.Y == y
}

func (v *Vector2UInt) Serialize() string {
	//TODO
	return fmt.Sprintf("Vector2UInt(%d, %d)", v.X, v.Y)
}

func (v *Vector2UInt) ToString() string {
	return fmt.Sprintf("Vector2UInt(%d, %d)", v.X, v.Y)
}

//---------------------------------------------------------

func (r *RectangleUInt) IsEmpty() bool {
	return r.Size.IsZero()
}

func (r *RectangleUInt) Left() uint {
	return r.Location.X
}

func (r *RectangleUInt) Top() uint {
	return r.Location.X
}

func (r *RectangleUInt) Right() uint {
	return r.Location.X + r.Size.X
}

func (r *RectangleUInt) Bottom() uint {
	return r.Location.Y + r.Size.Y
}

func (r *RectangleUInt) BottomLeft() Vector2UInt {
	return Vector2UInt{X: r.Left(), Y: r.Bottom()}
}

func (r *RectangleUInt) BottomRight() Vector2UInt {
	return Vector2UInt{X: r.Right(), Y: r.Bottom()}
}

func (r *RectangleUInt) TopRight() Vector2UInt {
	return Vector2UInt{X: r.Right(), Y: r.Top()}
}

func (r *RectangleUInt) TopLeft() Vector2UInt {
	return Vector2UInt{X: r.Left(), Y: r.Top()}
}

func (r *RectangleUInt) Width() uint {
	return r.Size.X
}

func (r *RectangleUInt) Height() uint {
	return r.Size.Y
}

// Contains method will check and see if `rect` argument is in its
// region or not (even if it's only 0.1 point, it will return true).
func (r *RectangleUInt) Contains(rect RectangleUInt) bool {
	return r.ContainsPouint(rect.TopLeft()) ||
		r.ContainsPouint(rect.TopRight()) ||
		r.ContainsPouint(rect.BottomLeft()) ||
		r.ContainsPouint(rect.BottomRight())
}

func (r *RectangleUInt) ContainsPouint(point Vector2UInt) bool {
	xb := point.X > r.Left() && point.X < r.Right()
	yb := point.Y > r.Top() && point.Y < r.Bottom()

	return xb && yb
}

// Encloses method will check and see if the current rectangle
// completely closes the `rect` argument or not.
// it should contains all four major points of the argument,
// which are `top-left`, `top-right`, `bottom-left` and `bottom-right`.
// if this rectangle doesn't contain even "one of the them", this
// method will return false.
func (r *RectangleUInt) Encloses(rect RectangleUInt) bool {
	return r.ContainsPouint(rect.TopLeft()) &&
		r.ContainsPouint(rect.TopRight()) &&
		r.ContainsPouint(rect.BottomLeft()) &&
		r.ContainsPouint(rect.BottomRight())
}

func (r *RectangleUInt) IsLarger(rect RectangleUInt) bool {
	return r.Size.IsLargerXY(rect.Size)
}

func (r *RectangleUInt) IsLargerOrEqual(rect RectangleUInt) bool {
	return r.Size.IsLargerOrEqualXY(rect.Size)
}

func (r *RectangleUInt) IsSmaller(rect RectangleUInt) bool {
	return r.Size.IsSmallerXY(rect.Size)
}

func (r *RectangleUInt) IsSmallerOrEqual(rect RectangleUInt) bool {
	return r.Size.IsSmallerOrEqualXY(rect.Size)
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
