package parser_test

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/marsrover/internal/http/parser"
	"github.com/uesleicarvalhoo/marsrover/orchestrator"
	"github.com/uesleicarvalhoo/marsrover/orchestrator/fixture"
	fixturePlateau "github.com/uesleicarvalhoo/marsrover/plateau/fixture"
	"github.com/uesleicarvalhoo/marsrover/rover"
	fixtureRover "github.com/uesleicarvalhoo/marsrover/rover/fixture"
)

func TestParseMission(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about       string
		fileContent string
		expected    orchestrator.MissionParams
		expectErr   string
	}{
		{
			about:       "Valid mission file with 2 rovers",
			fileContent: "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM",
			expected: fixture.AnyMissionParams().
				WithPlateauLimits(
					fixturePlateau.AnyLimits().
						WithX(5).
						WithY(5).
						Build(),
				).
				WithRoverInstructions([]orchestrator.RoverInstructions{
					{
						Params: fixtureRover.AnyParams().
							WithName("Rover-1").
							WithCoordinates(fixtureRover.AnyCoordinates().
								WithX(1).
								WithY(2).
								Build()).
							WithDirection(rover.North).
							Build(),
						Commands: []rover.Command{
							rover.TurnLeft, rover.MoveForward, rover.TurnLeft,
							rover.MoveForward, rover.TurnLeft, rover.MoveForward,
							rover.TurnLeft, rover.MoveForward, rover.MoveForward,
						},
					},
					{
						Params: fixtureRover.AnyParams().
							WithName("Rover-2").
							WithCoordinates(fixtureRover.AnyCoordinates().
								WithX(3).
								WithY(3).
								Build()).
							WithDirection(rover.East).
							Build(),
						Commands: []rover.Command{
							rover.MoveForward, rover.MoveForward, rover.TurnRight,
							rover.MoveForward, rover.MoveForward, rover.TurnRight,
							rover.MoveForward, rover.TurnRight, rover.TurnRight,
							rover.MoveForward,
						},
					},
				}).
				Build(),
		},
		{
			about:       "Empty data returns error",
			fileContent: "",
			expectErr:   "mission data is empty",
		},
		{
			about:       "File with invalid plateau line",
			fileContent: "INVALID\n1 2 N\nLMLMLMLMM",
			expectErr:   "invalid plateau limits: \"INVALID\"",
		},
		{
			about:       "File with missing rover commands",
			fileContent: "5 5\n1 2 N",
			expectErr:   "missing commands for rover 1 at line: 0",
		},
		{
			about:       "File with invalid command",
			fileContent: "5 5\n1 2 N\n\nLMX",
			expectErr:   "invalid rover position at line 3: \"LMX\"",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			ctx := context.Background()
			reader := strings.NewReader(tc.fileContent)

			// Act
			result, err := parser.ParseMission(ctx, reader)

			// Assert
			if tc.expectErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
