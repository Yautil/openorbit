package components

type SpeedInterpolationComponent struct {
	Moving             bool
	MaxSpeed           float32
	CurSpeed           float32
	InterpolationDelta float32
	SlowdownMultiplier float32
}
