package user_service

import (
	"github.com/TwiLightDM/diploma-user-service/proto/userservicepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type UserClient struct {
	user userservicepb.UserServiceClient
	conn *grpc.ClientConn
}

func NewUserClient(address string) *UserClient {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}

	return &UserClient{
		user: userservicepb.NewUserServiceClient(conn),
		conn: conn,
	}
}

func (c *UserClient) Close() error {
	return c.conn.Close()
}
