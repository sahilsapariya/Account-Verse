package database

import (
	"server/config"
	"server/constants"
	"server/database/providers"
	"server/database/providers/sql"
	"server/logs"
)

var Provider providers.Provider

func InitDB() error {
	var err error
	logger := logs.InitLog("info")

	cfg := config.LoadConfig()

	isSQLite := cfg.DBType == constants.DbTypeSqlite

	if isSQLite {
		logger.Info("Initializing the SQLite database...")
		Provider, err = sql.NewProvider()
		if err != nil {
			logger.Error("Error while initializing the SQLite database: ", err)
			return err
		}
	}

	return nil
}
