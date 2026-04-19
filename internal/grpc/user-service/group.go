package user_service

import (
	"context"

	"github.com/TwiLightDM/diploma-user-service/proto/groupservicepb"
)

func (c *UserClient) CreateGroup(ctx context.Context, title, description, ownerId string) (*groupservicepb.CreateGroupResponse, error) {
	return c.group.CreateGroup(ctx, &groupservicepb.CreateGroupRequest{
		Title:       title,
		Description: description,
		OwnerId:     ownerId,
	})
}

func (c *UserClient) ReadGroup(ctx context.Context, id string) (*groupservicepb.ReadGroupResponse, error) {
	return c.group.ReadGroup(ctx, &groupservicepb.ReadGroupRequest{
		Id: id,
	})
}

func (c *UserClient) ReadAllGroupsByOwnerId(ctx context.Context, ownerId string) (*groupservicepb.ReadAllGroupsByOwnerIdResponse, error) {
	return c.group.ReadAllGroupsByOwnerId(ctx, &groupservicepb.ReadAllGroupsByOwnerIdRequest{
		OwnerId: ownerId,
	})
}

func (c *UserClient) UpdateGroup(ctx context.Context, id, title, description string) (*groupservicepb.UpdateGroupResponse, error) {
	return c.group.UpdateGroup(ctx, &groupservicepb.UpdateGroupRequest{
		Id:          id,
		Title:       title,
		Description: description,
	})
}

func (c *UserClient) DeleteGroup(ctx context.Context, id string) (*groupservicepb.DeleteGroupResponse, error) {
	return c.group.DeleteGroup(ctx, &groupservicepb.DeleteGroupRequest{
		Id: id,
	})
}
