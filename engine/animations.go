package engine

import (
	eanimations "simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/utils"
)

type Animations struct {
	Animations []eanimations.IAnimation
}

func (animations *Animations) AddAnimation(animation eanimations.IAnimation) {
	animations.Animations = append(animations.Animations, animation)
}

func (animations *Animations) Update() {
	for _, animation := range animations.Animations {
		animation.Update()
	}

	utils.RemoveDead(&animations.Animations)
}

var animations *Animations = nil

func GetAnimations() *Animations {
	needToInitAnimations := animations == nil

	if needToInitAnimations {
		animations = &Animations{}
	}

	return animations
}
