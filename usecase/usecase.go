package usecase

import (
	"context"

	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
	"github.com/sapawarga/userpost-service/repository"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Post struct {
	repo   repository.PostI
	logger kitlog.Logger
}

func NewPost(repo repository.PostI, logger kitlog.Logger) *Post {
	return &Post{
		repo:   repo,
		logger: logger,
	}
}

func (p *Post) GetListPost(ctx context.Context, params *model.GetListRequest) (*model.UserPostWithMetadata, error) {
	logger := kitlog.With(p.logger, "method", "GetListPost")
	var limit, offset int64 = 10, 0

	if params.Page != nil && params.Limit != nil {
		limit = helper.GetInt64FromPointer(params.Limit)
		offset = (helper.GetInt64FromPointer(params.Page) - 1) * limit
	}

	req := &model.UserPostRequest{
		ActivityName: params.ActivityName,
		Username:     params.Username,
		Category:     params.Category,
		Status:       params.Status,
		Offset:       helper.SetPointerInt64(offset),
		Limit:        helper.SetPointerInt64(limit),
		SortBy:       params.SortBy,
		OrderBy:      params.ActivityName,
	}

	listData, err := p.repo.GetListPost(ctx, req)
	if err != nil {
		level.Error(logger).Log("error_get_userpost", err)
		return nil, err
	}

	userPosts := make([]*model.UserPostResponse, 0)

	for _, v := range listData {
		userPost := &model.UserPostResponse{
			ID:            v.ID,
			Title:         v.Title,
			Tag:           helper.SetPointerString(v.Tag.String),
			ImagePath:     v.ImagePath.String,
			Images:        v.Images.String,
			LikesCount:    v.LikesCount,
			CommentCounts: v.CommentCounts,
			Status:        v.Status,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		}
		if v.LastUserPostCommentID.Valid {
			comment, err := p.repo.GetLastComment(ctx, v.LastUserPostCommentID.Int64)
			if err != nil {
				level.Error(logger).Log("error_get_last_comment", err)
				return nil, err
			}
			userPost.LastUserPostCommentID = helper.SetPointerInt64(v.LastUserPostCommentID.Int64)
			userPost.LastComment = comment
		}
		if v.CreatedBy.Valid {
			user, err := p.repo.GetActor(ctx, v.CreatedBy.Int64)
			if err != nil {
				level.Error(logger).Log("error_get_actor", err)
				return nil, err
			}
			userPost.Actor = user
		}
		// TODO get isLiked by who is user login. Get the user login from context
		// isLiked, err := p.repo.GetIsLikedByUser()
	}

	total, err := p.repo.GetMetadataPost(ctx, req)
	if err != nil {
		level.Error(logger).Log("error_get_metadata", err)
		return nil, err
	}

	metadata := &model.Metadata{
		Page:      helper.GetInt64FromPointer(params.Page),
		TotalPage: helper.GetInt64FromPointer(total) / limit,
		Total:     helper.GetInt64FromPointer(total),
	}

	return &model.UserPostWithMetadata{
		Data:     userPosts,
		Metadata: metadata,
	}, nil
}
