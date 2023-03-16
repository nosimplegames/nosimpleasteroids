package engine

import (
	"simple-games.com/asteroids/render"
)

type Entity struct {
	Children []IEntity
}

func (entity *Entity) HandleInput() {
}

func (entity *Entity) Update() {
}

func (entity Entity) IsAlive() bool {
	return true
}

func (entity Entity) Draw(target render.RenderTarget) {
}

func (entity *Entity) AddChild(child IEntity) {
	entity.Children = append(entity.Children, child)
}

func (entity Entity) GetChildren() []IEntity {
	return entity.Children
}
