package rover_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/marsrover/rover"
)

func TestDirection_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about    string
		dir      rover.Direction
		expected string
	}{
		{
			about:    "North",
			dir:      rover.North,
			expected: "N",
		},
		{
			about:    "East",
			dir:      rover.East,
			expected: "E",
		},
		{
			about:    "South",
			dir:      rover.South,
			expected: "S",
		},
		{
			about:    "West",
			dir:      rover.West,
			expected: "W",
		},
		{
			about:    "Invalid direction",
			dir:      rover.Direction(99),
			expected: "invalid",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Act
			result := tc.dir.String()

			// Assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestParseDirection(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about     string
		input     string
		expected  rover.Direction
		expectErr string
	}{
		{
			about:    "Parse N",
			input:    "N",
			expected: rover.North,
		},
		{
			about:    "Parse E",
			input:    "E",
			expected: rover.East,
		},
		{
			about:    "Parse S",
			input:    "S",
			expected: rover.South,
		},
		{
			about:    "Parse W",
			input:    "W",
			expected: rover.West,
		},
		{
			about:     "Invalid input",
			input:     "X",
			expectErr: "invalid direction",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.about, func(t *testing.T) {
			// Act
			dir, err := rover.ParseDirection(tc.input)

			// Assert
			if tc.expectErr != "" {
				assert.Error(t, err)
				assert.Equal(t, rover.Direction(0), dir)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, dir)
		})
	}
}

func TestDirection_IsValid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about    string
		dir      rover.Direction
		expected bool
	}{
		{
			about:    "Valid North",
			dir:      rover.North,
			expected: true,
		},
		{
			about:    "Valid East",
			dir:      rover.East,
			expected: true,
		},
		{
			about:    "Valid South",
			dir:      rover.South,
			expected: true,
		},
		{
			about:    "Valid West",
			dir:      rover.West,
			expected: true,
		},
		{
			about:    "Invalid direction",
			dir:      rover.Direction(99),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.about, func(t *testing.T) {
			// Act
			isValid := tc.dir.IsValid()

			// Assert
			assert.Equal(t, tc.expected, isValid)
		})
	}
}
