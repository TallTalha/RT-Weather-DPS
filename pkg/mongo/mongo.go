package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient global değişken olarak MongoDB client'ını saklar
var MongoClient *mongo.Client

// Connect fonksiyonu MongoDB'ye bağlanmayı sağlar
func Connect(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Bağlantıyı kontrol et
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	log.Println("Connected to MongoDB!")
}

// Disconnect fonksiyonu MongoDB bağlantısını keser
func Disconnect() {
	if err := MongoClient.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB.")
}
