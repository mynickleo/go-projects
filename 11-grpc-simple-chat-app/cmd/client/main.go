package main

import (
	"context"
	"grpc-simple-chat-app/pkg/client"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, ":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	chatClient := client.NewChatClient(conn)

	go chatClient.StreamMessages("client1")

	log.Println("Connected to gRPC server")
	chatClient.SendMessage("client1", "Hello, world!")
	log.Println("Message sent from client")
}
