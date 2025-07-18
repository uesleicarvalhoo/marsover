package ioc

import (
	"sync"

	"github.com/uesleicarvalhoo/marsrover/internal/config"
	"github.com/uesleicarvalhoo/marsrover/orchestrator"
)

var (
	missionSvcOnce sync.Once
	missionSvc     orchestrator.MissionUseCase
)

func OrchestratorMissionService() orchestrator.MissionUseCase {
	missionSvcOnce.Do(func() {
		missionSvc = orchestrator.NewMissionService(
			config.GetInt("ORCHESTRATOR_CONCURRENCY"),
		)
	})

	return missionSvc
}
