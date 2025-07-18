package rover

type Command string

const (
	TurnLeft    = "L"
	TurnRight   = "R"
	MoveForward = "M"
)

func (c Command) String() string {
	switch c {
	case TurnLeft:
		return "L"
	case TurnRight:
		return "R"
	case MoveForward:
		return "M"
	}

	return "invalid"
}

func ParseCommand(s string) (Command, error) {
	switch s {
	case "L":
		return TurnLeft, nil
	case "R":
		return TurnRight, nil
	case "M":
		return MoveForward, nil
	}

	return "", &ErrInvalidCommand{
		str: s,
	}
}

func (c Command) IsValid() bool {
	switch c {
	case MoveForward, TurnLeft, TurnRight:
		return true
	default:
		return false
	}
}
