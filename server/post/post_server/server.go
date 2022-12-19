package main

import (
	"fmt"
	pd "github.com/weichenluo/Twitter-Raft/server/post/post_driver"
	"github.com/weichenluo/Twitter-Raft/server/post/postpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("server started")
	postPort := "localhost:4001"
	log.Println("postPort =", postPort)
	lis, err := net.Listen("tcp", postPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	postpb.RegisterPostServiceServer(s, &pd.Server{})

	pd.Init()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
