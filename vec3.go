package vec3

import "math"

// Vector3 struct
type Vector3 struct {
	X, Y, Z float32
}

// Add two Vector3's together
func Add(a, b Vector3) Vector3 {
	return Vector3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Mult -iply a Vector
func Mult(a Vector3, b float32) Vector3 {
	return Vector3{a.X * b, a.Y * b, a.Z * b}
}

// Length of a Vector3
func (a Vector3) Length() float32 {
	return float32(math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z)))
}

// Distance of a Vector3
func Distance(a, b Vector3) float32 {
	xDiff := a.X - b.X
	yDiff := b.Y - b.Y
	zDiff := a.Z - b.Z
	return float32(math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff + zDiff*zDiff)))
}

// DistanceSquared of a Vector3
func DistanceSquared(a, b Vector3) float32 {
	xDiff := a.X - b.X
	yDiff := b.Y - b.Y
	zDiff := a.Z - b.Z
	return xDiff*xDiff + yDiff*yDiff + zDiff*zDiff
}

// Normalize a Vector
func Normalize(a Vector3) Vector3 {
	len := a.Length()
	return Vector3{a.X / len, a.Y / len, a.Z / len}
}
