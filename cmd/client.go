package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"

	pb "github.com/TallTalha/weather-system/protos/api" // Proto dosyalarının yolu
)

const (
	grpcAddress = "localhost:50051"        // gRPC sunucusunun adresi
	wsAddress   = "ws://localhost:8080/ws" // WebSocket sunucusunun adresi
	defaultCity = "Istanbul"               // Varsayılan şehir
)

func connectWebSocket() (*websocket.Conn, error) {
	u, err := url.Parse(wsAddress)
	if err != nil {
		return nil, fmt.Errorf("could not parse wsAddress: %v", err)
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("dial: %v", err)
	}
	return c, nil
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// gRPC sunucusuna bağlan
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// gRPC istemcisini oluştur
	grpcClient := pb.NewWeatherServiceClient(conn)

	// WebSocket sunucusuna bağlan
	wsConn, err := connectWebSocket()
	if err != nil {
		log.Fatalf("failed to connect to WebSocket: %v", err)
	}
	defer wsConn.Close()

	go func() {
		defer wsConn.Close()
		for {
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	// QueryWeather RPC'sini çağır ve WebSocket üzerinden gönder
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := grpcClient.QueryWeather(ctx, &pb.WeatherRequest{City: defaultCity})
	if err != nil {
		log.Fatalf("could not query weather: %v", err)
	}

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the WebSocket connection by sending a close message.
			err := wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		default:
			// Receive data from gRPC stream
			weather, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving from stream: %v", err)
			}

			// Send data to WebSocket connection
			err = wsConn.WriteJSON(weather)
			if err != nil {
				log.Fatalf("Error sending to WebSocket: %v", err)
			}
			time.Sleep(2 * time.Second) // Rate limit for demonstration purposes
		}
	}
}
