package repository

import (
	"context"
	"log"

	"github.com/TallTalha/weather-system/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// WeatherRepository MongoDB ile hava durumu veritabanı işlemlerini yönetir
type WeatherRepository struct {
	DB *mongo.Database
}

// NewWeatherRepository yeni bir WeatherRepository örneği oluşturur
func NewWeatherRepository(db *mongo.Database) *WeatherRepository {
	return &WeatherRepository{
		DB: db,
	}
}

// InsertWeatherData hava durumu verisini MongoDB'ye ekler
func (repo *WeatherRepository) InsertWeatherData(data model.WeatherData) error {
	collection := repo.DB.Collection("weather")
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Printf("Failed to insert weather data: %v", err)
		return err
	}
	log.Println("Weather data inserted successfully")
	return nil
}

// Diğer MongoDB işlemleri burada tanımlanabilir (örneğin, FindWeatherData, UpdateWeatherData vb.)
