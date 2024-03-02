package env

import (
	"github.com/caarlos0/env/v8"
)

type Env struct {
	Env string `env:"ENV"`
}

func Load() (Env, error) {
	e := Env{}
	if err := env.Parse(&e); err != nil {
		return e, err
	}
	return e, nil
}
