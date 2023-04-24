package game

import (
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type Shield struct {
	engine.Entity

	Animation *animations.AnimationList
	Size      math.Vector

	OnDestroy events.Signal
}

func (shield Shield) IsAlive() bool {
	return shield.Animation.IsAlive()
}

func (shield Shield) Draw(target render.RenderTarget, transform math.Transform) {
	spriteAnimation := shield.Animation.CurrentAnimation.(*animations.SpriteAnimation)
	sprite := spriteAnimation.GetCurrentSprite()
	sprite.Target = target
	sprite.Transform = transform

	sprite.Render()
}

func (shield Shield) GetSize() math.Vector {
	return shield.Size
}

func (shield Shield) GetCollisionMask() string {
	return SpaceshipShieldCollisionMask
}

func (shield Shield) CanCollide() bool {
	return shield.IsAlive()
}

func (shield Shield) CanCollideWith(collisionMask string) bool {
	return collisionMask == EnergyBlastCollisionMask
}

func (shield *Shield) OnCollision(collisionMask string) {
	shield.Animation.Stop()
	shield.OnDestroy.Fire()
}
