package components

import (
	"engo.io/engo"
)

type KeyboardMovementComponent struct {
	Moving             bool
	MaxSpeed           float32
	CurSpeed           engo.Point
	InterpolationDelta float32
	SlowdownMultiplier float32
	DirectionVector    engo.Point
	MovementVector     engo.Point
}
