package main

import (
	"fmt"
	ud "github.com/weichenluo/Twitter-Raft/server/user/user_driver"
	"github.com/weichenluo/Twitter-Raft/server/user/userpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("server started")
	userPort := "localhost:3001"
	log.Println("userPort =", userPort)
	lis, err := net.Listen("tcp", userPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &ud.Server{})

	ud.Init()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
