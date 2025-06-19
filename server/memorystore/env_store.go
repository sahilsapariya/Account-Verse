package memorystore

import (
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"server/cli"
	"server/constants"
)

// RequiredEnv holds information about required envs
type RequiredEnv struct {
	EnvPath            string `json:"ENV_PATH"`
	DatabaseURL        string `json:"DATABASE_URL"`
	DatabaseType       string `json:"DATABASE_TYPE"`
	DatabaseName       string `json:"DATABASE_NAME"`
	DatabaseHost       string `json:"DATABASE_HOST"`
	DatabasePort       string `json:"DATABASE_PORT"`
	DatabaseUsername   string `json:"DATABASE_USERNAME"`
	DatabasePassword   string `json:"DATABASE_PASSWORD"`
}

// RequiredEnvStore is a simple in-memory store for sessions.
type RequiredEnvStore struct {
	mutex       sync.Mutex
	requiredEnv RequiredEnv
}

// GetRequiredEnv to get required env
func (r *RequiredEnvStore) GetRequiredEnv() RequiredEnv {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.requiredEnv
}

// SetRequiredEnv to set required env
func (r *RequiredEnvStore) SetRequiredEnv(requiredEnv RequiredEnv) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.requiredEnv = requiredEnv
}

var RequiredEnvStoreObj *RequiredEnvStore

// InitRequiredEnv to initialize EnvData and throw error if required env are not present
// This includes env that only configurable via env vars and not the ui
func InitRequiredEnv() error {
	envPath := os.Getenv(constants.EnvKeyEnvPath)

	if envPath == "" {
		if envPath == "" {
			envPath = `.env`
		}
	}

	if cli.ARG_ENV_FILE != nil && *cli.ARG_ENV_FILE != "" {
		envPath = *cli.ARG_ENV_FILE
	}
	log.Info("env path: ", envPath)

	err := godotenv.Load(envPath)
	if err != nil {
		log.Infof("using OS env instead of %s file", envPath)
	}

	dbURL := os.Getenv(constants.EnvKeyDatabaseURL)
	dbType := os.Getenv(constants.EnvKeyDatabaseType)
	dbName := os.Getenv(constants.EnvKeyDatabaseName)
	dbPort := os.Getenv(constants.EnvKeyDatabasePort)
	dbHost := os.Getenv(constants.EnvKeyDatabaseHost)
	dbUsername := os.Getenv(constants.EnvKeyDatabaseUsername)
	dbPassword := os.Getenv(constants.EnvKeyDatabasePassword)

	// set default db name for non sql dbs
	if dbName == "" {
		dbName = "authorizer"
	}

	if strings.TrimSpace(dbType) == "" {
		if cli.ARG_DB_TYPE != nil && *cli.ARG_DB_TYPE != "" {
			dbType = strings.TrimSpace(*cli.ARG_DB_TYPE)
		}

		if dbType == "" {
			log.Debug("DATABASE_TYPE is not set")
			return errors.New("invalid database type. DATABASE_TYPE is empty")
		}
	}

	if strings.TrimSpace(dbURL) == "" {
		if cli.ARG_DB_URL != nil && *cli.ARG_DB_URL != "" {
			dbURL = strings.TrimSpace(*cli.ARG_DB_URL)
		}

		// In dynamoDB these field are not always mandatory
		if dbType != constants.DbTypeDynamoDB && dbURL == "" && dbPort == "" && dbHost == "" && dbUsername == "" && dbPassword == "" {
			log.Debug("DATABASE_URL is not set")
			return errors.New("invalid database url. DATABASE_URL is required")
		}
	}

	if dbName == "" {
		if dbName == "" {
			dbName = "authorizer"
		}
	}

	requiredEnv := RequiredEnv{
		EnvPath:                   envPath,
		DatabaseURL:               dbURL,
		DatabaseType:              dbType,
		DatabaseName:              dbName,
		DatabaseHost:              dbHost,
		DatabasePort:              dbPort,
		DatabaseUsername:          dbUsername,
		DatabasePassword:          dbPassword,
	}

	RequiredEnvStoreObj = &RequiredEnvStore{
		requiredEnv: requiredEnv,
	}

	return nil
}