package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectDatabase() error {
	uri := "mongodb+srv://anhhuy:%40nhHuy1010@cluster0.fmkkfsr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	fmt.Println("connect")

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Kết nối thành công tới MongoDB!")
	db = client.Database("todolist")
	return nil
}
func GetInstance() *mongo.Database {
	return db
}
