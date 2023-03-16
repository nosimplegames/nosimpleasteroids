package particles

import "simple-games.com/asteroids/utils"

type Particle struct {
	LifeTime float64
}

func (particle *Particle) Live() {
	particle.LifeTime -= utils.FrameTime
}

func (particle Particle) IsAlive() bool {
	return particle.LifeTime > 0
}

func (particle *Particle) Die() {
	particle.LifeTime = 0
}
