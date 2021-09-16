// wotoMath Project
// Copyright (C) 2021 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMath

import (
	"fmt"
	"math"
	"strings"
)

//---------------------------------------------------------

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinUInt(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func MaxUInt(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func MinInt8(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

func MaxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func MinUInt8(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}

func MaxUInt8(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

func MinInt16(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

func MaxInt16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func MinUInt16(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

func MaxUInt16(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func MinUInt32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func MaxUInt32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MinUInt64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func MaxUInt64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func Sqrt(a int) float64 {
	return math.Sqrt(float64(a * a))
}

//---------------------------------------------------------

func EqualFloats64(a, b float64) bool {
	if !EqualSignFloats64(a, b) {
		return false
	}

	a_fraction := GetFloatingCount64(a)
	b_fraction := GetFloatingCount64(b)

	// there is a checker for 0 values of tolerance here so
	// it increases the speed of running in a case one of the
	// values doesn't have any number after floating point.
	// for example: `10.1` and `10.00`
	if a_fraction == 0 || b_fraction == 0 {
		return int(a) == int(b)
	}

	return EqualFloatsT64(a, b,
		MinInt(a_fraction, b_fraction))
}

func EqualFloats32(a, b float32) bool {
	if !EqualSignFloats32(a, b) {
		return false
	}

	a_fraction := GetFloatingCount32(b)
	b_fraction := GetFloatingCount32(b)

	// there is a checker for 0 values of tolerance here so
	// it increases the speed of running in a case one of the
	// values doesn't have any number after floating point.
	// for example: `10.1` and `10.00`
	if a_fraction == 0 || b_fraction == 0 {
		return int(a) == int(b)
	}

	return EqualFloatsT32(a, b, MinInt(a_fraction, b_fraction))
}

func GetFloatingCount64(a float64) int {
	str := fmt.Sprintf("%g", a)

	strs := strings.Split(str, ".")
	if len(strs) < 2 {
		return 0
	}

	return len(strs[1])
}

func GetFloatingCount32(a float32) int {
	var count int
	for a != float32(int(a)) {
		a *= 10
		count++
	}

	return count
}

func EqualFloatsT64(a, b float64, tolerance int) bool {
	if !EqualSignFloats64(a, b) {
		return false
	}

	ep := math.Abs(a - b)
	maxAllowDiff := 1 / math.Pow(10, float64(tolerance))

	return ep < maxAllowDiff
}

func EqualFloatsT32(a, b float32, tolerance int) bool {
	if !EqualSignFloats32(a, b) {
		return false
	}

	ep := math.Abs(float64(a - b))
	maxAllowDiff := 1 / math.Pow(10, float64(tolerance))

	return ep < maxAllowDiff
}

func EqualSignFloats64(a, b float64) bool {
	return (a > 0 && b > 0) || (a < 0 && b < 0)
}

func EqualSignFloats32(a, b float32) bool {
	return (a > 0 && b > 0) || (a < 0 && b < 0)
}

//---------------------------------------------------------

func NewVector2(x float64, y float64) *Vector2 {
	return &Vector2{X: x, Y: y}
}

func NewVector2Int(x int, y int) *Vector2Int {
	return &Vector2Int{X: x, Y: y}
}

func NewVector2UInt(x uint, y uint) *Vector2UInt {
	return &Vector2UInt{X: x, Y: y}
}

func NewVector2Int8(x int8, y int8) *Vector2Int8 {
	return &Vector2Int8{X: x, Y: y}
}

func NewVector2UInt8(x uint8, y uint8) *Vector2UInt8 {
	return &Vector2UInt8{X: x, Y: y}
}

func NewVector2Int16(x int16, y int16) *Vector2Int16 {
	return &Vector2Int16{X: x, Y: y}
}

func NewVector2UInt16(x uint16, y uint16) *Vector2UInt16 {
	return &Vector2UInt16{X: x, Y: y}
}

func NewVector2Int32(x int32, y int32) *Vector2Int32 {
	return &Vector2Int32{X: x, Y: y}
}

func NewVector2UInt32(x uint32, y uint32) *Vector2UInt32 {
	return &Vector2UInt32{X: x, Y: y}
}

func NewVector2Int64(x int64, y int64) *Vector2Int64 {
	return &Vector2Int64{X: x, Y: y}
}

func NewVector2UInt64(x uint64, y uint64) *Vector2UInt64 {
	return &Vector2UInt64{X: x, Y: y}
}

func NewRectangle(x, y, width, height float64) *Rectangle {
	return &Rectangle{
		Location: Vector2{X: x, Y: y},
		Size:     Vector2{X: width, Y: height},
	}
}

func NewRectangleVect(location, size Vector2) *Rectangle {
	return &Rectangle{
		Location: location,
		Size:     size,
	}
}

func NewRectangleInt(x, y, width, height int) *RectangleInt {
	return &RectangleInt{
		Location: Vector2Int{X: x, Y: y},
		Size:     Vector2Int{X: width, Y: height},
	}
}

func NewRectangleUInt(x, y, width, height uint) *RectangleUInt {
	return &RectangleUInt{
		Location: Vector2UInt{X: x, Y: y},
		Size:     Vector2UInt{X: width, Y: height},
	}
}

func NewRectangleIntVector(location, size Vector2Int) *RectangleInt {
	return &RectangleInt{
		Location: location,
		Size:     size,
	}
}

func NewRectangleUIntVector(location, size Vector2UInt) *RectangleUInt {
	return &RectangleUInt{
		Location: location,
		Size:     size,
	}
}

func NewRectangleInt8(x, y, width, height int8) *RectangleInt8 {
	return &RectangleInt8{
		Location: Vector2Int8{X: x, Y: y},
		Size:     Vector2Int8{X: width, Y: height},
	}
}

func NewRectangleUInt8(x, y, width, height uint8) *RectangleUInt8 {
	return &RectangleUInt8{
		Location: Vector2UInt8{X: x, Y: y},
		Size:     Vector2UInt8{X: width, Y: height},
	}
}

func NewRectangleInt8Vector(location, size Vector2Int8) *RectangleInt8 {
	return &RectangleInt8{
		Location: location,
		Size:     size,
	}
}

func NewRectangleUInt8Vector(location, size Vector2UInt8) *RectangleUInt8 {
	return &RectangleUInt8{
		Location: location,
		Size:     size,
	}
}

func NewRectangleInt16(x, y, width, height int16) *RectangleInt16 {
	return &RectangleInt16{
		Location: Vector2Int16{X: x, Y: y},
		Size:     Vector2Int16{X: width, Y: height},
	}
}

func NewRectangleUInt16(x, y, width, height uint16) *RectangleUInt16 {
	return &RectangleUInt16{
		Location: Vector2UInt16{X: x, Y: y},
		Size:     Vector2UInt16{X: width, Y: height},
	}
}

func NewRectangleInt16Vector(location, size Vector2Int16) *RectangleInt16 {
	return &RectangleInt16{
		Location: location,
		Size:     size,
	}
}

func NewRectangleUInt16Vector(location, size Vector2UInt16) *RectangleUInt16 {
	return &RectangleUInt16{
		Location: location,
		Size:     size,
	}
}

func NewRectangleInt32(x, y, width, height int32) *RectangleInt32 {
	return &RectangleInt32{
		Location: Vector2Int32{X: x, Y: y},
		Size:     Vector2Int32{X: width, Y: height},
	}
}

func NewRectangleUInt32(x, y, width, height uint32) *RectangleUInt32 {
	return &RectangleUInt32{
		Location: Vector2UInt32{X: x, Y: y},
		Size:     Vector2UInt32{X: width, Y: height},
	}
}

func NewRectangleInt32Vector(location, size Vector2Int32) *RectangleInt32 {
	return &RectangleInt32{
		Location: location,
		Size:     size,
	}
}

func NewRectangleUInt32Vector(location, size Vector2UInt32) *RectangleUInt32 {
	return &RectangleUInt32{
		Location: location,
		Size:     size,
	}
}

func NewRectangleInt64(x, y, width, height int64) *RectangleInt64 {
	return &RectangleInt64{
		Location: Vector2Int64{X: x, Y: y},
		Size:     Vector2Int64{X: width, Y: height},
	}
}

func NewRectangleUInt64(x, y, width, height uint64) *RectangleUInt64 {
	return &RectangleUInt64{
		Location: Vector2UInt64{X: x, Y: y},
		Size:     Vector2UInt64{X: width, Y: height},
	}
}

func NewRectangleInt64Vector(location, size Vector2Int64) *RectangleInt64 {
	return &RectangleInt64{
		Location: location,
		Size:     size,
	}
}

func NewRectangleUInt64Vector(location, size Vector2UInt64) *RectangleUInt64 {
	return &RectangleUInt64{
		Location: location,
		Size:     size,
	}
}

//---------------------------------------------------------

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

//---------------------------------------------------------

func DotInt(ihs *Vector2Int, rhs *Vector2Int) int {
	return ihs.X*rhs.X + ihs.Y*rhs.Y
}

func LerpInt(a *Vector2Int, b *Vector2Int, t int) *Vector2Int {
	return NewVector2Int(
		a.X+(b.X-a.X)*t,
		a.Y+(b.Y-a.Y)*t,
	)
}

func DistanceInt(a *Vector2Int, b *Vector2Int) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return Sqrt(dx*dx + dy*dy)
}

func ReflectInt(ihs *Vector2Int, rhs *Vector2Int) *Vector2Int {
	factor := -2.0 * DotInt(ihs, rhs)
	return NewVector2Int(
		factor*ihs.X+rhs.X,
		factor*ihs.Y+rhs.Y,
	)
}

//---------------------------------------------------------
