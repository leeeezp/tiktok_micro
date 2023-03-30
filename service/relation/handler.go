package main

import (
	"context"
	relation "tiktok_micro/kitex_gen/relation"
	video "tiktok_micro/kitex_gen/video"
)

// UserRpcImpl implements the last service interface defined in the IDL.
type UserRpcImpl struct{}

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

// GetPublishList implements the VideoRpcImpl interface.
func (s *VideoRpcImpl) GetPublishList(ctx context.Context, req *video.PublishListReq) (resp *video.PublishListResp, err error) {
	// TODO: Your code here...
	return
}

// GetFeed implements the VideoRpcImpl interface.
func (s *VideoRpcImpl) GetFeed(ctx context.Context, req *video.FeedReq) (resp *video.FeedResp, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the VideoRpcImpl interface.
func (s *VideoRpcImpl) CommentAction(ctx context.Context, req *video.CommentReq) (resp *video.CommentResp, err error) {
	// TODO: Your code here...
	return
}

// GetCommentList implements the VideoRpcImpl interface.
func (s *VideoRpcImpl) GetCommentList(ctx context.Context, req *video.CommentListReq) (resp *video.CommentListResp, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction implements the VideoRpcImpl interface.
func (s *VideoRpcImpl) FavoriteAction(ctx context.Context, req *video.FavoriteReq) (resp *video.FavoriteResp, err error) {
	// TODO: Your code here...
	return
}

// GetFavoriteList implements the VideoRpcImpl interface.
func (s *VideoRpcImpl) GetFavoriteList(ctx context.Context, req *video.FavoriteListReq) (resp *video.FavoriteListResp, err error) {
	// TODO: Your code here...
	return
}
