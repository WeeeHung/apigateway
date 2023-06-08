package main

import (
	server1 "github.com/cloudwego/RPC_Servers/server1"
	server2 "github.com/cloudwego/RPC_Servers/server2"
	server3 "github.com/cloudwego/RPC_Servers/server3"
	server4 "github.com/cloudwego/RPC_Servers/server4"
	"github.com/cloudwego/kitex/server"
)

func main() {
	// Create the handlers for each server
	handler1 := &server1.ExampleService{}
	handler2 := &server2.ExampleService{}
	handler3 := &server3.ExampleService{}
	handler4 := &server4.ExampleService{}

	// Create the servers
	server1 := server1.NewServer(handler1).Start()
	server2 := server2.NewServer(handler2).Start()
	server3 := server3.NewServer(handler3).Start()
	server4 := server4.NewServer(handler4).Start()

	// Register the servers with Kitex
	server.RegisterServers(server1, server2, server3, server4)

	// Start the servers
	if err := server.Run(); err != nil {
		panic(err)
	}
}
