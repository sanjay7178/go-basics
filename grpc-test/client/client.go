package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	pb "github.com/sanjay7178/go-basics/grpc_test/grpc_test"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func printGreeter(client pb.GreeterClient ,name string) {
	ctx  , cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r , err := client.SayHello(ctx , &pb.HelloRequest{Name: name})
	if err != nil {
		log.WithField("error stacktrace" , fmt.Sprintf("%+v", err)).Errorf("%v", err)
		panic(err)
	}
	println(r.Message)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = "testdata/ca.pem"
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			panic(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.WithField("error stacktrace" , fmt.Sprintf("%+v", err)).Errorf("%v", err)
		panic(err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	printGreeter(client , "world")
}

