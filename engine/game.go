package engine

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"simple-games.com/asteroids/math"
)

type Game struct {
	BackgroundColor color.Color
	Cameras         []*Camera
	WindowSize      math.Vector
	Size            math.Vector
}

func (game *Game) Update(*ebiten.Image) error {
	GetUpdatables().Update()
	GetEntities().Update()
	GetWorld().Update()
	GetAnimations().Update()

	var err error = nil

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		err = errors.New("")
	}

	return err
}

func (game Game) Draw(screen *ebiten.Image) {
	for _, camera := range game.Cameras {
		cameraTexture, _ := ebiten.NewImage(int(game.Size.X), int(game.Size.Y), ebiten.FilterDefault)
		cameraTexture.Fill(game.BackgroundColor)
		GetEntities().Draw(cameraTexture)

		camera.SetTexture(cameraTexture)
		camera.Draw(screen, camera.GetTransform())
	}
}

func (game Game) Layout(_, _ int) (int, int) {
	return int(game.WindowSize.X), int(game.WindowSize.Y)
}

func (game *Game) Run() {
	ebiten.SetWindowSize(int(game.WindowSize.X), int(game.WindowSize.Y))
	game.SetDefaultBackgroundColor()
	game.InitDefaultCamera()
	ebiten.RunGame(game)
}

func (game *Game) SetDefaultBackgroundColor() {
	hasBackgroundColor := game.BackgroundColor != nil

	if !hasBackgroundColor {
		game.BackgroundColor = color.Black
	}
}

func (game *Game) InitDefaultCamera() {
	hasCameras := len(game.Cameras) > 0

	if hasCameras {
		return
	}

	defaultCamera := &Camera{}
	defaultCamera.RenderingBox = math.Box{
		Position: game.Size.By(0.5),
		Size:     game.Size,
	}
	defaultCamera.SetViewport(math.Box{
		Position: game.WindowSize.By(0.5),
		Size:     game.WindowSize,
	})
	game.Cameras = append(game.Cameras, defaultCamera)
}

func (game *Game) GetDefaultCamera() *Camera {
	game.InitDefaultCamera()

	return game.Cameras[0]
}
