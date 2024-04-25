package service

import (
	"C:\Users\talha\OneDrive\Belgeler\Go_Projects\weather-system\dto"
	"weather-system/model"
	"weather-system/pkg/mongo"
	"weather-system/pkg/rabbitmq"
)

// WeatherService hava durumu ile ilgili işlemleri tanımlar
type WeatherService struct {
	MongoClient    *mongo.MongoClient
	RabbitMQClient *rabbitmq.RabbitMQClient
}

// NewWeatherService yeni bir WeatherService örneği oluşturur
func NewWeatherService(mongoClient *mongo.MongoClient, rabbitMQClient *rabbitmq.RabbitMQClient) *WeatherService {
	return &WeatherService{
		MongoClient:    mongoClient,
		RabbitMQClient: rabbitMQClient,
	}
}

// ProcessWeatherData hava durumu verilerini işler ve gerektiğinde başka sistemlere gönderir
func (ws *WeatherService) ProcessWeatherData(data dto.WeatherData) {
	// DTO'dan model verisine dönüştürme
	weatherModel := model.WeatherData{
		Temperature: data.Temperature,
		Humidity:    data.Humidity,
		City:        data.City,
		// Timestamp dönüşümü gerekirse burada yapılabilir
	}

	// MongoDB'ye veri kaydet
	ws.MongoClient.InsertWeatherData(weatherModel)

	// RabbitMQ üzerinden mesaj gönder
	message := []byte("New weather data processed for " + data.City)
	ws.RabbitMQClient.PublishMessage("weatherDataQueue", message)
}
