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
	shield.Origin = math.Vector{X: 16, Y: 16}
	shield.Position = math.Vector{X: 8, Y: 8}
	shield.Size = math.Vector{X: 32, Y: 32}

	animation := factory.GetAnimation(shield)
	animation.OnStop.AddCallback(shield.Die)

	engine.GetAnimations().AddAnimation(animation)
	engine.GetWorld().AddCollisinable(shield)

	return shield
}

func (factory ShieldFactory) GetAnimation(shield *Shield) *animations.AnimationList {
	shieldAnimation := GetAssets().SpaceshipShieldAnimatorFactory.Create(shield)

	animation := animations.AnimationLoopFactory{
		Animation: shieldAnimation,
		LoopCount: 4,
	}.Create()

	return &animation
}
