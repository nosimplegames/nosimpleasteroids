package assets

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
	"simple-games.com/asteroids/render"
)

func LoadTexture(textureFileName string) render.Texture {
	textureFile, err := os.Open(textureFileName)

	if err != nil {
		panic(err)
	}

	image, _, err := image.Decode(textureFile)

	if err != nil {
		panic(err)
	}

	texture, err := ebiten.NewImageFromImage(image, ebiten.FilterDefault)

	if err != nil {
		panic(err)
	}

	return texture
}
