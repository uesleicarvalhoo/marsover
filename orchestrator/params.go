package orchestrator

import (
	"github.com/uesleicarvalhoo/marsrover/plateau"
	"github.com/uesleicarvalhoo/marsrover/rover"
)

type RoverInstructions struct {
	Params   rover.Params
	Commands []rover.Command
}

type MissionParams struct {
	PlateauLimits     plateau.Limits
	RoverInstructions []RoverInstructions
}
