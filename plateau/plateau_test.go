package plateau_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/marsrover/plateau"
	"github.com/uesleicarvalhoo/marsrover/plateau/fixture"
)

func TestNew(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about     string
		limits    plateau.Limits
		expectErr string
	}{
		{
			about: "Create plateau with valid dimensions",
			limits: fixture.AnyLimits().
				WithX(5).
				WithY(5).
				Build(),
		},
		{
			about: "Fail when X is negative",
			limits: fixture.AnyLimits().
				WithX(-1).
				WithY(5).
				Build(),
			expectErr: "invalid plateau dimensions",
		},
		{
			about: "Fail when Y is negative",
			limits: fixture.AnyLimits().
				WithX(5).
				WithY(-1).
				Build(),
			expectErr: "invalid plateau dimensions",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Act
			p, err := plateau.New(tc.limits)

			// Assert
			if tc.expectErr != "" {
				assert.Error(t, err)
				assert.Nil(t, p)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, p)
			assert.Equal(t, 0, p.MinX)
			assert.Equal(t, 0, p.MinY)
			assert.Equal(t, tc.limits.X, p.MaxX)
			assert.Equal(t, tc.limits.Y, p.MaxY)
		})
	}
}

func TestValidateCoordinates(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		about     string
		plateau   plateau.Plateau
		x, y      int
		expectErr string
	}{
		{
			about: "Coordinates within bounds",
			plateau: fixture.AnyPlateau().
				WithMaxX(5).
				WithMaxY(5).
				Build(),
			x: 3,
			y: 3,
		},
		{
			about: "X coordinate out of bounds (too high)",
			plateau: fixture.AnyPlateau().
				WithMaxX(5).
				WithMaxY(5).
				Build(),
			x:         6,
			y:         3,
			expectErr: "coordinates out of plateau bounds",
		},
		{
			about: "Y coordinate out of bounds (too high)",
			plateau: fixture.AnyPlateau().
				WithMaxX(5).
				WithMaxY(5).
				Build(),
			x:         3,
			y:         6,
			expectErr: "coordinates out of plateau bounds",
		},
		{
			about: "X coordinate out of bounds (too low)",
			plateau: fixture.AnyPlateau().
				WithMaxX(5).
				WithMaxY(5).
				Build(),
			x:         -1,
			y:         3,
			expectErr: "coordinates out of plateau bounds",
		},
		{
			about: "Y coordinate out of bounds (too low)",
			plateau: fixture.AnyPlateau().
				WithMaxX(5).
				WithMaxY(5).
				Build(),
			x:         3,
			y:         -1,
			expectErr: "coordinates out of plateau bounds",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Act
			err := tc.plateau.ValidateCoordinates(tc.x, tc.y)

			// Assert
			if tc.expectErr != "" {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
