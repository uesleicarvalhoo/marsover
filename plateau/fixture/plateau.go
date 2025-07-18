package fixture

import "github.com/uesleicarvalhoo/marsrover/plateau"

type PlateauBuilder struct {
	minX int
	minY int
	maxX int
	maxY int
}

func AnyPlateau() PlateauBuilder {
	return PlateauBuilder{
		minX: 0,
		minY: 0,
		maxX: 5,
		maxY: 5,
	}
}

func (b PlateauBuilder) WithMinX(minX int) PlateauBuilder {
	b.minX = minX
	return b
}

func (b PlateauBuilder) WithMinY(minY int) PlateauBuilder {
	b.minY = minY
	return b
}

func (b PlateauBuilder) WithMaxX(maxX int) PlateauBuilder {
	b.maxX = maxX
	return b
}

func (b PlateauBuilder) WithMaxY(maxY int) PlateauBuilder {
	b.maxY = maxY
	return b
}

func (b PlateauBuilder) Build() plateau.Plateau {
	return plateau.Plateau{
		MinX: b.minX,
		MinY: b.minY,
		MaxX: b.maxX,
		MaxY: b.maxY,
	}
}

func (b PlateauBuilder) BuildWithError() (*plateau.Plateau, error) {
	return plateau.New(plateau.Limits{
		X: b.maxX,
		Y: b.maxY,
	})
}
