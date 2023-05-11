package main

import (
	"simple-games.com/asteroids/assets"
	"simple-games.com/asteroids/asteroidsscene"
	"simple-games.com/asteroids/engine"
)

func main() {
	entities := engine.GetEntities()
	scene := asteroidsscene.AsteroidsSceneFactory{}.Create()
	entities.AddEntity(scene)

	asteroidsGame := engine.Game{
		BackgroundColor: assets.BackgroundColor,
		Size:            assets.GameSize,
		WindowSize:      assets.GameSize,
	}
	// asteroidsGame.GetDefaultCamera().SetScale(math.Vector{X: 2, Y: 2})
	asteroidsGame.Run()
}
