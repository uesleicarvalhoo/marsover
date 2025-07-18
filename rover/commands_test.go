package rover_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/marsrover/rover"
)

func TestCommand_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about    string
		command  rover.Command
		expected string
	}{
		{
			about:    "MoveForward string",
			command:  rover.MoveForward,
			expected: "M",
		},
		{
			about:    "TurnLeft string",
			command:  rover.TurnLeft,
			expected: "L",
		},
		{
			about:    "TurnRight string",
			command:  rover.TurnRight,
			expected: "R",
		},
		{
			about:    "Invalid command string",
			command:  rover.Command("X"),
			expected: "invalid",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Act
			result := tc.command.String()

			// Assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestParseCommand(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about     string
		input     string
		expected  rover.Command
		expectErr string
	}{
		{
			about:    "Parse M",
			input:    "M",
			expected: rover.MoveForward,
		},
		{
			about:    "Parse L",
			input:    "L",
			expected: rover.TurnLeft,
		},
		{
			about:    "Parse R",
			input:    "R",
			expected: rover.TurnRight,
		},
		{
			about:     "Invalid input",
			input:     "X",
			expectErr: "invalid command",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.about, func(t *testing.T) {
			// Act
			cmd, err := rover.ParseCommand(tc.input)

			// Assert
			if tc.expectErr != "" {
				assert.Error(t, err)
				assert.Equal(t, rover.Command(""), cmd)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, cmd)
		})
	}
}

func TestCommand_IsValid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about    string
		command  rover.Command
		expected bool
	}{
		{
			about:    "Valid MoveForward",
			command:  rover.MoveForward,
			expected: true,
		},
		{
			about:    "Valid TurnLeft",
			command:  rover.TurnLeft,
			expected: true,
		},
		{
			about:    "Valid TurnRight",
			command:  rover.TurnRight,
			expected: true,
		},
		{
			about:    "Invalid command",
			command:  rover.Command("X"),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.about, func(t *testing.T) {
			// Act
			result := tc.command.IsValid()

			// Assert
			assert.Equal(t, tc.expected, result)
		})
	}
}
