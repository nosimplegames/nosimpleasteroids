package engine

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
	"simple-games.com/asteroids/utils"
)

type Entity struct {
	math.Transformable

	Children []IEntity

	Parent IEntity
	IsDead bool
}

func (entity *Entity) HandleInput() {
}

func (entity *Entity) Update() {
}

func (entity Entity) IsAlive() bool {
	return !entity.IsDead
}

func (entity *Entity) Die() {
	entity.IsDead = true
}

func (entity Entity) Draw(target render.RenderTarget, combinedTransform math.Transform) {
}

func (entity *Entity) AddChild(child IEntity) {
	entity.Children = append(entity.Children, child)
	child.SetParent(entity)
}

func (entity Entity) GetChildren() []IEntity {
	return entity.Children
}

func (entity *Entity) RemoveDeadChildren() {
	utils.RemoveDead(&entity.Children)
}

func (entity *Entity) SetParent(parent IEntity) {
	entity.Parent = parent
}

func (entity Entity) GetParent() IEntity {
	return entity.Parent
}

func (entity Entity) GetPosition() math.Vector {
	fullTransform := entity.GetAncestorsTransform()
	transformedX, transformedY := fullTransform.Apply(entity.Position.X, entity.Position.Y)

	return math.Vector{
		X: transformedX,
		Y: transformedY,
	}
}

func (entity Entity) GetAncestorsTransform() math.Transform {
	transform := math.Transform{}

	hasParent := entity.Parent != nil

	if hasParent {
		parentTransform := entity.Parent.GetAncestorsTransform()
		transform.Concat(parentTransform)
		transform.Concat(entity.Parent.GetTransform())
	}

	return transform
}
