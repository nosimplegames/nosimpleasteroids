package engine

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/ebitenutil"
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
	if !entity.IsRunning() {
		return
	}

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
		EntityDrawer{
			Entity:    entity,
			Transform: math.Transform{},
			Target:    target,
		}.Draw()
	}
	ebitenutil.DebugPrint(target, strconv.Itoa(len(entities.Entities)))
}

func (entities *Entities) AddEntity(entity IEntity) {
	entities.Entities = append(entities.Entities, entity)
}

func (entities Entities) IsThereEntitiesOfType(entityType string) bool {
	for _, entity := range entities.Entities {
		seeker := EntitySeeker{
			Entity: entity,
		}

		if seeker.IsThereOfType(entityType) {
			return true
		}
	}

	return false
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
