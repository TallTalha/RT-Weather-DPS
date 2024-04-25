package service

import (
	"log"

	"github.com/TallTalha/weather-system/dto"
	"github.com/TallTalha/weather-system/model"
	"github.com/TallTalha/weather-system/pkg/mongo"
	"github.com/TallTalha/weather-system/pkg/rabbitmq"
)

type WeatherService struct {
	MongoClient    *mongo.MongoClient
	RabbitMQClient *rabbitmq.RabbitMQClient
}

func NewWeatherService(mongoClient *mongo.MongoClient, rabbitMQClient *rabbitmq.RabbitMQClient) *WeatherService {
	return &WeatherService{
		MongoClient:    mongoClient,
		RabbitMQClient: rabbitMQClient,
	}
}

func (ws *WeatherService) ProcessWeatherData(data dto.WeatherData) {
	// DTO'dan model verisine dönüştürme
	weatherModel := model.WeatherData{
		Temperature: data.Temperature,
		City:        data.City,
		// Timestamp dönüşümü burada yapılabilir
	}

	// MongoDB'ye veri kaydet
	// Repository kullanarak veri kaydet
	if err := ws.WeatherRepo.InsertWeatherData(data); err != nil {
		log.Printf("Error inserting weather data: %v", err)
	}

	// RabbitMQ üzerinden mesaj gönder
	message := []byte("New weather data processed for " + data.City)
	ws.RabbitMQClient.PublishMessage("weatherDataQueue", message)
}
