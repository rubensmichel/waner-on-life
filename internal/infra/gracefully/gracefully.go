package gracefully

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Shutdowner interface {
	Shutdown() error
}

type gracefully struct {
	funcs []Shutdowner
	trap  chan os.Signal
}

func New() *gracefully {
	grf := gracefully{
		trap: make(chan os.Signal, 1),
	}
	signal.Notify(grf.trap, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	return &grf
}

func (grf *gracefully) Add(f Shutdowner) *gracefully {
	grf.funcs = append(grf.funcs, f)
	return grf
}

func (grf *gracefully) Wait() {
	ctx := context.Background()

	<-grf.trap
	fmt.Print(ctx, "Gracefully shutting down")
	for _, f := range grf.funcs {
		err := f.Shutdown()
		if err != nil {
			fmt.Print(ctx, fmt.Errorf("unable to gracefully shutdown: %w", err))
		}
	}
	fmt.Print(ctx, "Gracefully shutting down")
}
