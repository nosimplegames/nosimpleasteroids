package animators

import "simple-games.com/asteroids/animations"

type AlphaAnimator struct {
	Animator[IColorable, *animations.NumberAnimation]
}
