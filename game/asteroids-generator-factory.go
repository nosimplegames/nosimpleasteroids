package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type AsteroidsGeneratorFactory struct {
}

const asteroidsGeneratorTimeoutId = "asteroids-generator-timeout"

func (factory AsteroidsGeneratorFactory) Create() *AsteroidsGenerator {
	generator := &AsteroidsGenerator{}

	generator.GenerationPoints = []AsteroidGenerationPoint{
		{
			Position: math.Vector{
				X: 400,
				Y: -40,
			},
			MinRotation: math.DegreesToRads(45),
			MaxRotation: math.DegreesToRads(135),
		},
		{
			Position: math.Vector{
				X: 840,
				Y: 400,
			},
			MinRotation: math.DegreesToRads(135),
			MaxRotation: math.DegreesToRads(225),
		},
		{
			Position: math.Vector{
				X: 400,
				Y: 840,
			},
			MinRotation: math.DegreesToRads(225),
			MaxRotation: math.DegreesToRads(315),
		},
		{
			Position: math.Vector{
				X: -40,
				Y: 400,
			},
			MinRotation: math.DegreesToRads(315),
			MaxRotation: math.DegreesToRads(395),
		},
	}

	engine.GetTimer().AddTimeout(&engine.Timeout{
		Time:     2,
		Callback: generator.GenerateAsteroid,
		IsLoop:   true,
		Id:       asteroidsGeneratorTimeoutId,
	})

	return generator
}

func InitAsteroidsGenerator() {
	AsteroidsGeneratorFactory{}.Create()
}

func StopAsteroidsGenerator() {
	engine.GetTimer().RemoveTimeout(asteroidsGeneratorTimeoutId)
}
