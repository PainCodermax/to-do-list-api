package signals

import (
	"context"
	"time"

	"github.com/PainCodermax/to-do-list-api/server"
	log "github.com/sirupsen/logrus"
)

type Shutdown struct {
	serverShutdownTimeout time.Duration
}

func NewShutdown(serverShutdownTimeout time.Duration) (*Shutdown, error) {
	srv := &Shutdown{
		serverShutdownTimeout: serverShutdownTimeout,
	}

	return srv, nil
}

func (s *Shutdown) Graceful(stopCh <-chan struct{}, svr *server.Server) {
	ctx := context.Background()
	<-stopCh
	ctx, cancel := context.WithTimeout(ctx, s.serverShutdownTimeout)
	defer cancel()

	log.Info("Shutting down HTTP/HTTPS server. ", s.serverShutdownTimeout)
	if err := svr.Shutdown(ctx); err != nil {
		log.Warn("HTTP server graceful shutdown failed", err)
	}

	log.Error("Shutdown complete.")
}
