package internal

import (
	"fmt"

	"github.com/rubensmichel/waner-on-life/internal/env"
	"github.com/rubensmichel/waner-on-life/internal/infra/db"
	"github.com/rubensmichel/waner-on-life/internal/infra/db/postgres"
	"gorm.io/gorm"
)

type Factory struct {
	Env    env.Env
	dtBase *gorm.DB
	DBUser db.Users
}

func NewFactory() (*Factory, error) {
	ft := Factory{}

	appEnv, err := env.Load()
	if err != nil {
		return nil, err
	}
	ft.Env = appEnv

	ft.DBUser = db.NewUserDB(ft.dtBase)

	ft.ConnectDB()

	return &ft, nil
}

func (f *Factory) ConnectDB() {
	var err error
	f.dtBase, err = postgres.ConnectDb()
	if err != nil {
		fmt.Println(err)
	}

}

func (f *Factory) Shutdown() error {
	return nil
}
