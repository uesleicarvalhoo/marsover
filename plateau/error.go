package plateau

import "fmt"

type ErrCoordinatesOutOfRange struct {
	MaxX       int
	MaxY       int
	MinX       int
	MinY       int
	RequestedX int
	RequestedY int
}

func (e ErrCoordinatesOutOfRange) Error() string {
	return fmt.Sprintf(
		"coordinates out of range: requested (%d, %d); allowed X in [%d–%d], Y in [%d–%d]",
		e.RequestedX, e.RequestedY,
		e.MinX, e.MaxX,
		e.MinY, e.MaxY,
	)
}

type ErrInvalidPlateauDimensions struct {
	MinX, MinY int
	MaxX, MaxY int
}

func (e ErrInvalidPlateauDimensions) Error() string {
	return fmt.Sprintf(
		"invalid plateau dimensions: expected X≥%d and Y≥%d, got X=%d, Y=%d",
		e.MinX, e.MinY, e.MaxX, e.MaxY,
	)
}
