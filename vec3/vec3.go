package vec3

import (
	"fmt"
	"math"

	"github.com/ichibankunio/fvec/vec2"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func New(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

// func (v Vec3) Add(other Vec3) Vec3 {
// 	return New(v.X+other.X, v.Y+other.Y, v.Z+other.Z)
// }

func NewFromVec2(v vec2.Vec2) Vec3 {
	return Vec3{v.X, v.Y, 0}
}

func (v Vec3) Add(other Vec3) Vec3 {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z

	return v
}

// func (v Vec3) Sub(other Vec3) Vec3 {
// 	return New(v.X-other.X, v.Y-other.Y, v.Z-other.Z)
// }

func (v Vec3) Sub(other Vec3) Vec3 {
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z

	return v
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// func VecFromAngle(angle, magnitude float64) Vec3 {
// 	return Vec3{
// 		math.Cos(angle) * magnitude,
// 		math.Sin(angle) * magnitude,
// 	}
// }

func (v Vec3) Clone() Vec3 {
	return Vec3{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
}

func NewArray(arr [][3]float64) []Vec3 {
	out := make([]Vec3, len(arr))
	for i, _ := range arr {
		out[i] = New(arr[i][0], arr[i][1], arr[i][2])
	}

	return out
}

// func (v Vec3) Heading() float64 {
// 	return math.Atan2(v.X, v.Y)
// }

// func (v Vec3) Scale(s float64) Vec3 {
// 	return New(v.X*s, v.Y*s, v.Z*s)
// }

func (v Vec3) Scale(s float64) Vec3 {
	v.X *= s
	v.Y *= s
	v.Z *= s

	return v
}

// func (v Vec3) Mul(other Vec3) Vec3 {
// 	return New(v.X*other.X, v.Y*other.Y, v.Z*other.Z)
// }

func (v Vec3) Mul(other Vec3) Vec3 {
	v.X *= other.X
	v.Y *= other.Y
	v.Z *= other.Z

	return v
}

// func (v Vec3) Floor() Vec3 {
// 	return New(math.Floor(v.X), math.Floor(v.Y), math.Floor(v.Z))
// }

func (v Vec3) Floor() Vec3 {
	v.X = math.Floor(v.X)
	v.Y = math.Floor(v.Y)
	v.Z = math.Floor(v.Z)

	return v
}

// func (v Vec3) Abs() Vec3 {
// 	return New(math.Abs(v.X), math.Abs(v.Y), math.Floor(v.Z))
// }

func (v Vec3) Abs() Vec3 {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
	v.Z = math.Abs(v.Z)

	return v
}

// func (v Vec3) Sign() Vec3 {
// 	sign := New(1, 1, 1)
// 	if v.X < 0 {
// 		sign.X = -1
// 	}
// 	if v.Y < 0 {
// 		sign.Y = -1
// 	}
// 	if v.Z < 0 {
// 		sign.Z = -1
// 	}

// 	return sign
// }

func (v Vec3) Sign() Vec3 {
	v.X = 1
	v.Y = 1
	v.Z = 1

	if v.X < 0 {
		v.X = -1
	}
	if v.Y < 0 {
		v.Y = -1
	}
	if v.Z < 0 {
		v.Z = -1
	}

	return v
}

func (v Vec3) XY() vec2.Vec2 {
	return vec2.New(v.X, v.Y)
}

func (v Vec3) Print() {
	fmt.Printf("{%f, %f, %f}\n", v.X, v.Y, v.Z)
}
