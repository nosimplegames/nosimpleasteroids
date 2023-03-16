package assets

import (
	"fmt"
	"os"

	"golang.org/x/image/font/sfnt"
)

type FontManager struct {
	Fonts map[string]*sfnt.Font
}

func (fontManager *FontManager) GetFont(fontFileName string) *sfnt.Font {
	font := fontManager.Fonts[fontFileName]
	fontAlreadyLoaded := font != nil

	if !fontAlreadyLoaded {
		font = fontManager.LoadFont(fontFileName)
		fontManager.Fonts[fontFileName] = font
	}

	return font
}

func (fontManager FontManager) LoadFont(fontFileName string) *sfnt.Font {
	fontBytes := fontManager.GetFontBytes(fontFileName)
	font, err := sfnt.Parse(fontBytes)

	if err != nil {
		fmt.Println("Error getting font")
		panic(err)
	}

	return font
}

func (fontManager FontManager) GetFontBytes(fontFileName string) []byte {
	bytes, err := os.ReadFile(fontFileName)

	if err != nil {
		fmt.Println("Error getting font bytes")
		panic(err)
	}

	return bytes
}

var fontManager *FontManager = nil

func GetFontManager() *FontManager {
	needToInitFontManager := fontManager == nil

	if needToInitFontManager {
		fontManager = &FontManager{
			Fonts: map[string]*sfnt.Font{},
		}
	}

	return fontManager
}
