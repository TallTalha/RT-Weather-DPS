package grpc

import (
	"log"
	"net"

	pb "github.com/TallTalha/weather-system/proto/pb" // Önceki adımda oluşturduğunuz protobuf paket yolu
	"google.golang.org/grpc"
)

// WeatherServiceServer gRPC servisini tanımlar
type WeatherServiceServer struct {
	pb.UnimplementedWeatherServiceServer
}

// NewServer yeni bir gRPC sunucusu örneği oluşturur
func NewServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterWeatherServiceServer(server, &WeatherServiceServer{})
	return server
}

// StartServer gRPC sunucusunu belirli bir portta başlatır
func StartServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := NewServer()
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
