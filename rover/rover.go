package rover

import (
	"context"

	"github.com/uesleicarvalhoo/marsrover/plateau"
)

type Rover struct {
	Name        string
	plateau     *plateau.Plateau
	coordinates Coordinates
	direction   Direction
}

func New(plat *plateau.Plateau, params Params) (*Rover, error) {
	if plat == nil {
		return nil, ErrNoPlateuProvided
	}

	err := plat.ValidateCoordinates(params.Coordinates.X, params.Coordinates.Y)
	if err != nil {
		return nil, err
	}

	r := &Rover{
		Name:        params.Name,
		coordinates: params.Coordinates,
		plateau:     plat,
		direction:   params.Direction,
	}

	return r, nil
}

func (r *Rover) ExecuteCommands(ctx context.Context, commands []Command) (Coordinates, Direction, error) {
	for _, cmd := range commands {
		if _, _, err := r.ExecuteCommand(ctx, cmd); err != nil {
			return Coordinates{}, 0, err
		}
	}

	return r.coordinates, r.direction, nil
}

func (r *Rover) ExecuteCommand(ctx context.Context, cmd Command) (Coordinates, Direction, error) {
	switch cmd {
	case TurnLeft:
		r.turnLeft()
	case TurnRight:
		r.turnRight()
	case MoveForward:
		if err := r.move(); err != nil {
			return Coordinates{}, 0, err
		}

	default:
		return Coordinates{}, 0, ErrCommandNotConfigured{
			cmd: cmd,
		}
	}

	return r.coordinates, r.direction, nil
}

func (r *Rover) Position() (Coordinates, Direction) {
	return r.coordinates, r.direction
}

func (r *Rover) turnLeft() {
	r.direction = (r.direction + 3) % 4
}

func (r *Rover) turnRight() {
	r.direction = (r.direction + 1) % 4
}

func (r *Rover) move() error {
	x, y := r.coordinates.X, r.coordinates.Y
	switch r.direction {
	case North:
		y++
	case East:
		x++
	case South:
		y--
	case West:
		x--
	default:
		return ErrInvalidDirection{
			str: r.direction.String(),
		}
	}

	if err := r.plateau.ValidateCoordinates(x, y); err != nil {
		return err
	}

	r.coordinates.X, r.coordinates.Y = x, y

	return nil
}
