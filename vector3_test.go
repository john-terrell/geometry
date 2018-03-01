package geometry

/*
	Copyright 2018 John W. Terrell.  All Rights Reserved.

	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

import (
	"testing"
)

func TestVector3_XAxis(t *testing.T) {
	expected := Vector3{1, 0, 0}
	var received = XAxis()

	if received != expected {
		t.Errorf("XAxis failed - received: %.2f, expected: %.2f", received, expected)
	}

}

func TestVector3_YAxis(t *testing.T) {
	expected := Vector3{0, 1, 0}
	var received = YAxis()

	if received != expected {
		t.Errorf("YAxis failed - received: %.2f, expected: %.2f", received, expected)
	}

}

func TestVector3_ZAxis(t *testing.T) {
	expected := Vector3{0, 0, 1}
	var received = ZAxis()

	if received != expected {
		t.Errorf("ZAxis failed - received: %.2f, expected: %.2f", received, expected)
	}

}

func TestVector3_Scaled(t *testing.T) {
	v := Vector3{1, 2, 3}
	expected := Vector3{3, 6, 9}
	var received = v.Scaled(3)

	if received != expected {
		t.Errorf("Scaled failed - received: %.2f, expected: %.2f", received, expected)
	}
}

func TestVector3_Inverse(t *testing.T) {
	v := Vector3{1, 2, 3}
	expected := Vector3{-1, -2, -3}
	var received = v.Inverse()

	if received != expected {
		t.Errorf("Inverse failed - received: %.2f, expected: %.2f", received, expected)
	}
}

func TestVector3_SquaredMagnitude(t *testing.T) {
	v := Vector3{0, 2, 0}
	const expected = 4.0
	var received = v.SquaredMagnitude()

	if received != expected {
		t.Errorf("SquaredMagnitude failed - received: %.2f, expected: %.2f", received, expected)
	}
}

func TestVector3_Magnitude(t *testing.T) {
	v := Vector3{0, 2, 0}
	const expected = 2.0
	var received = v.Magnitude()

	if received != expected {
		t.Errorf("Magnitude failed - received: %.2f, expected: %.2f", received, expected)
	}
}

func TestVector3_Invert(t *testing.T) {
	v1 := Vector3{1, 2, 3}
	expected := Vector3{-1, -2, -3}
	v1.Invert()

	if v1 != expected {
		t.Errorf("Inversion of vector failed - received: %.2f, expected: %.2f", v1, expected)
	}
}

func TestVector3_Add(t *testing.T) {
	v1 := Vector3{1, 2, 3}
	v2 := Vector3{-1, -2, -3}

	expected := Vector3{0, 0, 0}
	var received = AddVectors(v1, v2)

	if received != expected {
		t.Errorf("Addition of vectors failed - received: %.2f, expected: %.2f", received, expected)
	}
}

func TestVector3_DotProduct(t *testing.T) {
	v1 := Vector3{0, 1, 0}
	v2 := Vector3{0, 1, 0}
	var expected = 1.0
	var received = DotProduct(v1, v2)

	if received != expected {
		t.Errorf("DotProduct of coincident vectors failed - received: %.2f, expected: %.2f", received, expected)
	}

	v2 = Vector3{1, 0, 0}
	expected = 0.0
	received = DotProduct(v1, v2)

	if received != expected {
		t.Errorf("DotProduct of perpendicular vectors failed - received: %.2f, expected: %.2f", received, expected)
	}
}

func TestVector3_CrossProduct(t *testing.T) {
	v1 := Vector3{0, 0, 1}
	v2 := Vector3{0, 1, 0}
	var expected = Vector3{-1, 0, 0}
	var received = CrossProduct(v1, v2)

	if received != expected {
		t.Errorf("CrossProduct of coincident vectors failed - received: %+v, expected: %+v", received, expected)
	}
}

func TestVector3_Normalize(t *testing.T) {
	v1 := Vector3{0, 0, 2}
	v1.Normalize()
	var expected = 1.0
	var received = v1.Magnitude()
	if received != expected {
		t.Errorf("Normalization of vector failed - received: %+v, expected: %+v", received, expected)
	}
}
