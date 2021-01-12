package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aaorlov/stream/log"
	"github.com/aaorlov/stream/server"
	"github.com/aaorlov/stream/version"
)

const defaultShutDownWaitTime time.Duration = time.Duration(5) * time.Second

// Application encapsulates a server application.
type Application struct {
	logger           log.Logger
	waitStopCh       chan os.Signal
	shutDownWaitSecs time.Duration
	srv              server.Srv
}

// New returns a runnable application given an output and a command line arguments array.
func New() *Application {
	return &Application{
		waitStopCh:       make(chan os.Signal, 1),
		shutDownWaitSecs: defaultShutDownWaitTime}
}

// Run runs did-steam application until either a stop signal is received or an error occurs.
func (a *Application) Run() error {
	// initialize logger
	if err := a.initLogger(); err != nil {
		return err
	}

	a.showVersion()

	a.srv = server.New()
	a.srv.Start()

	// ...wait for stop signal to shutdown
	sig := a.waitForStopSignal()
	log.Infof("received %s signal... shutting down...", sig.String())
	return a.gracefullyShutdown()
}

func (a *Application) showVersion() {
	log.Infof("stream version: %v\n", version.ApplicationVersion)
}

func (a *Application) initLogger() error {
	l, err := log.New("info")
	if err != nil {
		return err
	}
	a.logger = l
	log.Set(a.logger)
	return nil
}

func (a *Application) waitForStopSignal() os.Signal {
	signal.Notify(a.waitStopCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	return <-a.waitStopCh
}

func (a *Application) gracefullyShutdown() error {
	// wait until application has been shut down
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(a.shutDownWaitSecs))
	defer cancel()

	select {
	case <-a.shutdown(ctx):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (a *Application) shutdown(ctx context.Context) <-chan bool {
	c := make(chan bool, 1)
	go func() {
		if err := a.doShutdown(ctx); err != nil {
			log.Warnf("failed to shutdown: %s", err)
		}

		c <- true
	}()
	return c
}

func (a *Application) doShutdown(ctx context.Context) error {
	if err := a.srv.Shutdown(ctx); err != nil {
		return err
	}

	log.Unset()
	return nil
}
