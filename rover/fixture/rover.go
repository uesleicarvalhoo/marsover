package fixture

import (
	"log"

	"github.com/uesleicarvalhoo/marsrover/plateau"
	fixturePlateu "github.com/uesleicarvalhoo/marsrover/plateau/fixture"
	"github.com/uesleicarvalhoo/marsrover/rover"
)

type RoverBuilder struct {
	name        string
	plateau     plateau.Plateau
	coordinates rover.Coordinates
	direction   rover.Direction
}

func AnyRover() RoverBuilder {
	return RoverBuilder{
		name:    "Rover-1",
		plateau: fixturePlateu.AnyPlateau().Build(),
		coordinates: AnyCoordinates().
			Build(),
		direction: rover.North,
	}
}

func (b RoverBuilder) WithName(name string) RoverBuilder {
	b.name = name
	return b
}

func (b RoverBuilder) WithPlateau(p plateau.Plateau) RoverBuilder {
	b.plateau = p
	return b
}

func (b RoverBuilder) WithCoordinates(c rover.Coordinates) RoverBuilder {
	b.coordinates = c
	return b
}

func (b RoverBuilder) WithDirection(d rover.Direction) RoverBuilder {
	b.direction = d
	return b
}

func (b RoverBuilder) Build() rover.Rover {
	r, err := rover.New(&b.plateau, rover.Params{
		Name:        b.name,
		Coordinates: b.coordinates,
		Direction:   b.direction,
	})
	if err != nil {
		log.Fatalf("failed to build rover fixture: %v", err)
	}

	return *r
}

func (b RoverBuilder) BuildWithError() (rover.Rover, error) {
	r, err := rover.New(&b.plateau, rover.Params{
		Name:        b.name,
		Coordinates: b.coordinates,
		Direction:   b.direction,
	})
	if err != nil {
		return rover.Rover{}, err
	}

	return *r, nil
}
