package game

import (
	"math/rand"

	"simple-games.com/asteroids/engine"
)

type AsteroidsGenerator struct {
	GenerationPoints []AsteroidGenerationPoint
}

func (generator *AsteroidsGenerator) GenerateAsteroid() {
	entities := engine.GetEntities()

	asteroid := generator.GenerateAsteroidFactory().Create()
	entities.AddEntity(asteroid)
}

func (generator AsteroidsGenerator) GenerateAsteroidFactory() AsteroidFactory {
	generationPoint := generator.GetRandomGenerationPoint()
	asteriodSize := generator.GetRandomAsteroidSize()

	factory := AsteroidFactory{
		Size:     asteriodSize,
		Position: generationPoint.Position,
		Rotation: generationPoint.GetRandomRotation(),
	}

	return factory
}

func (generator AsteroidsGenerator) GetRandomGenerationPoint() AsteroidGenerationPoint {
	randomIndex := rand.Intn(len(generator.GenerationPoints))

	return generator.GenerationPoints[randomIndex]
}

func (generator AsteroidsGenerator) GetRandomAsteroidSize() AsteroidSize {
	asteroidSizes := []AsteroidSize{
		BigAsteroid,
		NormalAsteroid,
		SmallAsteroid,
	}
	asteroidSizeIndex := rand.Intn(len(asteroidSizes))

	return asteroidSizes[asteroidSizeIndex]
}
