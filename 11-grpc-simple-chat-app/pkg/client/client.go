package client

import (
	"context"
	"grpc-simple-chat-app/api"
	"log"
	"time"

	"google.golang.org/grpc"
)

type ChatClient struct {
	client api.ChatServiceClient
}

func NewChatClient(conn *grpc.ClientConn) *ChatClient {
	return &ChatClient{
		client: api.NewChatServiceClient(conn),
	}
}

func (c *ChatClient) SendMessage(user, content string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.client.SendMessage(ctx, &api.MessageRequest{User: user, Content: content})
	if err != nil {
		log.Printf("could not send message: %v", err)
	}
}

func (c *ChatClient) StreamMessages(user string) {
	stream, err := c.client.StreamMessages(context.Background(), &api.User{Name: user})
	if err != nil {
		log.Fatalf("could not stream messages: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("error receiving message: %v", err)
			return
		}
		log.Printf("%s: %s", msg.User, msg.Content)
	}
}
