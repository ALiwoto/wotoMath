// wotoMath Project
// Copyright (C) 2021 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package tests

import (
	"log"
	"testing"

	"github.com/ALiwoto/wotoMath/wotoMath"
)

//---------------------------------------------------------

func TestVector01(t *testing.T) {
	v1 := wotoMath.NewVector2(5, 5)
	v2 := wotoMath.NewVector2(5, 20)
	distance := v1.Distance(v2)
	if distance != 15 {
		t.Errorf("expected distance of 15, got %f", distance)
		return
	}

	v1 = wotoMath.NewVector2(19, 78)
	v2 = wotoMath.NewVector2(59, 68)
	distance = v1.Distance(v2)
	if !wotoMath.EqualFloatsT64(distance, 41.23, 2) {
		t.Errorf("expected distance of 41.23, got %f", distance)
		return
	}

	v1 = wotoMath.NewVector2(10, 0)
	v2 = wotoMath.NewVector2(103.6, -19)
	distance = v1.Distance(v2)
	if !wotoMath.EqualFloats64(distance, 95.5089) {
		t.Errorf("expected distance of 95.5089, got %f", distance)
		return
	}

	v1 = wotoMath.NewVector2(-5, -500)
	v2 = wotoMath.NewVector2(95, 290)
	distance = v1.Distance(v2)
	if !wotoMath.EqualFloats64(distance, 796.30396206473) {
		t.Errorf("expected distance of 796.30396206473, got %f", distance)
		return
	}

	v1 = wotoMath.NewVector2(-5.934847534, -500)
	v2 = wotoMath.NewVector2(95.54587, 290.78456)
	distance = v1.Distance(v2)
	if !wotoMath.EqualFloats64(distance, 797.2) {
		t.Errorf("expected distance of 797.2, got %f", distance)
		return
	}

	if !v1.EqualsScalarX(-5.934) {
		t.Errorf("expected v1.x of -5.934, got %f", v1.X)
		return
	}
}

func TestVector02(t *testing.T) {
	v1 := wotoMath.NewVector2(-5.934847534, -500)
	//v2 := wotoMath.NewVector2(95.54587, 290.78456)

	if !v1.EqualsScalarX(-5.934) {
		t.Errorf("expected v1.X of -5.934, got %f", v1.X)
		return
	}

	if !v1.EqualsScalarY(-500.000000000000000000000498) {
		t.Errorf("expected v1.Y of -500, got %f", v1.X)
		return
	}
}

//---------------------------------------------------------

func TestVectorInt01(t *testing.T) {
	v1 := wotoMath.NewVector2Int(5, 5)
	v2 := wotoMath.NewVector2Int(5, 20)
	distance := v1.Distance(v2)
	if distance != 15 {
		t.Errorf("expected distance of 15, got %d", distance)
		return
	}

	v1 = wotoMath.NewVector2Int(19, 78)
	v2 = wotoMath.NewVector2Int(59, 68)
	distance = v1.Distance(v2)
	if distance != 41 {
		t.Errorf("expected distance of 41, got %d", distance)
		return
	}

	v1 = wotoMath.NewVector2Int(10, 0)
	v2 = wotoMath.NewVector2Int(103, -19)
	distance = v1.Distance(v2)
	if distance != 94 {
		t.Errorf("expected distance of 95, got %d", distance)
		return
	}

	v1 = wotoMath.NewVector2Int(-5, -500)
	v2 = wotoMath.NewVector2Int(95, 290)
	distance = v1.Distance(v2)
	if distance != 796 {
		t.Errorf("expected distance of 796, got %d", distance)
		return
	}

	v1 = wotoMath.NewVector2Int(-5, -500)
	v2 = wotoMath.NewVector2Int(95, 290)
	distance = v1.Distance(v2)
	if distance != 796 {
		t.Errorf("expected distance of 796, got %d", distance)
		return
	}

	if !v1.EqualsScalarX(-5) {
		t.Errorf("expected v1.X of -5, got %d", v1.X)
		return
	}

	if !v1.EqualsScalarY(-500) {
		t.Errorf("expected v1.Y of -500, got %d", v1.Y)
		return
	}

	if !v2.EqualsScalarX(95) {
		t.Errorf("expected v2.X of 95, got %d", v2.X)
		return
	}

	if !v2.EqualsScalarY(290) {
		t.Errorf("expected v2.Y of 290, got %d", v2.Y)
		return
	}

	log.Println("ended")
}

//---------------------------------------------------------

func TestRectangle01(t *testing.T) {
	rect1 := wotoMath.NewRectangle(0, 0, 10, 10)
	rect2 := wotoMath.NewRectangle(5, 5, 10, 10)

	if !rect1.Contains(*rect2) {
		t.Error("rect1 should contain rect2; but Contains methor returned false")
		return
	}
}

//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
