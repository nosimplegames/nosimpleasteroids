package engine

import "simple-games.com/asteroids/utils"

type TimeoutCallback func()

type Timeout struct {
	Time        float64
	CurrentTime float64
	Callback    TimeoutCallback
	IsLoop      bool
}

func (timeout *Timeout) Update() {
	timeout.CurrentTime += utils.FrameTime
}

func (timeout Timeout) IsAlive() bool {
	return timeout.CurrentTime < timeout.Time || timeout.IsLoop
}

func (timeout *Timeout) Exec() {
	mustExecute := timeout.CurrentTime >= timeout.Time
	if !mustExecute {
		return
	}

	timeout.Callback()

	if timeout.IsLoop {
		timeout.Reset()
	}
}

func (timeout *Timeout) Reset() {
	timeout.CurrentTime = 0
}
