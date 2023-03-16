package engine

import "simple-games.com/asteroids/utils"

type Timer struct {
	Timeouts []ITimeout
}

func (timer *Timer) Update() {
	for _, timeout := range timer.Timeouts {
		timeout.Update()
	}

	for _, timeout := range timer.Timeouts {
		timeout.Exec()
	}

	utils.RemoveDead(&timer.Timeouts)
}

func (timer *Timer) AddTimeout(timeout ITimeout) {
	timer.Timeouts = append(timer.Timeouts, timeout)
}

var timer *Timer = nil

func GetTimer() *Timer {
	needToInitTimer := timer == nil

	if needToInitTimer {
		timer = &Timer{}
		GetUpdatables().AddUpdatable(timer)
	}

	return timer
}
