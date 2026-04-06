package user_service

import (
	"log"

	"github.com/TwiLightDM/diploma-user-service/proto/groupmemberservicepb"
	"github.com/TwiLightDM/diploma-user-service/proto/groupservicepb"
	"github.com/TwiLightDM/diploma-user-service/proto/userservicepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	user        userservicepb.UserServiceClient
	group       groupservicepb.GroupServiceClient
	groupMember groupmemberservicepb.GroupMemberServiceClient
	conn        *grpc.ClientConn
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
		user:        userservicepb.NewUserServiceClient(conn),
		group:       groupservicepb.NewGroupServiceClient(conn),
		groupMember: groupmemberservicepb.NewGroupMemberServiceClient(conn),
		conn:        conn,
	}
}

func (c *UserClient) Close() error {
	return c.conn.Close()
}
