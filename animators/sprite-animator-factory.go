package animators

import "simple-games.com/asteroids/animations"

type SpriteAnimatorFactory struct {
	SpriteAnimation animations.SpriteAnimation
}

func (factory SpriteAnimatorFactory) Create(target ISprite) SpriteAnimator {
	animator := SpriteAnimator{}
	animator.SpriteAnimation = *factory.SpriteAnimation.Copy().(*animations.SpriteAnimation)
	animator.Target = target
	animator.Target.SetTexture(animator.SpriteAnimation.Texture)

	return animator
}
