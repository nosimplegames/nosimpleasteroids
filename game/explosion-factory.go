package game

import (
	"simple-games.com/asteroids/animators"
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
	explosion := &engine.Sprite{}

	explosion.SetPosition(factory.Position)
	animator := factory.GetAnimator(explosion)
	animator.Animation.OnStop.AddCallback(explosion.Die)

	engine.GetAnimations().AddAnimation(animator)
	engine.GetEntities().AddEntity(explosion)
}

func (factory ExplosionFactory) GetAnimator(explosion *engine.Sprite) *animators.SpriteAnimator {
	assets := GetAssets()
	animatorFactory := assets.SmallExplosionAnimatorFactory
	mustUseBigExplosionAnimation := factory.ExplosionSize == BigExplosion

	if mustUseBigExplosionAnimation {
		animatorFactory = assets.BigExplosionAnimatorFactory
	}

	animation := animatorFactory.Create(explosion)
	return animation
}
