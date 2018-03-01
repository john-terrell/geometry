package geometry

/*
	Copyright 2018 John W. Terrell.  All Rights Reserved.

	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

import "testing"

func TestRay3_Parameterize(t *testing.T) {
	r := Ray3{
		Origin:    Vector3{X: 0, Y: 0, Z: 0},
		Direction: Vector3{X: 0, Y: 0, Z: 1},
	}

	const parameter = 2.5
	received := r.Parameterize(parameter)
	expected := Vector3{X: 0, Y: 0, Z: 2.5}

	if received != expected {
		t.Errorf("Parameterization of ray failed - received: %+v, expected: %+v", received, expected)
	}
}
