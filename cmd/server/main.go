package main

import (
	"context"
	"log"
	"net"

	pb "github.com/TallTalha/weather-system/proto/pb"
	"google.golang.org/grpc"
)

// server sınıfımız UnimplementedWeatherServiceServer'ı genişletir.
type server struct {
	pb.UnimplementedWeatherServiceServer
	// Burada MongoDB ve RabbitMQ bağlantılarınızı yönetecek alanlar tanımlayabilirsiniz.
}

// GetWeatherData RPC metodunu gerçekleştirir.
func (s *server) GetWeatherData(ctx context.Context, req *pb.WeatherRequest) (*pb.WeatherData, error) {
	// Gerçek veri dönüşünüz burada işlenir.
	// Örnek dönüş:
	return &pb.WeatherData{City: req.City, Temperature: 22.5, Timestamp: "2024-04-25T12:34:56Z"}, nil
}

// SendWeatherData RPC metodunu gerçekleştirir.
func (s *server) SendWeatherData(ctx context.Context, data *pb.WeatherData) (*pb.WeatherResponse, error) {
	// Veriyi işleyin ve MongoDB'ye kaydedin veya RabbitMQ üzerinden yayınlayın.
	log.Printf("Received data for city %s with temperature %f", data.City, data.Temperature)
	// Örnek dönüş:
	return &pb.WeatherResponse{Message: "Data received successfully"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterWeatherServiceServer(grpcServer, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
