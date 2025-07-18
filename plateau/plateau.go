package plateau

type Plateau struct {
	MinX int `json:"minX"`
	MinY int `json:"minY"`
	MaxX int `json:"maxX"`
	MaxY int `json:"maxY"`
}

func New(limits Limits) (*Plateau, error) {
	if limits.X < 0 || limits.Y < 0 {
		return nil, ErrInvalidPlateauDimensions{
			MinX: 0,
			MinY: 0,
			MaxX: limits.X,
			MaxY: limits.Y,
		}
	}
	return &Plateau{
		MinX: 0, MinY: 0,
		MaxX: limits.X,
		MaxY: limits.Y,
	}, nil
}

func (p *Plateau) ValidateCoordinates(x, y int) error {
	if x < p.MinX || x > p.MaxX || y < p.MinY || y > p.MaxY {
		return ErrCoordinatesOutOfRange{
			MinX:       p.MinX,
			MinY:       p.MinY,
			MaxX:       p.MaxX,
			MaxY:       p.MaxY,
			RequestedX: x,
			RequestedY: y,
		}
	}
	return nil
}
