package sql

import (
	"server/config"
	"server/graph/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type provider struct {
	db *gorm.DB
}

func NewProvider() (*provider, error) {
	var sqliteDB *gorm.DB
	var err error

	DBName := config.LoadConfig().DBName

	sqliteDB, _ = gorm.Open(sqlite.Open(DBName), &gorm.Config{})

	err = sqliteDB.AutoMigrate(&model.User{})

	if err != nil {
		return nil, err
	}

	return &provider{
		db: sqliteDB,
	}, nil
}
