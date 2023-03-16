package game

import (
	"math"
)

type Engine struct {
	MinSpeed      float64
	MaxSpeed      float64
	CurrentSpeed  float64
	Acceleration  float64
	Deaceleration float64

	IsAccelerating bool
}

func (engine *Engine) Accelerate() {
	engine.IsAccelerating = true
	engine.CurrentSpeed += engine.Acceleration
	engine.CurrentSpeed = math.Min(engine.CurrentSpeed, engine.MaxSpeed)
}

func (engine *Engine) Deacelerate() {
	engine.CurrentSpeed -= engine.Deaceleration
	engine.CurrentSpeed = math.Max(engine.CurrentSpeed, engine.MinSpeed)
}
