package asteroidsscene

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/game"
)

func DoesPlayerExist() bool {
	return engine.GetEntities().IsThereEntitiesOfType(game.SpaceshipEntityType)
}
