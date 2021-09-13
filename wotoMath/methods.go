package wotoMath

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
//---------------------------------------------------------
