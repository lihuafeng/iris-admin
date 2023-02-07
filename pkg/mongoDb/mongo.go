package mongoDb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

var (
	once sync.Once
	MongoDbClient *mongo.Client
	dataBase   = "goTest" // 数据库
	collection = "test"
)

func Instance() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var err error
	once.Do(func() {
		// "mongodb://user:password@localhost:27017".
		//credential := options.Credential{
		//		Username: "user",
		//		Password: "password",
		//	}
		//	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
		MongoDbClient, err= mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			panic(err)
		}
	})
	err = MongoDbClient.Ping(ctx, readpref.Primary())
	if err != nil{
		panic(err)
	}
	return MongoDbClient
}