namespace go gateway_service

struct ExampleRequest {
  1: required string message
}

struct ExampleResponse {
  1: required string message
}

service ExampleService {
  ExampleResponse ExampleMethod(1: ExampleRequest request)
}