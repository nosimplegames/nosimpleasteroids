package engine

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
	"simple-games.com/asteroids/utils"
)

type Entity struct {
	Children []IEntity

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
}

func (entity Entity) GetChildren() []IEntity {
	return entity.Children
}

func (entity *Entity) RemoveDeadChildren() {
	utils.RemoveDead(&entity.Children)
}
