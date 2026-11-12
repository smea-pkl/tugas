package main

// import (
// 	"context"
// 	"fmt"
// 	"net"
// 	usersv1 "pkl/backend/gen/pkl/users/v1"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"
// )

// type UserServer struct {
// 	usersv1.UnimplementedUserServiceServer
// }

// func (u *UserServer) Login(ctx context.Context, req *usersv1.LoginRequest) (*usersv1.LoginResponse, error) {
// 	username := req.Username
// 	password := req.Password
// 	if username != "heri" && password != "heri7777" {
// 		return &usersv1.LoginResponse{
// 			Token: "Tidak Bisa Login!",
// 		}, nil
// 	}
// 	token := "123456"
// 	return &usersv1.LoginResponse{
// 		Token: token,

// 	}, nil
// }

// func main() {
// 	listener, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		panic(err)
// 	}

// 	grpcServer := grpc.NewServer()

// 	usersv1.RegisterUserServiceServer(
// 		grpcServer,
// 		&UserServer{},
// 	)

// 	reflection.Register(grpcServer)

// 	fmt.Println("gRPC server running on :50051")

// 	if err := grpcServer.Serve(listener); err != nil {
// 		panic(err)
// 	}
// }