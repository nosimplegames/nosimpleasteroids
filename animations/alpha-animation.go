package animations

import "github.com/hajimehoshi/ebiten"

type AlphaAnimation struct {
	NumberAnimation
}

func (animation AlphaAnimation) GetColorM() ebiten.ColorM {
	colorM := ebiten.ColorM{}

	colorM.SetElement(3, 3, animation.CurrentValue)

	return colorM
}

func (animation AlphaAnimation) Copy() IAnimation {
	copy := &AlphaAnimation{
		NumberAnimation: NumberAnimation{
			InitialValue: animation.InitialValue,
			TargetValue:  animation.TargetValue,
			CurrentValue: animation.InitialValue,
			FrameStep:    animation.FrameStep,
		},
	}

	copy.State = AnimationRunning

	return copy
}

func (animation AlphaAnimation) Reverse() IAnimation {
	copy := &AlphaAnimation{
		NumberAnimation: NumberAnimation{
			InitialValue: animation.TargetValue,
			TargetValue:  animation.InitialValue,
			CurrentValue: animation.TargetValue,
			FrameStep:    -animation.FrameStep,
		},
	}

	copy.State = AnimationRunning

	return copy
}
