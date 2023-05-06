package main

import (
	"context"
	"e/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}   //The server struct implements the AddServiceServer interface defined in the proto package

func main() {
	listener, err := net.Listen("tcp", ":4040") //sets up a new TCP listener  
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()  //creates a new gRPC server instance
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)  

	if e := srv.Serve(listener); e != nil {   ///the server is started by calling the Serve() method on the gRPC server instance
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}
