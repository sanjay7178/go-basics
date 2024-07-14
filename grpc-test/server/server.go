package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"strconv"

	pb "github.com/sanjay7178/go-basics/grpc_test/grpc_test"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func newServer() *server {
	s := &server{}
	return s
}

func main() {
	flag.Parse()
	list, err := net.Listen("tcp", "localhost:"+strconv.Itoa(*port))
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = "testdata/server1.pem"
		}
		if *keyFile == "" {
			*keyFile = "testdata/server1.key"
		}
		if *jsonDBFile == "" {
			*jsonDBFile = "testdata/route_guide_db.json"
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			panic(err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(grpcServer, newServer())
	fmt.Printf("server is running on localhost:%d", *port)
	grpcServer.Serve(list)

}
