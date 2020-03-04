package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"mimoza.tests/user-service-server/proto"
	"net"
	"time"
)

type UsersHandler struct{}

func (s *UsersHandler) Create(ctx context.Context, req *proto.CreateUserRequest) (*proto.User, error) {
	if req.GetUser().GetName() == "petkan" {
		err := grpc.SendHeader(ctx, metadata.Pairs("location", fmt.Sprintf("/api/users/%s", req.GetUser().GetName())))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal server error")
		}

		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("User %s already exists", req.GetUser().GetName()))
	}

	return &proto.User{}, nil
}

func (s *UsersHandler) Get(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {
	if req.GetName() == "petkan" {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("User %s not found", req.GetName()))
	}

	u := &proto.User{}
	u.Name = req.GetName()
	u.Email = "joro@abv.bg"
	u.Avatar = "joro/profile.png"
	u.CreationTime = time.Now().Unix()
	u.LastLoginTime = time.Now().Unix()

	return u, nil
}

func (s *UsersHandler) Update(ctx context.Context, req *proto.UpdateUserRequest) (*proto.User, error) {
	for _, path := range req.GetUpdateMask().GetPaths() {
		fmt.Println(path)
	}

	return nil, nil
}

func (s *UsersHandler) Delete(ctx context.Context, req *proto.DeleteUserRequest) (*empty.Empty, error) {
	return nil, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9991")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterUsersHandlerServer(srv, &UsersHandler{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
