package engine

import (
	"simple-games.com/asteroids/render"
)

type IEntity interface {
	HandleInput()
	Update()
	Draw(target render.RenderTarget)
	IsAlive() bool
	GetChildren() []IEntity
}
