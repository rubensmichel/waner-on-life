package env

import (
	"github.com/caarlos0/env/v8"
)

type Env struct {
	Env        string `env:"ENV"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBHost     string `env:"DB_HOST"`
}

func Load() (Env, error) {
	e := Env{}
	if err := env.Parse(&e); err != nil {
		return e, err
	}
	return e, nil
}
