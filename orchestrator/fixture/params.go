package fixture

import (
	"github.com/uesleicarvalhoo/marsrover/orchestrator"
	"github.com/uesleicarvalhoo/marsrover/plateau"
	"github.com/uesleicarvalhoo/marsrover/plateau/fixture"
	"github.com/uesleicarvalhoo/marsrover/rover"
	fixtureRover "github.com/uesleicarvalhoo/marsrover/rover/fixture"
)

type MissionParamsBuilder struct {
	plateauLimits     plateau.Limits
	roverInstructions []orchestrator.RoverInstructions
}

func AnyMissionParams() MissionParamsBuilder {
	return MissionParamsBuilder{
		plateauLimits: fixture.AnyLimits().Build(),
		roverInstructions: []orchestrator.RoverInstructions{
			{
				Params:   fixtureRover.AnyParams().Build(),
				Commands: []rover.Command{rover.MoveForward, rover.TurnLeft, rover.MoveForward},
			},
		},
	}
}

func (b MissionParamsBuilder) WithPlateauLimits(l plateau.Limits) MissionParamsBuilder {
	b.plateauLimits = l
	return b
}

func (b MissionParamsBuilder) AddRoverInstruction(r orchestrator.RoverInstructions) MissionParamsBuilder {
	b.roverInstructions = append(b.roverInstructions, r)
	return b
}

func (b MissionParamsBuilder) WithRoverInstructions(r []orchestrator.RoverInstructions) MissionParamsBuilder {
	b.roverInstructions = r
	return b
}

func (b MissionParamsBuilder) Build() orchestrator.MissionParams {
	return orchestrator.MissionParams{
		PlateauLimits:     b.plateauLimits,
		RoverInstructions: b.roverInstructions,
	}
}
