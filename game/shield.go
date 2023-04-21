package game

import (
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type Shield struct {
	engine.Entity
	math.Transformable

	Animation *animations.AnimationList
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
