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
