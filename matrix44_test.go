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

func TestMakeIdentifyMatrix(t *testing.T) {
	expected := Matrix44{
		m: [4][4]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}

	var received = MakeIdentityMatrix()

	if received != expected {
		t.Errorf("MakeIdentifyMatrix failed - received: %.2f, expected: %.2f", received, expected)
	}
}
