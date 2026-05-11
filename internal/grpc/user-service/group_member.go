package user_service

import (
	"context"

	"github.com/TwiLightDM/diploma-user-service/proto/groupmemberservicepb"
)

func (c *UserClient) CreateGroupMember(ctx context.Context, email, groupId string) (*groupmemberservicepb.CreateGroupMemberResponse, error) {
	return c.groupMember.CreateGroupMember(ctx, &groupmemberservicepb.CreateGroupMemberRequest{
		Email:   email,
		GroupId: groupId,
	})
}

func (c *UserClient) ReadAllGroupMembersByUserId(ctx context.Context, userId string) (*groupmemberservicepb.ReadAllGroupMembersByUserIdResponse, error) {
	return c.groupMember.ReadAllGroupMembersByUserId(ctx, &groupmemberservicepb.ReadAllGroupMembersByUserIdRequest{
		UserId: userId,
	})
}

func (c *UserClient) ReadAllGroupMembersByGroupId(ctx context.Context, groupId string) (*groupmemberservicepb.ReadAllGroupMembersByGroupIdResponse, error) {
	return c.groupMember.ReadAllGroupMembersByGroupId(ctx, &groupmemberservicepb.ReadAllGroupMembersByGroupIdRequest{
		GroupId: groupId,
	})
}

func (c *UserClient) DeleteGroupMember(ctx context.Context, id string) (*groupmemberservicepb.DeleteGroupMemberResponse, error) {
	return c.groupMember.DeleteGroupMember(ctx, &groupmemberservicepb.DeleteGroupMemberRequest{
		Id: id,
	})
}
