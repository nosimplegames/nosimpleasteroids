package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type Player struct {
	Spaceship

	HealthBar           *PlayerHealthBar
	IsControllerEnabled bool
}

func (player *Player) HandleInput() {
	if !player.IsControllerEnabled {
		return
	}

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

	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		player.ActivateShield()
	}

	player.Turret.PrepareForNextShoot()
}

func (player *Player) Update() {
	if !player.IsControllerEnabled {
		return
	}

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

func (player Player) Draw(target render.RenderTarget, transform math.Transform) {
	player.Spaceship.Draw(target, transform)
}
