package engine

import (
	"simple-games.com/asteroids/math"
)

type IEntity interface {
	ILiving
	IDrawable

	HandleInput()
	Update()

	GetChildren() []IEntity
	RemoveDeadChildren()

	GetTransform() math.Transform
	GetAncestorsTransform() math.Transform

	SetParent(IEntity)
	GetParent() IEntity

	GetId() string
	GetType() string
}
