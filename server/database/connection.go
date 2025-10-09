package database

import (
	"context"
	"log"
	"server/config"
	"server/database/mongodb"
	"server/database/sql"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Database struct {
	Type  string
	SQL   *gorm.DB
	Mongo *mongo.Database
}

func NewDatabase(cfg *config.Config) *Database {
	db := &Database{
		Type: cfg.DBType,
	}

	switch cfg.DBType {
	case "sqlite", "postgres", "mysql":
		sqlDB, err := sql.NewSQLConnection(cfg)
		if err != nil {
			log.Fatalf("Failed to connect to SQL database: %v", err)
		}
		db.SQL = sqlDB
	case "mongodb", "mongo":
		mongoDB, err := mongodb.NewMongoConnection(cfg)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}
		db.Mongo = mongoDB
	default:
		log.Fatalf("Unsupported database type: %s", cfg.DBType)
	}

	return db
}

func (db *Database) Close() error {
	if db.SQL != nil {
		sqlDB, err := db.SQL.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}

	if db.Mongo != nil {
		return db.Mongo.Client().Disconnect(context.Background())
	}

	return nil
}
