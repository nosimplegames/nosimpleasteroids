package animators

import "simple-games.com/asteroids/animations"

type SpriteAnimator struct {
	animations.SpriteAnimation
	Target ISprite
}

func (animator *SpriteAnimator) Update() {
	animator.SpriteAnimation.Update()

	if animator.IsRunning() {
		animator.Target.SetRect(animator.GetCurrentRect())
	}
}

func (animator SpriteAnimator) Copy() animations.IAnimation {
	return &SpriteAnimator{
		SpriteAnimation: *animator.SpriteAnimation.Copy().(*animations.SpriteAnimation),
		Target:          animator.Target,
	}
}
func (animator SpriteAnimator) Reverse() animations.IAnimation {
	return &SpriteAnimator{
		SpriteAnimation: *animator.SpriteAnimation.Reverse().(*animations.SpriteAnimation),
		Target:          animator.Target,
	}
}
