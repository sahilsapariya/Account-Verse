package database

import (
	log "github.com/sirupsen/logrus"

	"server/constants"
	"server/database/providers"
	"server/database/providers/sql"
	"server/memorystore"
)

// Provider returns the current database provider
var Provider providers.Provider

func InitDB() error {
	var err error

	envs := memorystore.RequiredEnvStoreObj.GetRequiredEnv()

	isSQL := envs.DatabaseType != constants.DbTypeArangodb && envs.DatabaseType != constants.DbTypeMongodb && envs.DatabaseType != constants.DbTypeCassandraDB && envs.DatabaseType != constants.DbTypeScyllaDB && envs.DatabaseType != constants.DbTypeDynamoDB && envs.DatabaseType != constants.DbTypeCouchbaseDB
	
	if isSQL {
		log.Info("Initializing SQL Driver for: ", envs.DatabaseType)
		Provider, err = sql.NewProvider()
		if err != nil {
			log.Fatal("Failed to initialize SQL driver: ", err)
			return err
		}
	}
	

	return nil
}