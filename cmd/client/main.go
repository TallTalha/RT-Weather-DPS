package main

import (
	"context"
	"log"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"

	pb "github.com/TallTalha/weather-system/proto/pb"
)

func main() {
	// WebSocket bağlantısı
	wsURL := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
	}
	defer conn.Close()

	// gRPC bağlantısı
	gRPCURL := "localhost:50051"
	grpcConn, err := grpc.Dial(gRPCURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer grpcConn.Close()

	client := pb.NewWeatherServiceClient(grpcConn)

	// WebSocket'ten veri alıp gRPC'ye gönderme
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("Received via WebSocket: %s\n", message)

		// gRPC üzerinden sunucuya veri gönderme
		response, err := client.SendWeatherData(context.Background(), &pb.WeatherData{
			City: string(message), // Örnek olarak mesajı city alanına yerleştirildi
		})
		if err != nil {
			log.Fatalf("Could not send weather data: %v", err)
		}
		log.Printf("Received from gRPC server: %s\n", response)
	}
}
