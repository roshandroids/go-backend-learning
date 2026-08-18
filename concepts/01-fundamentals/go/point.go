// Package fundamentals holds Stage 1 exercises: package structure, structs,
// methods, value vs pointer receivers, and export rules.
package fundamentals

import "math"

// Point represents a 2D coordinate. Exported because it's capitalized —
// there is no "public" keyword in Go.
type Point struct {
	X int
	Y int
}

// DistanceTo returns the Euclidean distance between p and other.
// Value receiver: p is a COPY inside this method, not a reference to
// the caller's original Point.
func (p Point) DistanceTo(other Point) float64 {
	dx := float64(p.X - other.X)
	dy := float64(p.Y - other.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// Translate shifts p by (dx, dy) in place.
// Pointer receiver: this mutates the caller's original Point.
func (p *Point) Translate(dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Scale is intentionally the WRONG choice of receiver — a value receiver
// means any mutation here happens to a copy and is thrown away when the
// method returns. See ScaleInPlace for the fix, and point_test.go for the
// test that proves the difference.
//
// TODO(exercise, Level 2): multiply p.X and p.Y by factor.
func (p Point) Scale(factor float64) {
	// TODO: p.X = int(float64(p.X) * factor); p.Y = int(float64(p.Y) * factor)
}

// ScaleInPlace is the corrected version: a pointer receiver mutates the
// caller's original Point, same as Translate above.
//
// TODO(exercise, Level 2): multiply p.X and p.Y by factor, in place.
func (p *Point) ScaleInPlace(factor float64) {
	// TODO: p.X = int(float64(p.X) * factor); p.Y = int(float64(p.Y) * factor)
}
