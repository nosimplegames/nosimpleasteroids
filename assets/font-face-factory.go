package assets

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type FontFaceFactory struct {
	Size         float64
	DPI          float64
	FontFileName string
}

func (factory FontFaceFactory) Create() font.Face {
	font := GetFontManager().GetFont(factory.FontFileName)
	faceOptions := factory.GetFaceOptions()

	face, err := opentype.NewFace(font, faceOptions)

	if err != nil {
		panic(err)
	}

	return face
}

func (factory FontFaceFactory) GetFaceOptions() *opentype.FaceOptions {
	return &opentype.FaceOptions{
		Size:    factory.Size,
		DPI:     factory.DPI,
		Hinting: font.HintingFull,
	}
}
