package grpc

import (
	"context"
	"log"
	"time"

	pb "github.com/TallTalha/weather-system/proto/pb" // Önceki adımda oluşturduğunuz protobuf paket yolu
	"google.golang.org/grpc"
)

// WeatherClient gRPC istemcisini tanımlar
type WeatherClient struct {
	Client pb.WeatherServiceClient
}

// NewWeatherClient yeni bir gRPC istemcisi örneği oluşturur
func NewWeatherClient(conn *grpc.ClientConn) *WeatherClient {
	client := pb.NewWeatherServiceClient(conn)
	return &WeatherClient{Client: client}
}

// GetWeatherData belirli bir şehir için hava durumu verisi alır
func (wc *WeatherClient) GetWeatherData(city string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := wc.Client.GetWeatherData(ctx, &pb.WeatherRequest{City: city})
	if err != nil {
		log.Fatalf("could not get weather: %v", err)
	}
	log.Printf("Weather in %s: %s", city, r)
}
