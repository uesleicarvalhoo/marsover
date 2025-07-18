package formatter

import (
	"fmt"
	"strings"

	"github.com/uesleicarvalhoo/marsrover/rover"
)

func FormatResultsPlain(finalRovers []rover.Params) string {
	var sb strings.Builder
	for i, r := range finalRovers {
		if i > 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(fmt.Sprintf("%d %d %s", r.Coordinates.X, r.Coordinates.Y, r.Direction.String()))
	}
	return sb.String()
}
