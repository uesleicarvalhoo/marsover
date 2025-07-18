package fixture

import "github.com/uesleicarvalhoo/marsrover/rover"

type ParamsBuilder struct {
	name        string
	coordinates rover.Coordinates
	direction   rover.Direction
}

func AnyParams() ParamsBuilder {
	return ParamsBuilder{
		name:        "Rover1",
		coordinates: AnyCoordinates().Build(),
		direction:   rover.North,
	}
}

func (b ParamsBuilder) WithName(name string) ParamsBuilder {
	b.name = name
	return b
}

func (b ParamsBuilder) WithCoordinates(c rover.Coordinates) ParamsBuilder {
	b.coordinates = c
	return b
}

func (b ParamsBuilder) WithDirection(d rover.Direction) ParamsBuilder {
	b.direction = d
	return b
}

func (b ParamsBuilder) Build() rover.Params {
	return rover.Params{
		Name:        b.name,
		Coordinates: b.coordinates,
		Direction:   b.direction,
	}
}
