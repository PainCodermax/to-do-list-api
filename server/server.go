package server

import (
	"context"
	"net/http"

	"github.com/PainCodermax/to-do-list-api/internal/middleware"
	"github.com/PainCodermax/to-do-list-api/pkg/setting"
	"github.com/PainCodermax/to-do-list-api/server/api"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Router *gin.Engine
	Svr    *http.Server
	Config *setting.Configuration
}

func NewServer(cfg *setting.Configuration) *Server {
	r := gin.New()

	middleware.UseDefault(r, cfg)

	api.SetRouters(r, cfg)

	srv := &http.Server{
		Addr:           ":" + cfg.HTTPPort,
		Handler:        r,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		Config: cfg,
		Svr:    srv,
		Router: r,
	}
}

func (s *Server) Start() error {
	go func() {
		log.Info("Starting HTTP Server at :", s.Config.HTTPPort)
		if err := s.Svr.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("HTTP server expcetpion. ", err)
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Svr.Shutdown(ctx)
}
