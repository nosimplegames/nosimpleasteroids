package game

import (
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type Score struct {
	engine.Entity

	Points int
}

func (score *Score) Increase(amount int) {
	score.Points += amount
}

func (score Score) Draw(target *ebiten.Image) {
	render.Text{
		Text:   score.GetScoreString(),
		Target: target,
		Position: math.Vector{
			X: 725,
			Y: 10,
		},
		FontFace: GetAssets().UIFontFace,
	}.Render()
}

func (score Score) GetScoreString() string {
	numberOfCharacters := 10
	scoreString := strconv.Itoa(score.Points)

	needToAppendZeros := len(scoreString) < numberOfCharacters

	if needToAppendZeros {
		requiredZeros := numberOfCharacters - len(scoreString)
		scoreString = strings.Repeat("0", requiredZeros) + scoreString
	}

	return scoreString
}

var score *Score = nil

func GetScore() *Score {
	needToInitScore := score == nil

	if needToInitScore {
		score = &Score{}
	}

	return score
}
