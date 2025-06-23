package mongodb

import (
    "context"
    "server/config"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(cfg *config.Config) (*mongo.Database, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
    if err != nil {
        return nil, err
    }

    // Test the connection
    if err := client.Ping(ctx, nil); err != nil {
        return nil, err
    }

    return client.Database(cfg.MongoDatabase), nil
}
