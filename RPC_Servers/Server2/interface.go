package rpcserver2

import (
	"context"

	gateway_service "github.com/cloudwego/RPC_Servers/kitex_gen/gateway_service"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return exampleServiceServiceInfo
}

var exampleServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ExampleService"
	handlerType := (*gateway_service.ExampleService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ExampleMethod": kitex.NewMethodInfo(exampleMethodHandler, newExampleServiceExampleMethodArgs, newExampleServiceExampleMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "gateway_service",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func exampleMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gateway_service.ExampleServiceExampleMethodArgs)
	realResult := result.(*gateway_service.ExampleServiceExampleMethodResult)
	success, err := handler.(gateway_service.ExampleService).ExampleMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newExampleServiceExampleMethodArgs() interface{} {
	return gateway_service.NewExampleServiceExampleMethodArgs()
}

func newExampleServiceExampleMethodResult() interface{} {
	return gateway_service.NewExampleServiceExampleMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ExampleMethod(ctx context.Context, request *gateway_service.ExampleRequest) (r *gateway_service.ExampleResponse, err error) {
	var _args gateway_service.ExampleServiceExampleMethodArgs
	_args.Request = request
	var _result gateway_service.ExampleServiceExampleMethodResult
	if err = p.c.Call(ctx, "ExampleMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
