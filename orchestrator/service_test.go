package orchestrator_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/marsrover/orchestrator"
	"github.com/uesleicarvalhoo/marsrover/orchestrator/fixture"
	fixturePlateau "github.com/uesleicarvalhoo/marsrover/plateau/fixture"
	"github.com/uesleicarvalhoo/marsrover/rover"
	fixtureRover "github.com/uesleicarvalhoo/marsrover/rover/fixture"
)

func TestService_Execute(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about       string
		params      orchestrator.MissionParams
		expectErr   string
		expectedRes []rover.Params
	}{
		{
			about: "Execute with single rover successfully",
			params: fixture.AnyMissionParams().
				WithPlateauLimits(
					fixturePlateau.AnyLimits().
						WithX(5).
						WithY(5).
						Build(),
				).
				WithRoverInstructions([]orchestrator.RoverInstructions{
					{
						Params: fixtureRover.AnyParams().
							WithCoordinates(
								fixtureRover.AnyCoordinates().
									WithX(1).
									WithY(2).
									Build()).
							WithDirection(rover.North).
							Build(),
						Commands: []rover.Command{rover.MoveForward, rover.TurnRight, rover.MoveForward},
					},
				}).
				Build(),
			expectedRes: []rover.Params{
				fixtureRover.AnyParams().
					WithCoordinates(fixtureRover.AnyCoordinates().
						WithX(2).
						WithY(3).
						Build()).
					WithDirection(rover.East).
					Build(),
			},
		},
		{
			about: "Error when rover execution fails",
			params: fixture.AnyMissionParams().
				WithPlateauLimits(
					fixturePlateau.AnyLimits().WithX(5).WithY(5).Build(),
				).
				WithRoverInstructions([]orchestrator.RoverInstructions{
					{
						Params: fixtureRover.AnyParams().
							WithCoordinates(fixtureRover.AnyCoordinates().
								WithX(0).
								WithY(0).
								Build()).
							WithDirection(rover.North).
							Build(),
						Commands: []rover.Command{"INVALID"},
					},
				}).
				Build(),
			expectErr: "error while running rover commands: command not configured",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			sut := orchestrator.NewMissionService(2)

			// Act
			res, err := sut.Execute(context.Background(), tc.params)

			// Assert
			if tc.expectErr != "" {
				assert.ErrorContains(t, err, tc.expectErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, len(tc.expectedRes), len(res))
		})
	}
}
