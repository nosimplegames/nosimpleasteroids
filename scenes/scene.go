package scenes

import (
	"simple-games.com/asteroids/engine"
)

type Scene struct {
	engine.Entity

	Acts    Acts
	HasInit bool
}

func (scene *Scene) Update() {
	shouldInitAct := !scene.HasInit && !scene.Acts.IsEmpty()

	if shouldInitAct {
		scene.PrepareAct(scene.Acts.Dequeue())
		scene.HasInit = true
	}
}

func (scene *Scene) PrepareAct(act *IAct) {
	isActValid := act != nil

	if !isActValid {
		return
	}

	(*act).Openning(scene)
	scene.AddChild(*act)
	(*act).OnDie(func() {
		scene.PrepareAct(scene.Acts.Dequeue())
	})
}
