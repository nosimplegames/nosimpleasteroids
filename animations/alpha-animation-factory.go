package animations

type AlphaAnimationFactory struct {
	InitialValue float64
	TargetValue  float64
	Duration     float64
}

func (factory AlphaAnimationFactory) Create() AlphaAnimation {
	animation := AlphaAnimation{}

	animation.State = AnimationRunning
	animation.NumberAnimation = NumberAnimationFactory{
		InitialValue: factory.InitialValue,
		TargetValue:  factory.TargetValue,
		Duration:     factory.Duration,
	}.Create()

	return animation
}
