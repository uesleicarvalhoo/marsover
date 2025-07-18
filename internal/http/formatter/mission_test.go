package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/marsrover/internal/http/formatter"
	"github.com/uesleicarvalhoo/marsrover/rover"
	"github.com/uesleicarvalhoo/marsrover/rover/fixture"
)

func TestFormatResultsPlain(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about    string
		input    []rover.Params
		expected string
	}{
		{
			about: "Single rover result",
			input: []rover.Params{
				fixture.AnyParams().
					WithCoordinates(fixture.AnyCoordinates().
						WithX(1).
						WithY(3).
						Build()).
					WithDirection(rover.North).
					Build(),
			},
			expected: "1 3 N",
		},
		{
			about: "Multiple rovers result",
			input: []rover.Params{
				fixture.AnyParams().
					WithCoordinates(fixture.AnyCoordinates().
						WithX(1).
						WithY(3).
						Build()).
					WithDirection(rover.North).
					Build(),
				fixture.AnyParams().
					WithCoordinates(fixture.AnyCoordinates().
						WithX(5).
						WithY(1).
						Build()).
					WithDirection(rover.East).
					Build(),
			},
			expected: "1 3 N\n5 1 E",
		},
		{
			about:    "Empty result list",
			input:    []rover.Params{},
			expected: "",
		},
		{
			about: "All directions covered",
			input: []rover.Params{
				fixture.AnyParams().
					WithCoordinates(fixture.AnyCoordinates().
						WithX(0).
						WithY(0).
						Build()).
					WithDirection(rover.North).
					Build(),
				fixture.AnyParams().
					WithCoordinates(fixture.AnyCoordinates().
						WithX(1).
						WithY(1).
						Build()).
					WithDirection(rover.East).
					Build(),
				fixture.AnyParams().
					WithCoordinates(fixture.AnyCoordinates().
						WithX(2).
						WithY(2).
						Build()).
					WithDirection(rover.South).
					Build(),
				fixture.AnyParams().
					WithCoordinates(fixture.AnyCoordinates().
						WithX(3).
						WithY(3).
						Build()).
					WithDirection(rover.West).
					Build(),
			},
			expected: "0 0 N\n1 1 E\n2 2 S\n3 3 W",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Act
			output := formatter.FormatResultsPlain(tc.input)

			// Assert
			assert.Equal(t, tc.expected, output)
		})
	}
}
