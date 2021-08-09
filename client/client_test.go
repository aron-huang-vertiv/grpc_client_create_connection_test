package client

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	foo "github.com/aron-huang-vertiv/grpc_client_create_connection_test/proto/foo"
)

const (
	address = "localhost:7788"
)

func BenchmarkOneConnection(b *testing.B) {
	conn := CreateNewConnect(address)
	defer conn.Close()
	c := foo.NewFooServiceClient(conn)
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
			defer cancel()
			_, err := c.Foo(ctx, &foo.FooReq{
				Name:  "OneConnection",
				Count: int64(ii),
			})
			if err != nil {
				log.Printf("count %v, err:%v", ii, err)
			}
			// log.Printf("%v", resp)
		}(i)
	}
	wg.Wait()
}

func BenchmarkManyConnections(b *testing.B) {
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			conn := CreateNewConnect(address)
			defer conn.Close()
			c := foo.NewFooServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
			defer cancel()
			_, err := c.Foo(ctx, &foo.FooReq{
				Name:  "ManyConnection",
				Count: int64(ii),
			})
			if err != nil {
				log.Printf("count %v, err:%v", ii, err)
			}
			// log.Printf("%v", resp)
		}(i)
	}
	wg.Wait()
}
