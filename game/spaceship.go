package game

import (
	"github.com/hajimehoshi/ebiten"
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type OnDestroyCallback func()

type Spaceship struct {
	math.Transformable
	events.EventTarget

	Propulsor    Engine
	RightRotator Engine
	LeftRotator  Engine

	Turret IWeapon

	LifePoints int
	OnDestroy  OnDestroyCallback

	RespawningAnimation *animations.AnimationList
}

func (spaceship Spaceship) Draw(screen *ebiten.Image) {
	render.Sprite{
		Target:    screen,
		Transform: spaceship.GetTransform(),
		Texture:   GetAssets().SpaceshipTexture,
		ColorM:    spaceship.GetColorM(),
	}.Render()
}

func (spaceship Spaceship) IsRespawning() bool {
	return spaceship.RespawningAnimation != nil && spaceship.RespawningAnimation.IsAlive()
}

func (spaceship Spaceship) GetColorM() ebiten.ColorM {
	if spaceship.IsRespawning() {
		alphaAnimation := spaceship.RespawningAnimation.CurrentAnimation.(*animations.AlphaAnimation)
		return alphaAnimation.GetColorM()
	}

	return ebiten.ColorM{}
}

func (spaceship Spaceship) IsAlive() bool {
	return spaceship.LifePoints > 0
}

func (spaceship Spaceship) GetSize() math.Vector {
	return math.Vector{
		X: 12,
		Y: 12,
	}
}

func (spaceship Spaceship) CanCollide() bool {
	return !spaceship.IsRespawning()
}

func (spaceship Spaceship) CanCollideWith(collisionMask string) bool {
	return collisionMask == AsteroidCollisionMask
}

func (spaceship Spaceship) GetCollisionMask() string {
	return SpaceshipCollisionMask
}

func (spaceship *Spaceship) OnCollision(collisionMask string) {
	spaceship.LifePoints -= 1
	spaceship.DispatchEvent(events.Event{
		Type: LifePointsChanged,
		Data: spaceship.LifePoints,
	})

	gotDestroyed := spaceship.LifePoints == 0

	if gotDestroyed {
		spaceship.OnDestroy()
	}
}
