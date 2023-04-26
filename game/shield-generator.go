package game

type ShieldGenerator struct {
	IsShieldActivated bool
}

func (generator *ShieldGenerator) ActivateShield() *Shield {
	shield := ShieldFactory{}.Create()
	generator.IsShieldActivated = true

	shield.OnDie.AddCallback(func() {
		generator.IsShieldActivated = false
	})

	return shield
}
