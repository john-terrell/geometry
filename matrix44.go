package geometry

/*
	Copyright 2018 John W. Terrell.  All Rights Reserved.

	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

// Matrix44 is a 4x4 transformation matrix.
type Matrix44 struct {
	m [4][4]float64
}

// MakeIdentityMatrix returns a new identity matrix.
func MakeIdentityMatrix() Matrix44 {
	return Matrix44{
		m: [4][4]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}
