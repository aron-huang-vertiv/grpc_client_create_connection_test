package client

import (
	"log"

	"google.golang.org/grpc"
)



func CreateNewConnect(address string) *grpc.ClientConn{
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect : %v" ,err)
	}
	return conn 
}
