package animations

import "simple-games.com/asteroids/events"

type AnimationState int

const (
	AnimationRunning AnimationState = iota
	AnimationPaused
	AnimationStopped
)

type Animation struct {
	State AnimationState
	events.EventTarget
}

func (animation Animation) IsAlive() bool {
	return animation.State != AnimationStopped
}

func (animation Animation) IsRunning() bool {
	return animation.State == AnimationRunning
}

func (animation *Animation) Stop() {
	animation.State = AnimationStopped
	animation.DispatchEvent(events.Event{
		Type: AnimationStoppedEvent,
	})
}
