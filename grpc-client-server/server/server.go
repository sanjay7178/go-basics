package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/sanjay7178/go-basics/grpc-client-server/grpcbidir"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
    pb.UnimplementedYourServiceServer
    mu      sync.Mutex
    clients map[string]chan *pb.Message
}

func newServer() *server {
    return &server{
        clients: make(map[string]chan *pb.Message),
    }
}

func (s *server) RegisterClient(ctx context.Context, info *pb.ClientInfo) (*pb.ClientResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, exists := s.clients[info.ClientId]; !exists {
        s.clients[info.ClientId] = make(chan *pb.Message, 10)
        log.Printf("Registered client with ID: %s", info.ClientId)
    }

    return &pb.ClientResponse{Message: "Client registered successfully"}, nil
}

func (s *server) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if ch, exists := s.clients[req.ClientId]; exists {
        ch <- &pb.Message{Message: req.Message}
        return &pb.MessageResponse{Status: "Message sent successfully"}, nil
    }

    return &pb.MessageResponse{Status: "Client not found"}, nil
}

func (s *server) ReceiveMessages(req *pb.ReceiveRequest, stream pb.YourService_ReceiveMessagesServer) error {
    s.mu.Lock()
    ch, exists := s.clients[req.ClientId]
    s.mu.Unlock()

    if !exists {
        return nil
    }

    for msg := range ch {
        if err := stream.Send(msg); err != nil {
            return err
        }
    }

    return nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterYourServiceServer(s, newServer())
    reflection.Register(s)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
