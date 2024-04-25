package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	// Protobuf ve gRPC kütüphanelerini import edin.
	pb "github.com/TallTalha/weather-system/protos/api"
)

// server is used to implement WeatherServiceServer from the protobuf definition.
type server struct {
	pb.UnimplementedWeatherServiceServer
	mongoClient *mongo.Client
	rabbitConn  *amqp.Connection
}

// QueryWeather implements the QueryWeather RPC method.
func (s *server) QueryWeather(req *pb.WeatherRequest, stream pb.WeatherService_QueryWeatherServer) error {
	// Şimdilik basit bir mock veri döndürülüyor.
	// Gerçek uygulamada burası veri tabanından sorgu yapılacak.
	mockResponses := []*pb.WeatherResponse{
		{Temperature: 24.5, Humidity: 30.0},
		{Temperature: 26.0, Humidity: 45.0},
	}
	for _, res := range mockResponses {
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second) // Stream'de veri gönderme aralığını simüle etmek için.
	}
	return nil
}

// ListCities implements the ListCities RPC method.
func (s *server) ListCities(ctx context.Context, req *pb.ListCitiesRequest) (*pb.ListCitiesResponse, error) {
	// Gerçek uygulamada burası veri tabanından şehir listesi sorgulanacak.
	mockCities := []string{"Istanbul", "Ankara", "Izmir"}
	items := make([]*pb.CityEntry, len(mockCities))
	for i, city := range mockCities {
		items[i] = &pb.CityEntry{CityCode: fmt.Sprintf("TR-%d", i), CityName: city}
	}
	return &pb.ListCitiesResponse{Items: items}, nil
}

func main() {
	// MongoDB ve RabbitMQ ile bağlantı kurun.
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	rabbitConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitConn.Close()

	// gRPC sunucusunu başlatın.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Sunucumuzu gRPC sunucusuna kaydedin.
	pb.RegisterWeatherServiceServer(grpcServer, &server{
		mongoClient: mongoClient,
		rabbitConn:  rabbitConn,
	})

	// Sunucuyu başlatın ve gelen bağlantıları dinleyin.
	log.Println("Server is listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 50051: %v", err)
	}
}
