package core

import (
	"fmt"
	"math"
)

type Vec struct {
    X, Y, Z float64
}

func (v *Vec) plus(other *Vec) *Vec {
    return &Vec{X: (v.X + other.X), Y: (v.Y + other.Y), Z: (v.Z + other.Z)}
}

func (v *Vec) plusM(other *Vec) {
    v.X += other.X;
    v.Y += other.Y;
    v.Z += other.Z;
}

func (v *Vec) plusAll(d float64) {
    v.X += d;
    v.Y += d;
    v.Z += d;
}

func (v *Vec) Minus(other *Vec) *Vec {
    return &Vec{X: (v.X - other.X), Y: (v.Y - other.Y), Z: (v.Z - other.Z)}
}

func (v *Vec) minusM(other *Vec) {
    v.X -= other.X;
    v.Y -= other.Y;
    v.Z -= other.Z;
}

func (v *Vec) vecTimesM(other *Vec) {
    v.X *= other.X;
    v.Y *= other.Y;
    v.Z *= other.Z;
}

func (v *Vec) times(d float64) *Vec {
    return &Vec{X: v.X*d, Y: v.Y*d, Z: v.Z*d}
}

func (v *Vec) timesM(d float64) {
    v.X *= d;
    v.Y *= d;
    v.Z *= d;
}

func (v *Vec) dot(other *Vec) float64 {
    return v.X*other.X + v.Y*other.Y + v.Z*other.Z;
}

func (v *Vec) length() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec) Normalize() *Vec {
    len := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z);
    if len == 0.0 { return v }
    return &Vec{v.X/len, v.Y/len, v.Z/len}
}

func (v *Vec) normalizeM() {
    len := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z);
    if len == 0.0 { return }
    v.X /= len;
    v.Y /= len;
    v.Z /= len;
}

func (v *Vec) String() string {
    return fmt.Sprintf("%f %f %f", v.X, v.Y, v.Z)
}