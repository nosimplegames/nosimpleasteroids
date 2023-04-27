package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
)

type AsteroidsSeeker struct {
	engine.Living
	OnNoMoreAsteroids events.Signal
}

func (seeker *AsteroidsSeeker) Update() {
	isThereAsteroids := engine.GetEntities().IsThereEntitiesOfType(AsteroidEntityType)

	if !isThereAsteroids {
		seeker.OnNoMoreAsteroids.Fire()
		seeker.Die()
	}
}
