package app

import (
	"fmt"
	"git.garena.com/shopee/MLP/aip/platform/aip-user-service/cmd/app/options"
	"git.garena.com/shopee/MLP/aip/platform/aip-user-service/pkg/router"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type APIServer struct {
	*options.ServerRunOptions
}

func Run() error {
	s := &APIServer{
		ServerRunOptions : options.NewServerRunOptions(),
	}

	// init db

	// set routes
	engine := router.SetRouter()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.Port),
		Handler:      engine,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	return fmt.Errorf("server stopped")
}