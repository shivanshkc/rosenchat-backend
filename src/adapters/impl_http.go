package adapters

import (
	"context"
	"net/http"
	"rosenchat/src/router"
	"time"
)

// implHTTP implements IAdapter as an HTTP server.
type implHTTP struct {
	server  *http.Server
	handler http.Handler
}

func (i *implHTTP) Name() string {
	return "HTTP"
}

func (i *implHTTP) Start(ctx context.Context) error {
	if !conf.HTTPServer.Enabled {
		log.Infof(ctx, "HTTP server is not enabled.")
		return nil
	}

	i.server = &http.Server{
		Addr:    conf.HTTPServer.Addr,
		Handler: i.handler,
	}

	log.Infof(ctx, "Starting %s@%s HTTP server at %s", conf.Application.Name, conf.Application.Version, conf.HTTPServer.Addr)
	return i.server.ListenAndServe()
}

func (i *implHTTP) Stop(ctx context.Context) error {
	if i.server == nil {
		log.Infof(ctx, "Not doing anything because server pointer is nil.")
		return nil
	}

	log.Infof(ctx, "Stopping the HTTP server...")
	shutdownTimeoutDuration := time.Second * time.Duration(conf.HTTPServer.ShutDownTimeoutSec)
	shutdownCtx, cancelFunc := context.WithTimeout(context.Background(), shutdownTimeoutDuration)
	defer cancelFunc()

	errShutdown := i.server.Shutdown(shutdownCtx)
	if errShutdown == nil {
		log.Infof(ctx, "HTTP server closed gracefully.")
		return nil
	}

	log.Errorf(ctx, "Failed to close HTTP server gracefully because: %+v. Attempting a forced close...", errShutdown)
	if err := i.server.Close(); err != nil {
		log.Errorf(ctx, "Failed to force close HTTP server because: %+v", err)
		return err
	}

	log.Warnf(ctx, "HTTP server closed forcefully.")
	return nil
}

func (i *implHTTP) init() {
	i.handler = router.Get()
}
