package rover

type Direction int

const (
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)

func (d Direction) String() string {
	switch d {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	}

	return "invalid"
}

func (d Direction) IsValid() bool {
	switch d {
	case North, East, South, West:
		return true
	default:
		return false
	}
}

func ParseDirection(s string) (Direction, error) {
	switch s {
	case "N":
		return North, nil
	case "E":
		return East, nil
	case "S":
		return South, nil
	case "W":
		return West, nil
	}
	return 0, &ErrInvalidDirection{
		str: s,
	}
}
