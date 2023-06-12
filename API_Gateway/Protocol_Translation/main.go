package main

import (
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

func main() {
	// Parse IDL with Local Files
	// YOUR_IDL_PATH thrift file path, eg:./idl/example.thrift
	p, err := generic.NewThriftFileProvider("./service.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	cli, err := genericclient.NewClient("psm", g)
	if err != nil {
		panic(err)
	}
	// 'ExampleMethod' method name must be passed as param
	resp, err := cli.GenericCall(ctx, "ExampleMethod", "{\"Msg\": \"hello\"}")
	// resp is a JSON string
}
