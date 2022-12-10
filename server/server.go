package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	userPB "gouser/proto"

	"google.golang.org/grpc"
)

const (
	port = ":30001"
)

type UserManagementServer struct {
	userPB.UnimplementedUserManagementServicesServer
	list_user *userPB.UserListResponses
}

func NewUserManagement() *UserManagementServer {
	return &UserManagementServer {
		list_user: &userPB.UserListResponses{},
	}
}




func(s *UserManagementServer) CreateNewUser(ctx context.Context, protoBuff *userPB.UserRequest) (*userPB.UserResponse, error) {
	log.Printf("Received : %v", protoBuff.GetName())

	var userId int32 = int32(rand.Intn(1000))
	createdUser := &userPB.UserResponse{
		Name: protoBuff.GetName(),
		Age: protoBuff.GetAge(),
		Id: userId,
	}

	s.list_user.Users = append(s.list_user.Users, (*userPB.User)(createdUser))
	return &userPB.UserResponse{
		Name: protoBuff.GetName(),
		Age: protoBuff.GetAge(),
		Id: userId,
	}, nil
}

func(s *UserManagementServer) GetUsers(ctx context.Context, protoBuff *userPB.GetUsersParamRequest) (*userPB.UserListResponses, error) {
	return s.list_user, nil
}


func(server *UserManagementServer) Run() error {
	netListen, err :=  net.Listen("tcp", port);

	if err != nil {
		log.Fatalf("Failied to listen %v", err)
	}

	s := grpc.NewServer();
	userPB.RegisterUserManagementServicesServer(s, server)
	log.Printf("Server listening on port %v", netListen.Addr())

	return s.Serve(netListen)
}




func main() {
	var userServer *UserManagementServer = NewUserManagement()

	if err := userServer.Run(); err != nil {
		log.Fatalf("Failed to server %v", err)
	}


}
