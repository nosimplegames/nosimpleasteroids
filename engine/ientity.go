package engine

import (
	"simple-games.com/asteroids/math"
)

type IEntity interface {
	ILiving
	IDrawable
	math.ITransformable

	HandleInput()
	Update()

	GetChildren() []IEntity
	RemoveDeadChildren()

	GetAncestorsTransform() math.Transform
	GetSize() math.Vector

	SetParent(IEntity)
	GetParent() IEntity

	GetId() string
	GetType() string
}
