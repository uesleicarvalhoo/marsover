package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/marsrover/orchestrator"
	"github.com/uesleicarvalhoo/marsrover/rover"
)

type MissionUseCaseMock struct {
	mock.Mock
}

func (m *MissionUseCaseMock) Execute(ctx context.Context, params orchestrator.MissionParams) ([]rover.Params, error) {
	args := m.Called(ctx, params)

	return args.Get(0).([]rover.Params), args.Error(1)
}
