package game

import (
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/events"
)

type ShieldGenerator struct {
	IsShieldActivated bool
}

func (generator *ShieldGenerator) ActivateShield() *Shield {
	shield := ShieldFactory{}.Create()
	generator.IsShieldActivated = true

	shield.Animation.AddEventListener(events.EventListener{
		EventType: animations.AnimationStoppedEvent,
		Callback: func(event events.Event) {
			generator.IsShieldActivated = false
		},
	})

	return shield
}
