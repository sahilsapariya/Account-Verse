package sql

import (
	"fmt"
	"server/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLConnection(cfg *config.Config) (*gorm.DB, error) {
	var dsn string

	switch cfg.DBType {
	case "sqlite":
		dsn = cfg.DBName
		return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
		// Note: Add postgres driver import when needed
		// return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
		// Note: Add mysql driver import when needed
		// return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	return nil, fmt.Errorf("unsupported SQL database type: %s", cfg.DBType)
}
