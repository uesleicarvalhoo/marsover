package fixture

import "github.com/uesleicarvalhoo/marsrover/rover"

type CoordinatesBuilder struct {
	x int
	y int
}

func AnyCoordinates() CoordinatesBuilder {
	return CoordinatesBuilder{
		x: 1,
		y: 1,
	}
}

func (b CoordinatesBuilder) WithX(x int) CoordinatesBuilder {
	b.x = x
	return b
}

func (b CoordinatesBuilder) WithY(y int) CoordinatesBuilder {
	b.y = y
	return b
}

func (b CoordinatesBuilder) Build() rover.Coordinates {
	return rover.Coordinates{
		X: b.x,
		Y: b.y,
	}
}
