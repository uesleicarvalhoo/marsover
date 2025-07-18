package fixture

import "github.com/uesleicarvalhoo/marsrover/plateau"

type LimitsBuilder struct {
	x int
	y int
}

func AnyLimits() LimitsBuilder {
	return LimitsBuilder{
		x: 5,
		y: 5,
	}
}

func (b LimitsBuilder) WithX(x int) LimitsBuilder {
	b.x = x
	return b
}

func (b LimitsBuilder) WithY(y int) LimitsBuilder {
	b.y = y
	return b
}

func (b LimitsBuilder) Build() plateau.Limits {
	return plateau.Limits{
		X: b.x,
		Y: b.y,
	}
}
