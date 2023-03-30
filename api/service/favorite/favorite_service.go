/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"context"
	db2 "tiktok_micro/api/dal/db"
	"tiktok_micro/api/model/basic/feed"
	"tiktok_micro/api/model/common"
	"tiktok_micro/api/model/interact/favorite"
	feed_service "tiktok_micro/api/service/feed"

	"github.com/cloudwego/hertz/pkg/app"

	"tiktok_micro/api/pkg/constants"
	"tiktok_micro/api/pkg/errno"
)

type FavoriteService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewFavoriteService create favorite service
func NewFavoriteService(ctx context.Context, c *app.RequestContext) *FavoriteService {
	return &FavoriteService{ctx: ctx, c: c}
}

// FavoriteAction like a video and return result
func (r *FavoriteService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) (flag bool, err error) {
	_, err = db2.CheckVideoExistById(req.VideoId)
	if err != nil {
		return false, err
	}
	if req.ActionType != constants.FavoriteActionType && req.ActionType != constants.UnFavoriteActionType {
		return false, errno.ParamErr
	}
	current_user_id, _ := r.c.Get("current_user_id")

	new_favorite_relation := &db2.Favorites{
		UserId:  current_user_id.(int64),
		VideoId: req.VideoId,
	}
	favorite_exist, _ := db2.CheckFavoriteRelationExist(new_favorite_relation)
	if req.ActionType == constants.FavoriteActionType {
		if favorite_exist {
			return false, errno.FavoriteRelationAlreadyExistErr
		}
		flag, err = db2.AddNewFavorite(new_favorite_relation)
	} else {
		if !favorite_exist {
			return false, errno.FavoriteRelationNotExistErr
		}
		flag, err = db2.DeleteFavorite(new_favorite_relation)
	}
	return flag, err
}

// GetFavoriteList query favorite list by the user id in the request
func (r *FavoriteService) GetFavoriteList(req *favorite.DouyinFavoriteListRequest) (favoritelist []*common.Video, err error) {
	query_user_id := req.UserId
	_, err = db2.CheckUserExistById(query_user_id)

	if err != nil {
		return nil, err
	}
	current_user_id, _ := r.c.Get("current_user_id")

	video_id_list, err := db2.GetFavoriteIdList(query_user_id)

	dbVideos, err := db2.GetVideoListByVideoIDList(video_id_list)
	var videos []*feed.Video
	f := feed_service.NewFeedService(r.ctx, r.c)
	err = f.CopyVideos(&videos, &dbVideos, current_user_id.(int64))
	for _, item := range videos {
		video := &common.Video{
			Id: item.Id,
			Author: &common.User{
				Id:              item.Author.Id,
				Name:            item.Author.Name,
				FollowCount:     item.Author.FollowCount,
				FollowerCount:   item.Author.FollowerCount,
				Avatar:          item.Author.Avatar,
				BackgroundImage: item.Author.BackgroundImage,
				Signature:       item.Author.Signature,
				TotalFavorited:  item.Author.TotalFavorited,
				WorkCount:       item.Author.WorkCount,
			},
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
		}
		favoritelist = append(favoritelist, video)
	}
	return favoritelist, err
}
