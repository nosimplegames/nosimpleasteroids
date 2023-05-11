package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"simple-games.com/asteroids/math"
)

type Player struct {
	Spaceship
	IsControllerEnabled bool
	DidAccelerate       bool
	DidRotate           bool
	Direction           float64
}

func (player *Player) HandleInput() {
	if !player.IsControllerEnabled {
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		player.Propulsor.Accelerate()
		player.DidAccelerate = true
	} else {
		player.Propulsor.Deacelerate()
		player.DidAccelerate = false
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
		player.Turret.Shoot(player.GetPosition(), player.GetRotation())
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

	if player.DidAccelerate {
		player.Direction = player.GetRotation()
	}

	movementVector := math.MovementVector{
		Rotation: player.Direction,
		Speed:    player.Propulsor.CurrentSpeed,
	}.Calculate()
	player.Move(movementVector)
}
