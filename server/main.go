package main

import (
	"context"
	"log"
	"net"

	foo "github.com/aron-huang-vertiv/grpc_client_create_connection_test/proto/foo"
	"google.golang.org/grpc"
)

const (
	port = ":7788"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	foo.RegisterFooServiceServer(s, &FooService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type FooService struct {
	foo.UnimplementedFooServiceServer
}

func NewFooService() *FooService {
	return &FooService{}
}

func (f FooService) Foo(ctx context.Context, req *foo.FooReq) (*foo.FooResp, error) {
	log.Printf("Hi,%v-%v", req.Name, req.Count)
	return &foo.FooResp{
		Msg:   "Hi " + req.Name,
		Count: req.Count,
	}, nil
}
