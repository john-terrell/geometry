package geometry

/*
	Copyright 2018 John W. Terrell.  All Rights Reserved.

	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

import (
	"math"
)

// Vector3 is a 3D vector type with X, Y, and Z components.
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// ** Properties **

// SquaredMagnitude returns the square of the vector's magnitude.
func (v Vector3) SquaredMagnitude() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

// Magnitude returns the vector's magnitude.
func (v Vector3) Magnitude() float64 {
	return math.Sqrt(v.SquaredMagnitude())
}

// DotProduct returns the dot product of two vectors.
func DotProduct(a Vector3, b Vector3) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z)
}

// CrossProduct returns the cross product of two vectors. (Right handed)
func CrossProduct(left Vector3, right Vector3) Vector3 {
	return Vector3{
		X: left.Y*right.Z - left.Z*right.Y,
		Y: left.Z*right.X - left.X*right.Z,
		Z: left.X*right.Y - left.Y*right.X,
	}
}

// ** Modifiers **

// Scale modifes the magnitude of the vector by the given scaling multiplier.
func (v *Vector3) Scale(s float64) {
	v.X *= s
	v.Y *= s
	v.Z *= s
}

// Normalize sets the magnitude of the vector to unity without affecting its direction.
func (v *Vector3) Normalize() {
	var m = v.Magnitude()
	v.Scale(1 / m)
}
