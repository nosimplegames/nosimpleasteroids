package engine

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
	"simple-games.com/asteroids/utils"
)

type Entities struct {
	Entities []IEntity
}

func (entities *Entities) Update() {
	for _, entity := range entities.Entities {
		entities.EntityHandleInput(entity)
	}

	for _, entity := range entities.Entities {
		entities.EntityUpdate(entity)
	}

	entities.RemoveDead()
}

func (entities Entities) EntityHandleInput(entity IEntity) {
	entity.HandleInput()

	for _, child := range entity.GetChildren() {
		entities.EntityHandleInput(child)
	}
}

func (entities Entities) EntityUpdate(entity IEntity) {
	entity.Update()

	for _, child := range entity.GetChildren() {
		entities.EntityUpdate(child)
	}

	entity.RemoveDeadChildren()
}

func (entities *Entities) RemoveDead() {
	utils.RemoveDead(&entities.Entities)
}

func (entities Entities) Draw(target render.RenderTarget) {
	for _, entity := range entities.Entities {
		entities.DrawEntity(entity, target, math.Transform{})
	}
}

func (entities Entities) DrawEntity(entity IEntity, target render.RenderTarget, parentTransform math.Transform) {
	transform := entity.GetTransform()
	transform.Concat(parentTransform)
	entity.Draw(target, transform)

	for _, child := range entity.GetChildren() {
		entities.DrawEntity(child, target, transform)
	}
}

func (entities *Entities) AddEntity(entity IEntity) {
	entities.Entities = append(entities.Entities, entity)
}

var entities *Entities = nil

func GetEntities() *Entities {
	if entities == nil {
		entities = &Entities{
			Entities: []IEntity{},
		}
	}

	return entities
}
