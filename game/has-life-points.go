package game

import "simple-games.com/asteroids/events"

type HasLifePoints struct {
	maxLifePoints int
	lifePoints    int

	OnLifePointsChanged    events.Signal
	OnMaxLifePointsChanged events.Signal
}

func (hasLifePoints HasLifePoints) GetLifePoints() int {
	return hasLifePoints.lifePoints
}

func (hasLifePoints *HasLifePoints) SetMaxLifePoints(lifePoints int) {
	hasLifePoints.maxLifePoints = lifePoints
	hasLifePoints.OnMaxLifePointsChanged.TFire(hasLifePoints.maxLifePoints)
}

func (hasLifePoints *HasLifePoints) SetLifePoints(lifePoints int) {
	hasLifePoints.lifePoints = lifePoints
	hasLifePoints.OnLifePointsChanged.TFire(hasLifePoints.lifePoints)
}

func (hasLifePoints *HasLifePoints) DecreaseLifePoints(amout int) {
	hasLifePoints.SetLifePoints(hasLifePoints.lifePoints - amout)
}

func (hasLifePoints *HasLifePoints) ConnectHealthBar(healthBar IHealthBar) {
	healthBar.SetMaxValue(float64(hasLifePoints.maxLifePoints))
	healthBar.SetValue(float64(hasLifePoints.lifePoints))

	hasLifePoints.OnMaxLifePointsChanged.AddTCallback(func(data events.SignalData) {
		maxValue := data.(int)
		healthBar.SetMaxValue(float64(maxValue))
	})

	hasLifePoints.OnLifePointsChanged.AddTCallback(func(data events.SignalData) {
		value := data.(int)
		healthBar.SetValue(float64(value))
	})
}
