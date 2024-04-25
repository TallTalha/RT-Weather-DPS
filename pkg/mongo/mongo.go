package mongo

import (
	"context"
	"log"

	"github.com/TallTalha/weather-system/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient yapısı MongoDB bağlantısını ve ilgili işlemleri yönetir
type MongoClient struct {
	Client *mongo.Client
}

// NewMongoClient, yeni bir MongoClient örneği oluşturur ve döndürür
func NewMongoClient(uri string) *MongoClient {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	// Bağlantıyı kontrol et
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	return &MongoClient{
		Client: client,
	}
}

// Disconnect, MongoDB bağlantısını keser
func (mc *MongoClient) Disconnect() {
	if err := mc.Client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Failed to disconnect MongoDB:", err)
	}
}

// InsertWeatherData, verilen WeatherData nesnesini MongoDB'ye ekler
func (mc *MongoClient) InsertWeatherData(data model.WeatherData) error {
	collection := mc.Client.Database("yourDatabaseName").Collection("weather")
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Printf("Failed to insert weather data: %v", err)
		return err
	}
	log.Println("Weather data inserted successfully")
	return nil
}
