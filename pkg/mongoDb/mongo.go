package mongoDb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	uri        = "mongodb://127.0.0.1:27017/?maxPoolSize=20&w=majority"
	mon        *mongo.Client
	dataBase   = "goTest" // 数据库
	collection = "test"
)

func NewMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	​
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	​
	if err = client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	fmt.Println("successfully connected and pinged.")
	return client
}