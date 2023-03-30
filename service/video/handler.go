package main

import (
	"context"
	video "tiktok_micro/kitex_gen/video"
)

// VideoRpcImpl implements the last service interface defined in the IDL.
type VideoRpcImpl struct{}

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
