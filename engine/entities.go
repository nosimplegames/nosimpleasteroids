package engine

import (
	"github.com/hajimehoshi/ebiten"
	"simple-games.com/asteroids/utils"
)

type Entities struct {
	Entities []IEntity
}

func (entities *Entities) Update() {
	for _, entity := range entities.Entities {
		entity.HandleInput()

		for _, child := range entity.GetChildren() {
			child.HandleInput()
		}
	}

	for _, entity := range entities.Entities {
		entity.Update()

		for _, child := range entity.GetChildren() {
			child.Update()
		}
	}

	entities.RemoveDead()
}

func (entities *Entities) RemoveDead() {
	utils.RemoveDead(&entities.Entities)
}

func (entities *Entities) Draw(screen *ebiten.Image) {
	for _, entity := range entities.Entities {
		entity.Draw(screen)

		for _, child := range entity.GetChildren() {
			child.Draw(screen)
		}
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
