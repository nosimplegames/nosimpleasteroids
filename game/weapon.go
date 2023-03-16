package game

import "simple-games.com/asteroids/utils"

type Weapon struct {
	TimeSinceLastShoot float64
	TimeBetweenShoots  float64
}

func (weapon Weapon) CanShoot() bool {
	return weapon.TimeSinceLastShoot >= weapon.TimeBetweenShoots
}

func (weapon *Weapon) PrepareForNextShoot() {
	weapon.TimeSinceLastShoot += utils.FrameTime
}
