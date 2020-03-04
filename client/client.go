package main

import (
	"context"
	"fmt"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc"
	"time"
)
import "mimoza.tests/user-service-server/proto"

func main() {
	conn, err := grpc.Dial(":9991", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewUsersHandlerClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	req := &proto.UpdateUserRequest{User: &proto.User{Email: "fake@abv.bg"},
									UpdateMask: &field_mask.FieldMask{Paths: []string{"email"}}}

	if response, err := client.Update(ctx, req); err != nil {
		panic(err)
	} else {
		fmt.Println(response)
	}
}
