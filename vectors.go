package vec

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

// some predefines
type Vi2d = V2d[int]
type Vf2d = V2d[float64]

type Number interface {
	constraints.Integer | constraints.Float
}

type V2d[Type Number] struct {
	X, Y Type
}

func (this *V2d[Type]) Add(v V2d[Type]) {
	this.X += v.X
	this.Y += v.Y
}

func (this *V2d[Type]) Subtract(v V2d[Type]) {
	this.X -= v.X
	this.Y -= v.Y
}

func (this *V2d[Type]) Multiply(v V2d[Type]) {
	this.X *= v.X
	this.Y *= v.Y
}

func (this *V2d[Type]) Divide(v V2d[Type]) {
	this.X /= v.X
	this.Y /= v.Y
}

func NewVec[Type Number](x, y Type) *V2d[Type] {
	return &V2d[Type]{X: x, Y: y}
}

// useful arithmetics
func Add[Type Number](lhs, rhs V2d[Type]) (answer V2d[Type]) {
	answer.X = lhs.X + rhs.X
	answer.Y = lhs.Y + rhs.Y
	return
}

func Subtract[Type Number](lhs, rhs V2d[Type]) (answer V2d[Type]) {
	answer.X = lhs.X - rhs.X
	answer.Y = lhs.Y - rhs.Y
	return
}

func Multiply[Type Number](lhs, rhs V2d[Type]) (answer V2d[Type]) {
	answer.X = lhs.X * rhs.X
	answer.Y = lhs.Y * rhs.Y
	return
}

func Divide[Type Number](lhs, rhs V2d[Type]) (answer V2d[Type]) {
	answer.X = lhs.X / rhs.X
	answer.Y = lhs.Y / rhs.Y
	return
}

func Distance[Type Number](lhs, rhs V2d[Type]) (answer float64) {
	X := lhs.X - rhs.X
	Y := lhs.Y - rhs.Y

	answer = math.Sqrt(float64(X*X + Y*Y))
	return
}

// general utility

func (this *V2d[Type]) String() string {
	return fmt.Sprintf("(%s, %s)", fmt.Sprint(this.X), fmt.Sprint(this.Y))
}
