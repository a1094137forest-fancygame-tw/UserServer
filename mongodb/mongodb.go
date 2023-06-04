package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"UserServer/config"
)

var (
	PoolWr PoolWrapper
)

type DB interface {
	Insert(ctx context.Context, document interface{}) error
}

type PoolWrapper struct {
	master *mongo.Client
	slave  *mongo.Client
}

func (p *PoolWrapper) Write(collection string) *mongo.Collection {
	return (p.master.Database(config.MongoDBName).Collection(collection))
}

func (p *PoolWrapper) Read(collection string) *mongo.Collection {
	if p.slave == nil {
		return (p.master.Database(config.MongoDBName).Collection(collection))
	}
	return (p.slave.Database(config.MongoDBName).Collection(collection))
}

func Initialize() {
	log.Printf("[info] Connecting to MongoDB @ %s", config.MongoDBMaster)
	masterUri := fmt.Sprintf("mongodb://%s", config.MongoDBMaster)
	PoolWr.master = getPool(masterUri)
}

func getPool(uri string) *mongo.Client {
	options := options.Client().ApplyURI(uri).
		SetAuth(options.Credential{
			Username:      config.MongoDBUser,
			Password:      config.MongoDBPassword,
			AuthMechanism: config.MongoDBAuthMechanism,
			AuthSource:    config.MongoDBAuthSource,
		}).
		SetMaxPoolSize(100).
		SetConnectTimeout(30000 * time.Second)
	client, err := mongo.NewClient(options)
	if err != nil {
		panic(err)
	}

	return setPool(client)
}

func setPool(client *mongo.Client) *mongo.Client {

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return client
}
