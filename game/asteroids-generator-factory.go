package game

import (
	"simple-games.com/asteroids/assets"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type AsteroidsGeneratorFactory struct {
}

const asteroidsGeneratorTimeoutId = "asteroids-generator-timeout"

func (factory AsteroidsGeneratorFactory) Create() *AsteroidsGenerator {
	generator := &AsteroidsGenerator{}
	generator.TimeBetweenShoots = 2

	offset := 20.0
	generator.GenerationPoints = []AsteroidGenerationPoint{
		{
			Position: math.Vector{
				X: assets.GameSize.X * 0.5,
				Y: -offset,
			},
			MinRotation: math.DegreesToRads(45),
			MaxRotation: math.DegreesToRads(135),
		},
		{
			Position: math.Vector{
				X: assets.GameSize.X + offset,
				Y: assets.GameSize.Y * 0.5,
			},
			MinRotation: math.DegreesToRads(135),
			MaxRotation: math.DegreesToRads(225),
		},
		{
			Position: math.Vector{
				X: assets.GameSize.X * 0.5,
				Y: assets.GameSize.Y + offset,
			},
			MinRotation: math.DegreesToRads(225),
			MaxRotation: math.DegreesToRads(315),
		},
		{
			Position: math.Vector{
				X: -offset,
				Y: assets.GameSize.Y * 0.5,
			},
			MinRotation: math.DegreesToRads(315),
			MaxRotation: math.DegreesToRads(395),
		},
	}

	engine.GetUpdatables().AddUpdatable(generator)

	return generator
}

var asteroidsGenerator *AsteroidsGenerator = nil

func GetAsteroidsGenerator() *AsteroidsGenerator {
	mustInitGenerator := asteroidsGenerator == nil

	if mustInitGenerator {
		asteroidsGenerator = AsteroidsGeneratorFactory{}.Create()
	}

	return asteroidsGenerator
}

func StopAsteroidsGenerator() {
	engine.GetTimer().RemoveTimeout(asteroidsGeneratorTimeoutId)
}
