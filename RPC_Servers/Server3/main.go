package main

import (
	"context"
	"fmt"
	"net"

	"github.com/cloudwego/RPC_Servers/kitex_gen/gateway_service"
	server "github.com/cloudwego/kitex/server"
)

type exampleServiceImpl struct{}

func (s *exampleServiceImpl) ExampleMethod(ctx context.Context, request *gateway_service.ExampleRequest) (*gateway_service.ExampleResponse, error) {
	// Your implementation logic for ExampleMethod goes here
	response := &gateway_service.ExampleResponse{
		Message: fmt.Sprintf("Hello, %s!", request.Message),
	}
	return response, nil
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9993")
	handler := &exampleServiceImpl{}
	server := NewServer(handler, server.WithServiceAddr(addr))

	if err := server.Run(); err != nil {
		panic(err)
	}
}
