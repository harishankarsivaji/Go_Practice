package main

import (
	"context"
	"gRPC/proto"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := rv.Serve(listener); e != nil {
		panic(err)
	}

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return $proto.Response{Result: result}, nil

}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return $proto.Response{Result: result}, nil

}
