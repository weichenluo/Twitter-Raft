package handler

import (
	"github.com/weichenluo/Twitter-Raft/server/post/postpb"
	"github.com/weichenluo/Twitter-Raft/server/user/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Config struct {
	Clients struct {
		UserDB userpb.UserServiceClient
		PostDB postpb.PostServiceClient
	}

	Connections struct {
		UserPost *grpc.ClientConn
		User     *grpc.ClientConn
	}

	ServerAddresses struct {
		PostPort string
		UserPort string
	}
}

var connector Config

var page Page

func (c *Config) RegisterClients() {
	postPort := "4001"
	userPort := "3001"

	c.ServerAddresses.PostPort = postPort
	c.ServerAddresses.UserPort = userPort
}

func (c *Config) GetUserClient() userpb.UserServiceClient {
	return c.Clients.UserDB
}

func (c *Config) GetPostClient() postpb.PostServiceClient {
	return c.Clients.PostDB
}

func (c *Config) DialServers() {
	option := grpc.WithTransportCredentials(insecure.NewCredentials())
	var err error
	c.Connections.UserPost, err = grpc.Dial("localhost:"+c.ServerAddresses.PostPort, option)
	if err != nil {
		log.Fatalf("could not connect to UserPost Service: %v", err)
	} else {
		c.Clients.PostDB = postpb.NewPostServiceClient(c.Connections.UserPost)
		log.Println("SERVER: Successfully created a connection to User Post Service at", c.ServerAddresses.PostPort)
	}

	c.Connections.User, err = grpc.Dial("localhost:"+c.ServerAddresses.UserPort, option)
	if err != nil {
		log.Fatalf("could not connect to User Service: %v", err)
	} else {
		c.Clients.UserDB = userpb.NewUserServiceClient(c.Connections.User)
		log.Println("SERVER: Successfully created a connection to User Service at", c.ServerAddresses.UserPort)
	}
}

func InitConnectors() {
	connector.RegisterClients()
	connector.DialServers()
}
