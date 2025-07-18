package main

import (
	"net/http"

	"github.com/uesleicarvalhoo/marsrover/internal/config"
	server "github.com/uesleicarvalhoo/marsrover/internal/http"
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

	srv := server.NewServer()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("server error: %v", err)
	}
}
