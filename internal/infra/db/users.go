package db

import (
	"context"

	entity "github.com/rubensmichel/waner-on-life/internal/domain"
	"gorm.io/gorm"
)

type Users interface {
	Find(c context.Context) ([]entity.User, error)
}

type userDB struct {
	database *gorm.DB
}

func NewUserDB(database *gorm.DB) Users {
	return &userDB{
		database: database,
	}
}

func (repo *userDB) Find(ctx context.Context) ([]entity.User, error) {
	var user []entity.User

	response := repo.database.WithContext(ctx).Find(&user)

	if response.Error != nil {
		return nil, response.Error
	}

	return user, nil
}
