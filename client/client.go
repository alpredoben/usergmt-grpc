package main

import (
	"context"
	"fmt"
	"log"
	"time"

	userPB "gouser/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:30001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock());
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}

	defer conn.Close();

	c := userPB.NewUserManagementServicesClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	var new_users = make(map[string]int32)

	new_users["Alice"] = 43
	new_users["Bob"] = 30

	log.Println("User Details")

	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &userPB.UserRequest{
			Name: name,
			Age: age,
		})

		if err != nil {
			log.Fatalf("Could not create user %v", err)
		}

		log.Printf("ID: %d", r.GetId())
		log.Printf("Name: %s", r.GetName())
		log.Printf("Age: %d", r.GetAge())
		log.Println("------------------------")

	}

	params := &userPB.GetUsersParamRequest{}

	r, err := c.GetUsers(ctx, params)

	if err != nil {
		log.Fatalf("Could not retrieve users: %v", err)
	}

	log.Print("\n USER LIST : \n")
	fmt.Printf("r.GetUsers(): %v\n", r.GetUsers())

}