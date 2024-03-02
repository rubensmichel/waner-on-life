package internal

import (
	"github.com/rubensmichel/waner-on-life/internal/env"
)

type Factory struct {
	Env env.Env
}

func NewFactory() (*Factory, error) {
	ft := Factory{}

	appEnv, err := env.Load()
	if err != nil {
		return nil, err
	}
	ft.Env = appEnv

	return &ft, nil
}

func (f *Factory) Shutdown() error {
	return nil
}
