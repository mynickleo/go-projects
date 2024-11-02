package server

import (
	"context"
	"grpc-simple-chat-app/api"
	"log"
	"sync"
)

type ChatServer struct {
	api.UnimplementedChatServiceServer
	clients map[string]chan *api.MessageResponse
	mu      sync.Mutex
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(map[string]chan *api.MessageResponse),
	}
}

func (s *ChatServer) SendMessage(ctx context.Context, req *api.MessageRequest) (*api.MessageResponse, error) {
	message := &api.MessageResponse{
		User:    req.User,
		Content: req.Content,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for user, clientChan := range s.clients {
		select {
		case clientChan <- message:
			log.Printf("Message sent to user: %s", user)
		default:
			log.Printf("User %s's channel is full, message dropped", user)
		}
	}

	return message, nil
}

func (s *ChatServer) StreamMessages(user *api.User, stream api.ChatService_StreamMessagesServer) error {
	s.mu.Lock()
	msgChan := make(chan *api.MessageResponse, 10)
	s.clients[user.Name] = msgChan
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, user.Name)
		close(msgChan)
		s.mu.Unlock()
		log.Printf("User %s disconnected", user.Name)
	}()

	for msg := range msgChan {
		if err := stream.Send(msg); err != nil {
			log.Printf("Error sending message to user %s: %v", user.Name, err)
			return err
		}
	}
	return nil
}
