package game

import (
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type ShieldFactory struct {
}

func (factory ShieldFactory) Create() *Shield {
	shield := &Shield{}
	shield.Animation = factory.GetAnimation()
	engine.GetAnimations().AddAnimation(shield.Animation)
	shield.Origin = math.Vector{X: 16, Y: 16}
	shield.Position = math.Vector{X: 8, Y: 8}
	shield.Size = math.Vector{X: 32, Y: 32}

	engine.GetWorld().AddCollisinable(shield)

	return shield
}

func (factory ShieldFactory) GetAnimation() *animations.AnimationList {
	shieldAnimation := GetAssets().ShieldAnimation.Copy()

	animation := animations.AnimationLoopFactory{
		Animation: shieldAnimation,
		LoopCount: 4,
	}.Create()

	return &animation
}
