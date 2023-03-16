package game

import (
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type ExplosionSize int

const (
	SmallExplosion ExplosionSize = iota
	BigExplosion   ExplosionSize = iota
)

type ExplosionFactory struct {
	Position      math.Vector
	ExplosionSize ExplosionSize
}

func (factory ExplosionFactory) Create() {
	explosion := &Explosion{}

	explosion.Position = factory.Position
	explosion.Animation = factory.GetAnimation()
	engine.GetAnimations().AddAnimation(explosion.Animation)
	engine.GetEntities().AddEntity(explosion)
}

func (factory ExplosionFactory) GetAnimation() *animations.SpriteAnimation {
	assets := GetAssets()
	animation := assets.SmallExplosionAnimation
	mustUseBigExplosionAnimation := factory.ExplosionSize == BigExplosion

	if mustUseBigExplosionAnimation {
		animation = assets.BigExplosionAnimation
	}

	spriteAnimation := animation.Copy()
	return spriteAnimation.(*animations.SpriteAnimation)
}
