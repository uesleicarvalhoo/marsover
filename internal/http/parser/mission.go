package parser

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/uesleicarvalhoo/marsrover/orchestrator"
	"github.com/uesleicarvalhoo/marsrover/plateau"
	"github.com/uesleicarvalhoo/marsrover/rover"
)

func ParseMission(ctx context.Context, reader io.Reader) (orchestrator.MissionParams, error) {
	scanner := bufio.NewScanner(reader)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return orchestrator.MissionParams{}, err
		}

		return orchestrator.MissionParams{}, &ErrEmptyData{}
	}

	limits, err := parsePlateauLimits(scanner.Text())
	if err != nil {
		return orchestrator.MissionParams{}, err
	}

	line := 1
	roverIdx := 0
	var instructions []orchestrator.RoverInstructions

	for scanner.Scan() {
		line++
		value := strings.TrimSpace(scanner.Text())
		if value == "" {
			continue
		}
		roverIdx++

		cor, dir, err := parsePosition(line, value)
		if err != nil {
			return orchestrator.MissionParams{}, err
		}

		if !scanner.Scan() {
			return orchestrator.MissionParams{}, ErrMissingCommands{Index: roverIdx}
		}
		cmds, err := parseCommands(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return orchestrator.MissionParams{}, err
		}

		instructions = append(instructions, orchestrator.RoverInstructions{
			Params: rover.Params{
				Name:        fmt.Sprintf("Rover-%d", roverIdx),
				Coordinates: cor,
				Direction:   dir,
			},
			Commands: cmds,
		})
	}

	if err := scanner.Err(); err != nil {
		return orchestrator.MissionParams{}, err
	}

	return orchestrator.MissionParams{
		PlateauLimits:     limits,
		RoverInstructions: instructions,
	}, nil
}

func parsePlateauLimits(value string) (plateau.Limits, error) {
	parts := strings.Fields(value)
	if len(parts) != 2 {
		return plateau.Limits{}, ErrInvalidPlateauLine{Value: value}
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return plateau.Limits{}, ErrInvalidPlateauXValue{Value: parts[0], Err: err}
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return plateau.Limits{}, ErrInvalidPlateauYValue{Value: parts[1], Err: err}
	}
	return plateau.Limits{X: x, Y: y}, nil
}

func parsePosition(line int, value string) (rover.Coordinates, rover.Direction, error) {
	parts := strings.Fields(value)
	if len(parts) != 3 {
		return rover.Coordinates{}, 0, ErrInvalidPositionLine{Value: value, Line: line}
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return rover.Coordinates{}, 0, ErrInvalidPositionX{Value: parts[0], Err: err, Line: line}
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return rover.Coordinates{}, 0, ErrInvalidPositionY{Value: parts[1], Err: err, Line: line}
	}

	dirStr := parts[2]
	if len(dirStr) != 1 {
		return rover.Coordinates{}, 0, ErrInvalidDirectionFormat{Value: dirStr, Line: line}
	}
	dir, err := rover.ParseDirection(dirStr)
	if err != nil {
		return rover.Coordinates{}, 0, err
	}

	return rover.Coordinates{X: x, Y: y}, dir, nil
}

func parseCommands(value string) ([]rover.Command, error) {
	var cmds []rover.Command
	for _, r := range value {
		cmd, err := rover.ParseCommand(string(r))
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmd)
	}
	return cmds, nil
}
