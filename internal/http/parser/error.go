package parser

import "fmt"

type ErrEmptyData struct{}

func (e ErrEmptyData) Error() string {
	return "mission data is empty"
}

type ErrInvalidPlateauLine struct {
	Value string
}

func (e ErrInvalidPlateauLine) Error() string {
	return fmt.Sprintf("invalid plateau limits: %q", e.Value)
}

type ErrInvalidPlateauXValue struct {
	Value string
	Err   error
}

func (e ErrInvalidPlateauXValue) Error() string {
	return fmt.Sprintf("invalid plateau X value %q: %v", e.Value, e.Err)
}

type ErrInvalidPlateauYValue struct {
	Value string
	Err   error
}

func (e ErrInvalidPlateauYValue) Error() string {
	return fmt.Sprintf("invalid plateau Y value %q: %v", e.Value, e.Err)
}

type ErrInvalidPositionLine struct {
	Line  int
	Value string
}

func (e ErrInvalidPositionLine) Error() string {
	return fmt.Sprintf("invalid rover position at line %d: %q", e.Line, e.Value)
}

type ErrInvalidPositionX struct {
	Line  int
	Value string
	Err   error
}

func (e ErrInvalidPositionX) Error() string {
	return fmt.Sprintf("invalid rover X position at line %d - %q: %v", e.Line, e.Value, e.Err)
}

type ErrInvalidPositionY struct {
	Line  int
	Value string
	Err   error
}

func (e ErrInvalidPositionY) Error() string {
	return fmt.Sprintf("invalid rover Y position %q at line %d: %v", e.Value, e.Line, e.Err)
}

type ErrInvalidDirectionFormat struct {
	Line  int
	Value string
}

func (e ErrInvalidDirectionFormat) Error() string {
	return fmt.Sprintf("invalid direction format at line %d: %q", e.Line, e.Value)
}

type ErrMissingCommands struct {
	Index int
	Line  int
}

func (e ErrMissingCommands) Error() string {
	return fmt.Sprintf("missing commands for rover %d at line: %d", e.Index, e.Line)
}
