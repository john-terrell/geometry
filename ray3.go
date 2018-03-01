package geometry

/*
	Copyright 2018 John W. Terrell.  All Rights Reserved.

	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

// Ray3 is a 3D ray with origin and vector components.
type Ray3 struct {
	Origin    Vector3
	Direction Vector3
}

// Parameterize returns a point along the ray at the location defined by the given parameter.
func (r Ray3) Parameterize(t float64) Vector3 {
	result := Vector3{
		X: r.Origin.X + (r.Direction.X * t),
		Y: r.Origin.Y + (r.Direction.Y * t),
		Z: r.Origin.Z + (r.Direction.Z * t),
	}

	return result
}
