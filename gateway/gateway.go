package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"

	gw "mimoza.tests/user-service-server/proto"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", ":9991", "gRPC server endpoint")
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	err := gw.RegisterUsersHandlerHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}
