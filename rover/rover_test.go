package rover_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uesleicarvalhoo/marsrover/plateau"
	fixturePlateau "github.com/uesleicarvalhoo/marsrover/plateau/fixture"
	"github.com/uesleicarvalhoo/marsrover/rover"
	"github.com/uesleicarvalhoo/marsrover/rover/fixture"
)

func ptr[T any](v T) *T {
	return &v
}

func TestNew(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about     string
		plateau   *plateau.Plateau
		params    rover.Params
		expectErr string
	}{
		{
			about: "Create rover with valid params",
			plateau: ptr(fixturePlateau.AnyPlateau().
				WithMaxX(5).
				WithMaxY(5).Build()),
			params: fixture.AnyParams().
				WithCoordinates(
					fixture.AnyCoordinates().
						WithX(2).
						WithY(3).
						Build(),
				).
				WithDirection(rover.North).
				Build(),
		},
		{
			about:     "Fail when plateau is nil",
			plateau:   nil,
			params:    fixture.AnyParams().Build(),
			expectErr: "no plateau provided",
		},
		{
			about: "Fail when coordinates out of bounds",
			plateau: ptr(fixturePlateau.AnyPlateau().
				WithMaxX(2).
				WithMaxY(2).Build()),
			params: fixture.AnyParams().
				WithCoordinates(
					fixture.AnyCoordinates().
						WithX(3).
						WithY(3).
						Build()).
				Build(),
			expectErr: "coordinates out of plateau bounds",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.about, func(t *testing.T) {
			// Arrange

			// Act
			sut, err := rover.New(tc.plateau, tc.params)

			// Assert
			if tc.expectErr != "" {
				assert.Error(t, err)
				assert.Nil(t, sut)
				return
			}

			cor, pos := sut.Position()
			assert.NoError(t, err)
			require.NotNil(t, sut)
			assert.Equal(t, tc.params.Name, sut.Name)
			assert.Equal(t, tc.params.Coordinates, cor)
			assert.Equal(t, tc.params.Direction, pos)
		})
	}
}

func TestExecuteCommands(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about               string
		rover               rover.Rover
		commands            []rover.Command
		expectedDirection   rover.Direction
		expectedCoordinates rover.Coordinates
		expectErr           string
	}{
		{
			about: "Move forward and rotate right",
			rover: fixture.AnyRover().
				WithPlateau(
					fixturePlateau.AnyPlateau().
						WithMaxX(5).
						WithMaxY(5).
						Build(),
				).
				WithCoordinates(
					fixture.AnyCoordinates().
						WithX(1).
						WithY(2).
						Build(),
				).
				WithDirection(rover.North).
				Build(),
			commands: []rover.Command{
				rover.MoveForward, rover.TurnRight, rover.MoveForward,
			},
			expectedCoordinates: fixture.AnyCoordinates().
				WithX(2).
				WithY(3).
				Build(),
			expectedDirection: rover.East,
		},
		{
			about: "Move beyond plateau bounds",
			rover: fixture.AnyRover().
				WithPlateau(
					fixturePlateau.AnyPlateau().
						WithMaxX(2).
						WithMaxY(2).
						Build(),
				).
				WithCoordinates(
					fixture.AnyCoordinates().
						WithX(2).
						WithY(2).
						Build(),
				).
				WithDirection(rover.North).
				Build(),
			commands: []rover.Command{
				rover.MoveForward,
			},
			expectErr: "coordinates out of plateau bounds",
		},
		{
			about: "Invalid command in sequence",
			rover: fixture.AnyRover().
				WithPlateau(
					fixturePlateau.AnyPlateau().
						WithMaxX(5).
						WithMaxY(5).
						Build(),
				).
				WithCoordinates(
					fixture.AnyCoordinates().
						WithX(0).
						WithY(0).
						Build(),
				).
				WithDirection(rover.South).
				Build(),
			commands: []rover.Command{
				"INVALID",
			},
			expectErr: "invalid command",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.about, func(t *testing.T) {
			ctx := context.Background()

			// Arrange

			// Act
			cor, dir, err := tc.rover.ExecuteCommands(ctx, tc.commands)

			// Assert
			if tc.expectErr != "" {
				assert.Error(t, err)
				require.Equal(t, rover.Coordinates{}, cor)
				require.Equal(t, rover.Direction(0), dir)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCoordinates, cor)
			assert.Equal(t, tc.expectedDirection, dir)
		})
	}
}

func TestExecuteCommand(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about               string
		rover               rover.Rover
		command             rover.Command
		expectedCoordinates rover.Coordinates
		expectedDirection   rover.Direction
		expectErr           string
	}{
		{
			about: "Turn left from North to West",
			rover: fixture.AnyRover().
				WithDirection(rover.North).
				Build(),
			command:           rover.TurnLeft,
			expectedDirection: rover.West,
			expectedCoordinates: fixture.AnyCoordinates().
				WithX(1).
				WithY(1).
				Build(),
		},
		{
			about: "Turn right from North to East",
			rover: fixture.AnyRover().
				WithDirection(rover.North).
				Build(),
			command:           rover.TurnRight,
			expectedDirection: rover.East,
			expectedCoordinates: fixture.AnyCoordinates().
				WithX(1).
				WithY(1).
				Build(),
		},
		{
			about: "Move forward within plateau bounds",
			rover: fixture.AnyRover().
				WithPlateau(
					fixturePlateau.AnyPlateau().
						WithMaxX(5).
						WithMaxY(5).
						Build(),
				).
				WithCoordinates(
					fixture.AnyCoordinates().WithX(1).WithY(1).Build(),
				).
				WithDirection(rover.North).
				Build(),
			command: rover.MoveForward,
			expectedCoordinates: fixture.AnyCoordinates().
				WithX(1).
				WithY(2).
				Build(),
			expectedDirection: rover.North,
		},
		{
			about: "Move forward out of plateau bounds",
			rover: fixture.AnyRover().
				WithPlateau(
					fixturePlateau.AnyPlateau().
						WithMaxX(1).
						WithMaxY(1).
						Build(),
				).
				WithCoordinates(
					fixture.AnyCoordinates().WithX(1).WithY(1).Build(),
				).
				WithDirection(rover.North).
				Build(),
			command:   rover.MoveForward,
			expectErr: "coordinates out of plateau bounds",
		},
		{
			about:     "Invalid command",
			rover:     fixture.AnyRover().Build(),
			command:   "INVALID_COMMAND",
			expectErr: "command not configured",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Act
			coordinates, direction, err := tc.rover.ExecuteCommand(context.Background(), tc.command)

			// Assert
			if tc.expectErr != "" {
				assert.Error(t, err)
				assert.Equal(t, rover.Coordinates{}, coordinates)
				assert.Equal(t, rover.Direction(0), direction)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedCoordinates, coordinates)
			assert.Equal(t, tc.expectedDirection, direction)
		})
	}
}
