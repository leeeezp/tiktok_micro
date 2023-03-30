package main

import (
	"context"
	user "tiktok_micro/kitex_gen/user"
	relation "tiktok_micro/kitex_gen/relation"
)

// UserRpcImpl implements the last service interface defined in the IDL.
type UserRpcImpl struct{}

// Register implements the UserRpcImpl interface.
func (s *UserRpcImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserRpcImpl interface.
func (s *UserRpcImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserRpcImpl interface.
func (s *UserRpcImpl) GetUser(ctx context.Context, req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	// TODO: Your code here...
	return
}

// FollowAction implements the UserRpcImpl interface.
func (s *UserRpcImpl) FollowAction(ctx context.Context, req *relation.FollowActionReq) (resp *relation.FollowActionResp, err error) {
	// TODO: Your code here...
	return
}

// FollowList implements the UserRpcImpl interface.
func (s *UserRpcImpl) FollowList(ctx context.Context, req *relation.FollowListReq) (resp *relation.FollowListResp, err error) {
	// TODO: Your code here...
	return
}

// FollowerList implements the UserRpcImpl interface.
func (s *UserRpcImpl) FollowerList(ctx context.Context, req *relation.FollowerListReq) (resp *relation.FollowerListResp, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the UserRpcImpl interface.
func (s *UserRpcImpl) FriendList(ctx context.Context, req *relation.FriendListReq) (resp *relation.FriendListResp, err error) {
	// TODO: Your code here...
	return
}

// MessageAction implements the UserRpcImpl interface.
func (s *UserRpcImpl) MessageAction(ctx context.Context, req *relation.MessageActionReq) (resp *relation.MessageActionResp, err error) {
	// TODO: Your code here...
	return
}

// MessageList implements the UserRpcImpl interface.
func (s *UserRpcImpl) MessageList(ctx context.Context, req *relation.MessageListReq) (resp *relation.MessageListResp, err error) {
	// TODO: Your code here...
	return
}
