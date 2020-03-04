package proto

//go:generate protoc --go_out=plugins=grpc:. user_service.proto
//go:generate protoc --grpc-gateway_out=logtostderr=true,grpc_api_configuration=user_service.yaml:. user_service.proto
