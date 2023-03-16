package game

import (
	"github.com/hajimehoshi/ebiten"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type Player struct {
	engine.Entity
	Spaceship

	HealthBar *PlayerHealthBar
}

func (player *Player) HandleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		player.Propulsor.Accelerate()
	} else {
		player.Propulsor.Deacelerate()
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.LeftRotator.Accelerate()
	} else {
		player.LeftRotator.Deacelerate()
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.RightRotator.Accelerate()
	} else {
		player.RightRotator.Deacelerate()
	}

	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		player.Turret.Shoot(player.Position, player.Rotation)
	}

	player.Turret.PrepareForNextShoot()
}

func (player *Player) Update() {
	player.Rotate(-player.LeftRotator.CurrentSpeed)
	player.Rotate(player.RightRotator.CurrentSpeed)

	movementVector := math.MovementVector{
		Rotation: player.Rotation,
		Speed:    player.Propulsor.CurrentSpeed,
	}.Calculate()
	player.Move(movementVector)
}

func (player Player) IsAlive() bool {
	return player.Spaceship.IsAlive()
}

func (player Player) Draw(target render.RenderTarget) {
	player.Spaceship.Draw(target)
}
