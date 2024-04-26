package service

import (
	"log"

	"github.com/TallTalha/weather-system/dto"
	"github.com/TallTalha/weather-system/internal/repository"
	"github.com/TallTalha/weather-system/model"
	"github.com/TallTalha/weather-system/pkg/rabbitmq"
)

type WeatherService struct {
	WeatherRepo    *repository.WeatherRepository
	RabbitMQClient *rabbitmq.RabbitMQClient
}

func NewWeatherService(weatherRepo *repository.WeatherRepository, rabbitMQClient *rabbitmq.RabbitMQClient) *WeatherService {
	return &WeatherService{
		WeatherRepo:    weatherRepo,
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

	// MongoDB'ye veri kaydetmek için WeatherRepository kullan
	if err := ws.WeatherRepo.InsertWeatherData(weatherModel); err != nil {
		log.Printf("Error inserting weather data: %v", err)
	}

	// RabbitMQ üzerinden mesaj gönder
	message := []byte("New weather data processed for " + data.City)
	ws.RabbitMQClient.PublishMessage("weatherDataQueue", message)
}
