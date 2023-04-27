package game

import (
	"fmt"
	"math/rand"

	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
)

type AsteroidsGenerator struct {
	engine.Living
	Weapon

	GenerationPoints                         []AsteroidGenerationPoint
	GeneratedAsteroidsSinceLastTimeReduction int

	OnStop events.Signal
}

func (generator *AsteroidsGenerator) GenerateAsteroid() {
	entities := engine.GetEntities()

	asteroid := generator.GenerateAsteroidFactory().Create()
	entities.AddEntity(asteroid)

	generator.GeneratedAsteroidsSinceLastTimeReduction++
	generator.TimeSinceLastShoot = 0
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

func (generator *AsteroidsGenerator) Update() {
	mustGenerateAsteroid := generator.CanShoot() && generator.TimeBetweenShoots > 0

	if mustGenerateAsteroid {
		generator.GenerateAsteroid()

		mustReduceGenerationTime := generator.GeneratedAsteroidsSinceLastTimeReduction > asteroidsGenerator.GetRequiredAsteroidsToReduceTime()

		if mustReduceGenerationTime {
			generator.TimeBetweenShoots -= 0.2
			generator.GeneratedAsteroidsSinceLastTimeReduction = 0
			fmt.Println(generator.TimeBetweenShoots)

			mustStop := generator.TimeBetweenShoots <= 0

			if mustStop {
				generator.DieAfterAsteroids()
			}
		}
	}

	generator.PrepareForNextShoot()
}

func (generator AsteroidsGenerator) GetRequiredAsteroidsToReduceTime() int {
	if generator.TimeBetweenShoots > 1 {
		return 5
	}

	return 10
}

func (generator *AsteroidsGenerator) DieAfterAsteroids() {
	generator.TimeBetweenShoots = 0
	seeker := &AsteroidsSeeker{}
	seeker.OnNoMoreAsteroids.AddCallback(generator.Die)
	engine.GetUpdatables().AddUpdatable(seeker)
}
