package main

import (
	"log"
	"net"

	myGrpc "github.com/TallTalha/weather-system/internal/grpc"
	"github.com/TallTalha/weather-system/internal/repository"
	"github.com/TallTalha/weather-system/internal/service"
	"github.com/TallTalha/weather-system/pkg/mongo"
	"github.com/TallTalha/weather-system/pkg/rabbitmq"
	pb "github.com/TallTalha/weather-system/proto/pb"
	"google.golang.org/grpc"
)

func main() {
	// MongoDB bağlantısı
	mongoClient := mongo.NewMongoClient("mongodb://localhost:27017")
	defer mongoClient.Disconnect()

	// RabbitMQ bağlantısı
	rabbitMQClient := rabbitmq.NewRabbitMQClient("amqp://guest:guest@localhost:5672/")
	defer rabbitMQClient.Close()

	// `WeatherRepository` örneğini `mongoClient` kullanarak oluşturun
	weatherRepo := repository.NewWeatherRepository(mongoClient.Client.Database("weatherDatabase"))

	// Servis ve repository katmanlarının oluşturulması
	weatherService := service.NewWeatherService(weatherRepo, rabbitMQClient)

	// gRPC sunucu ayarları
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, myGrpc.NewWeatherServiceServer(weatherService))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
