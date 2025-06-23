package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port         string
    DBType       string
    DBName       string
    DBHost       string
    DBPort       string
    DBUser       string
    DBPassword   string
    MongoURI     string
    MongoDatabase string
}

func LoadConfig() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    return &Config{
        Port:         getEnv("PORT", "8080"),
        DBType:       getEnv("DB_TYPE", "sqlite"),
        DBName:       getEnv("DB_NAME", "app.db"),
        DBHost:       getEnv("DB_HOST", "localhost"),
        DBPort:       getEnv("DB_PORT", "5432"),
        DBUser:       getEnv("DB_USER", ""),
        DBPassword:   getEnv("DB_PASSWORD", ""),
        MongoURI:     getEnv("MONGO_URI", "mongodb://localhost:27017"),
        MongoDatabase: getEnv("MONGO_DATABASE", "myapp"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
