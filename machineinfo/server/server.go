package main

import (
    // "context"
    "log"
    "net"
    "time"

    pb "github.com/sanjay7178/go-basics/grpc_test/machineinfo"

    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedMachineInfoServiceServer
    infoChannel chan *pb.MachineInfoResponse
}

func NewServer() *server {
    return &server{
        infoChannel: make(chan *pb.MachineInfoResponse),
    }
}

func (s *server) GetMachineInfo(req *pb.MachineInfoRequest, stream pb.MachineInfoService_GetMachineInfoServer) error {
    for info := range s.infoChannel {
        if err := stream.Send(info); err != nil {
            return err
        }
    }
    return nil
}

func (s *server) triggerInfo() {
    info := &pb.MachineInfoResponse{
        MachineID: "12345",
        Username:  "admin",
        Password:  "password",
        Status:    "Info sent successfully",
    }
    s.infoChannel <- info
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := NewServer()
    grpcServer := grpc.NewServer()
    pb.RegisterMachineInfoServiceServer(grpcServer, s)
    log.Printf("server listening at %v", lis.Addr())

    go func() {
        for {
            time.Sleep(10 * time.Second) // Simulate triggering every 10 seconds
            s.triggerInfo()
        }
    }()

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
