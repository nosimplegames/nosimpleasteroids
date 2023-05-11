package engine

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/utils"
)

type Entity struct {
	math.Transformable
	Living
	Drawable

	Id   string
	Type string

	Children []IEntity
	Parent   IEntity

	Size math.Vector
}

func (entity Entity) GetId() string {
	return entity.Id
}

func (entity Entity) GetType() string {
	return entity.Type
}

func (entity *Entity) HandleInput() {
}

func (entity *Entity) Update() {
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
	entityPosition := entity.Transformable.GetPosition()
	transformedX, transformedY := fullTransform.Apply(entityPosition.X, entityPosition.Y)

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

func (entity Entity) GetSize() math.Vector {
	return entity.Size
}
