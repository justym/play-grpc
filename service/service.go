package service

import (
	"context"
	"log"
)

type Server struct{}

func (s *Server) SayMessage(ctx context.Context, msg *Message) (*Message, error) {
	log.Printf("[Received Message]: %s", msg.Body)
	return &Message{Body: "Hello gRPC !!"}, nil
}
