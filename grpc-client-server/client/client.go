package main

import (
    "context"
    "log"
    "time"

    pb "github.com/sanjay7178/go-basics/grpc-client-server/grpcbidir"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewYourServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    clientID := "client1"
    _, err = client.RegisterClient(ctx, &pb.ClientInfo{ClientId: clientID})
    if err != nil {
        log.Fatalf("could not register client: %v", err)
    }

    go func() {
        stream, err := client.ReceiveMessages(context.Background(), &pb.ReceiveRequest{ClientId: clientID})
        if err != nil {
            log.Fatalf("could not receive messages: %v", err)
        }

        for {
            msg, err := stream.Recv()
            if err != nil {
                log.Fatalf("could not receive message: %v", err)
            }
            log.Printf("Received message: %s", msg.Message)
        }
    }()

    _, err = client.SendMessage(ctx, &pb.MessageRequest{ClientId: clientID, Message: "Hello, server!"})
    if err != nil {
        log.Fatalf("could not send message: %v", err)
    }

    // Keep the client running to receive messages
    select {}
}
