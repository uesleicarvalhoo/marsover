package http

import (
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/uesleicarvalhoo/marsrover/docs" // swagger docs
	"github.com/uesleicarvalhoo/marsrover/internal/config"
	"github.com/uesleicarvalhoo/marsrover/internal/logger"
	"github.com/uesleicarvalhoo/marsrover/orchestrator"
)

func RegisterHandlers(mux *http.ServeMux, svc orchestrator.MissionUseCase) {
}

func NewServer() *http.Server {
	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	port := config.GetInt("HTTP_PORT")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  config.GetDuration("HTTP_SERVER_READ_TIMEOUT"),
		WriteTimeout: config.GetDuration("HTTP_SERVER_WRITE_TIMEOUT"),
	}

	logger.Info("HTTP server listening on http://localhost:%d", port)

	return srv
}
