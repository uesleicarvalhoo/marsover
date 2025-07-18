package main

import (
	"fmt"
	"net/http"

	"github.com/uesleicarvalhoo/marsrover/internal/config"
	transport "github.com/uesleicarvalhoo/marsrover/internal/http"
	"github.com/uesleicarvalhoo/marsrover/internal/ioc"
	"github.com/uesleicarvalhoo/marsrover/internal/logger"
)

// swagger doc info
// @title Mars Rover API
// @version 1.0
// @description API para controle de rovers em um platô
func main() {
	logger.Configure(logger.Params{
		Level:          config.GetString("LOG_LEVEL"),
		ServiceName:    config.GetString("SERVICE_NAME"),
		ServiceVersion: config.GetString("SERVICE_VERSION"),
		Env:            config.GetString("ENV"),
	})

	missionSvc := ioc.OrchestratorMissionService()

	mux := http.NewServeMux()
	transport.RegisterHandlers(mux, missionSvc)

	port := config.GetInt("HTTP_PORT")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  config.GetDuration("HTTP_SERVER_READ_TIMEOUT"),
		WriteTimeout: config.GetDuration("HTTP_SERVER_WRITE_TIMEOUT"),
	}

	logger.Info("HTTP server listening on http://localhost:%d", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("server error: %v", err)
	}
}
