package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"rosenchat/src/adapters"
	"syscall"
)

// application represents the entire backend application.
type application struct {
	// adapters are the various adapters/functionalities that the application has.
	adapters []adapters.IAdapter
}

// start method launches the application by starting all the adapters.
// It panics if any of the adapters fails to start.
//
// It also sets up a listener for any SIGTERM interruptions to gracefully
// stop the application.
func (a *application) start() {
	for _, adapter := range a.adapters {
		go func(adapter adapters.IAdapter) {
			if err := adapter.Start(context.Background()); err != nil {
				panic(fmt.Errorf("error from %s adapter: %w", adapter.Name(), err))
			}
		}(adapter)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt
		a.stop()
		close(interrupt)
	}()
}

// stop method stops the application by stopping all the adapters.
func (a *application) stop() {
	for _, adapter := range a.adapters {
		go func(adapter adapters.IAdapter) {
			if err := adapter.Stop(context.Background()); err != nil {
				fmt.Printf("adapter: %s failed to stop: %+v\n", adapter.Name(), err)
			} else {
				fmt.Printf("adapter: %s gracefully stopped.\n", adapter.Name())
			}
		}(adapter)
	}
}

func main() {
	app := &application{}
	app.adapters = append(
		app.adapters,
		adapters.GetHTTP(),
		adapters.GetGRPC(),
		adapters.GetCleanup(),
	)

	app.start()
	select {}
}
