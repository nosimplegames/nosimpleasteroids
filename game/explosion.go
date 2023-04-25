package game

import (
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type Explosion struct {
	engine.Entity

	Animation *animations.SpriteAnimation
}

func (explosion Explosion) IsAlive() bool {
	return explosion.Animation.IsAlive()
}

func (explosion *Explosion) Draw(target render.RenderTarget, transform math.Transform) {
	sprite := explosion.Animation.GetCurrentSprite()
	sprite.Target = target
	sprite.Transform = transform

	sprite.Render()
}
