package mongodb

import (
	"context"
	"rosenchat/src/configs"
	"rosenchat/src/logger"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var conf = configs.Get()
var log = logger.Get()

var clientOnce = &sync.Once{}
var clientSingleton Client

// Client is a wrapper for *mongo.Client.
type Client *mongo.Client

// GetClient provides the Client singleton.
func GetClient() Client {
	clientOnce.Do(func() {
		connectOpts := options.Client().ApplyURI(conf.Mongo.Addr)

		client, err := mongo.NewClient(connectOpts)
		if err != nil {
			panic("failed to create MongoDB client: " + err.Error())
		}

		timeoutDuration := time.Duration(conf.Mongo.OperationTimeoutSec) * time.Second
		connectCtx, connectCancelFunc := context.WithTimeout(context.Background(), timeoutDuration)
		defer connectCancelFunc()

		if err := client.Connect(connectCtx); err != nil {
			panic("failed to connect to MongoDB: " + err.Error())
		}

		pingCtx, pingCancelFunc := context.WithTimeout(context.Background(), timeoutDuration)
		defer pingCancelFunc()

		if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
			panic("failed to ping MongoDB: %+v" + err.Error())
		}

		log.Infof(context.Background(), "Connected with MongoDB.")
		clientSingleton = client
	})

	return clientSingleton
}
