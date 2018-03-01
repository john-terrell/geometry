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

// Scaled returns the inverse of the vector
func (v Vector3) Scaled(scaler float64) Vector3 {
	return Vector3{
		X: v.X * scaler,
		Y: v.Y * scaler,
		Z: v.Z * scaler,
	}
}

// Inverse returns the inverse of the vector
func (v Vector3) Inverse() Vector3 {
	return Vector3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

// SquaredMagnitude returns the square of the vector's magnitude.
func (v Vector3) SquaredMagnitude() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

// Magnitude returns the vector's magnitude.
func (v Vector3) Magnitude() float64 {
	return math.Sqrt(v.SquaredMagnitude())
}

// ** Modifiers **

// Invert reverses the direction of the vector by inverting each component.
func (v *Vector3) Invert() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

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

// ** Operations **

// XAxis returns a vector that points down the positive X axis
func XAxis() Vector3 {
	return Vector3{1, 0, 0}
}

// YAxis returns a vector that points down the positive Y axis
func YAxis() Vector3 {
	return Vector3{0, 1, 0}
}

// ZAxis returns a vector that points down the positive Z axis
func ZAxis() Vector3 {
	return Vector3{0, 0, 1}
}

// AddVectors returns a vector that is the sum of the two argument vectors
func AddVectors(left Vector3, right Vector3) Vector3 {
	return Vector3{
		X: left.X + right.X,
		Y: left.Y + right.Y,
		Z: left.Z + right.Z,
	}
}

// DotProduct returns the dot product of two vectors.
func DotProduct(left Vector3, right Vector3) float64 {
	return (left.X * right.X) + (left.Y * right.Y) + (left.Z * right.Z)
}

// CrossProduct returns the cross product of two vectors. (Right handed)
func CrossProduct(left Vector3, right Vector3) Vector3 {
	return Vector3{
		X: left.Y*right.Z - left.Z*right.Y,
		Y: left.Z*right.X - left.X*right.Z,
		Z: left.X*right.Y - left.Y*right.X,
	}
}
