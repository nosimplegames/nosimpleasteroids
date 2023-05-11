package ui

import "simple-games.com/asteroids/utils"

type JustifyContentPolicy = string

const (
	Start        JustifyContentPolicy = ""
	Spacebetween                      = "Spacebetween"
)

type LinearPositionCalculator struct {
	Size                 float64
	JustifyContentPolicy JustifyContentPolicy
	ElementsSizes        []float64
	Gap                  float64
}

func (calculator LinearPositionCalculator) Calculate() []float64 {
	positions := []float64{}
	spaceBetweenElements := calculator.CalculateSpaceBetweenElements()
	offsetPosition := -calculator.Size * 0.5

	for _, elementSize := range calculator.ElementsSizes {
		elementPosition := offsetPosition + elementSize*0.5
		positions = append(positions, elementPosition)
		offsetPosition += elementSize + spaceBetweenElements
	}

	return positions
}

func (calculator LinearPositionCalculator) CalculateElementsSize() float64 {
	return utils.Reduce(
		calculator.ElementsSizes,
		func(accumulator float64, elementSize float64) float64 {
			return accumulator + elementSize
		},
		0.0,
	)
}

func (calculator LinearPositionCalculator) CalculateFreeSpace() float64 {
	elementsSize := calculator.CalculateElementsSize()
	freeSpace := calculator.Size - elementsSize

	return freeSpace
}

func (calculator LinearPositionCalculator) CalculateSpaceBetweenElements() float64 {
	mayUtilizeGap := calculator.JustifyContentPolicy != Spacebetween

	if mayUtilizeGap {
		return calculator.Gap
	}

	freeSpace := calculator.CalculateFreeSpace()
	spaceBetweenElements := freeSpace / float64(len(calculator.ElementsSizes)-1)

	return spaceBetweenElements
}
