package game

type EnergyBlast struct {
	Inert
}

func (blast EnergyBlast) GetCollisionMask() string {
	return EnergyBlastCollisionMask
}

func (blast EnergyBlast) CanCollide() bool {
	return true
}

func (blast EnergyBlast) CanCollideWith(collisionMask string) bool {
	return collisionMask == SpaceshipCollisionMask ||
		collisionMask == SpaceshipShieldCollisionMask
}

func (blast EnergyBlast) OnCollision(collisionMask string) {
}

func (blast EnergyBlast) IsAlive() bool {
	return true
}
