package main

import (
    "context"
    "io"
    "log"
    "time"

    pb "github.com/sanjay7178/go-basics/grpc_test/machineinfo"

    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewMachineInfoServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    stream, err := c.GetMachineInfo(ctx)
    if err != nil {
        log.Fatalf("could not get machine info: %v", err)
    }

    go func() {
        for {
            response, err := stream.Recv()
            if err == io.EOF {
                return
            }
            if err != nil {
                log.Fatalf("failed to receive: %v", err)
            }

            log.Printf("Machine ID: %s", response.GetMachineID())
            log.Printf("Username: %s", response.GetUsername())
            log.Printf("Password: %s", response.GetPassword())

            // Sending back status to server
            status := &pb.MachineInfoResponse{
                Status: "Info received successfully",
            }
            if err := stream.Send(status); err != nil {
                log.Fatalf("failed to send status: %v", err)
            }
        }
    }()

    <-time.After(30 * time.Second)
}
