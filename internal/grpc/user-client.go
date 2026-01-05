package grpc

import (
	"context"
	userpb "github.com/TwiLightDM/diploma-user-service/proto/userservicepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type UserClient struct {
	client userpb.UserServiceClient
	conn   *grpc.ClientConn
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
		client: userpb.NewUserServiceClient(conn),
		conn:   conn,
	}
}

func (u *UserClient) Close() error {
	return u.conn.Close()
}

func (u *UserClient) Login(ctx context.Context, email, password string) (*userpb.LoginResponse, error) {
	return u.client.Login(ctx, &userpb.LoginRequest{
		Email:    email,
		Password: password,
	})
}

func (u *UserClient) SignUp(ctx context.Context, fullName, role, email, password string) (*userpb.SignUpResponse, error) {
	return u.client.SignUp(ctx, &userpb.SignUpRequest{
		FullName: fullName,
		Role:     role,
		Email:    email,
		Password: password,
	})
}

func (u *UserClient) ReadUser(ctx context.Context, id string) (*userpb.ReadUserResponse, error) {
	return u.client.ReadUser(ctx, &userpb.ReadUserRequest{
		Id: id,
	})
}

func (u *UserClient) UpdateUser(ctx context.Context, id, fullName, email string) (*userpb.UpdateUserResponse, error) {
	return u.client.UpdateUser(ctx, &userpb.UpdateUserRequest{
		Id:       id,
		FullName: fullName,
		Email:    email,
	})
}

func (u *UserClient) ChangePassword(ctx context.Context, id, password string) (*userpb.ChangePasswordResponse, error) {
	return u.client.ChangePassword(ctx, &userpb.ChangePasswordRequest{
		Id:       id,
		Password: password,
	})
}
