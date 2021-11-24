package adapters

import (
	"context"
	"net/http"
	"rosenchat/src/configs"
	"rosenchat/src/logger"
	"rosenchat/src/router"
	"sync"
	"time"
)

var httpOnce = &sync.Once{}
var httpSingleton IAdapter

var conf = configs.Get()
var log = logger.Get()

// GetHTTP provides the HTTP adapter singleton.
func GetHTTP() IAdapter {
	httpOnce.Do(func() {
		httpSingleton = &implHTTP{}
		httpSingleton.init()
	})

	return httpSingleton
}

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
		log.Infof("HTTP server is not enabled.")
		return nil
	}

	i.server = &http.Server{
		Addr:    conf.HTTPServer.Addr,
		Handler: i.handler,
	}

	log.Infof("Starting %s@%s HTTP server at %s", conf.Application.Name, conf.Application.Version, conf.HTTPServer.Addr)
	return i.server.ListenAndServe()
}

func (i *implHTTP) Stop(ctx context.Context) error {
	if i.server == nil {
		log.Infof("Not doing anything because server pointer is nil.")
		return nil
	}

	log.Infof("Stopping the HTTP server...")
	shutdownTimeoutDuration := time.Second * time.Duration(conf.HTTPServer.ShutDownTimeoutSec)
	shutdownCtx, cancelFunc := context.WithTimeout(context.Background(), shutdownTimeoutDuration)
	defer cancelFunc()

	errShutdown := i.server.Shutdown(shutdownCtx)
	if errShutdown == nil {
		log.Infof("HTTP server closed gracefully.")
		return nil
	}

	log.Errorf("Failed to close HTTP server gracefully because: %+v. Attempting a forced close...", errShutdown)
	if err := i.server.Close(); err != nil {
		log.Errorf("Failed to force close HTTP server because: %+v", err)
		return err
	}

	log.Warnf("HTTP server closed forcefully.")
	return nil
}

func (i *implHTTP) init() {
	i.handler = router.Get()
}
