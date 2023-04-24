package engine

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type IEntity interface {
	HandleInput()
	Update()
	Draw(target render.RenderTarget, combinedTransform math.Transform)
	IsAlive() bool
	Die()
	GetChildren() []IEntity
	RemoveDeadChildren()

	GetTransform() math.Transform
	GetAncestorsTransform() math.Transform

	SetParent(IEntity)
	GetParent() IEntity
}
