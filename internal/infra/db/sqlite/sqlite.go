package sqlite

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type InMemoryDB struct {
	*gorm.DB
	dsn string
}

func NewInMemoryDatabase(tables []interface{}) (*InMemoryDB, error) {
	database := build()
	err := database.connect()
	if err != nil {
		return nil, err
	}
	err = database.migrate(tables)
	if err != nil {
		return nil, err
	}
	return database, nil
}

func build() *InMemoryDB {
	return &InMemoryDB{dsn: "file::memory:?cache=shared"}
}

func (db *InMemoryDB) connect() error {
	newLogger := gormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormlogger.Config{
			SlowThreshold:             time.Nanosecond,
			LogLevel:                  gormlogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	database, err := gorm.Open(sqlite.Open(db.dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	db.DB = database
	return nil
}

func (db *InMemoryDB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	sqlDB.Close()
	return nil
}

func (db *InMemoryDB) migrate(tables []interface{}) error {
	for _, t := range tables {
		if db.Migrator().HasTable(&t) {
			continue
		}
		err := db.Migrator().CreateTable(&t)
		if err != nil {
			return err
		}
	}

	db.Reset(tables)

	return nil
}

func (db *InMemoryDB) Reset(tables []interface{}) error {
	for _, t := range tables {
		err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(t).Error
		if err != nil {
			return err
		}
	}
	return nil
}
