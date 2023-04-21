package animations

type NumberAnimation struct {
	Animation

	InitialValue float64
	TargetValue  float64
	FrameStep    float64
	CurrentValue float64
}

func (animation *NumberAnimation) Update() {
	animation.CurrentValue += animation.FrameStep

	if animation.HasFinished() {
		animation.CurrentValue = animation.TargetValue
		animation.Stop()
	}
}

func (animation NumberAnimation) HasFinished() bool {
	targetGreaterThanInitial := animation.TargetValue > animation.InitialValue

	if targetGreaterThanInitial {
		return animation.CurrentValue >= animation.TargetValue
	}

	return animation.CurrentValue <= animation.TargetValue
}
