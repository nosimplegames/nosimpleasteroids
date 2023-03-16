package animations

type AnimationLoopFactory struct {
	Animation        IAnimation
	LoopCount        int
	IncludeBackwards bool
}

func (factory AnimationLoopFactory) Create() AnimationList {
	animationList := AnimationList{}
	animationList.State = AnimationRunning

	for i := 0; i < factory.LoopCount; i++ {
		animationList.AddAnimation(factory.Animation.Copy())

		if factory.IncludeBackwards {
			animationList.AddAnimation(factory.Animation.Reverse())
		}
	}

	return animationList
}
