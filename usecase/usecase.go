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
	repoPost    repository.PostI
	repoComment repository.CommentI
	logger      kitlog.Logger
}

func NewPost(repoPost repository.PostI, repoComment repository.CommentI, logger kitlog.Logger) *Post {
	return &Post{
		repoPost:    repoPost,
		repoComment: repoComment,
		logger:      logger,
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

	listData, err := p.repoPost.GetListPost(ctx, req)
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
			comment, err := p.repoComment.GetLastComment(ctx, v.LastUserPostCommentID.Int64)
			if err != nil {
				level.Error(logger).Log("error_get_last_comment", err)
				return nil, err
			}
			userPost.LastUserPostCommentID = helper.SetPointerInt64(v.LastUserPostCommentID.Int64)
			commentResp := &model.Comment{
				ID:         comment.ID,
				UserPostID: comment.UserPostID,
				Text:       comment.Comment,
				CreatedAt:  comment.CreatedAt,
				UpdatedAt:  comment.UpdatedAt,
			}
			actorCreated, err := p.repoPost.GetActor(ctx, comment.CreatedBy)
			if err != nil {
				level.Error(logger).Log("error_get_actor_created", err)
				return nil, err
			}
			commentResp.CreatedBy = actorCreated
			actorUpdated, err := p.repoPost.GetActor(ctx, comment.UpdatedBy)
			if err != nil {
				level.Error(logger).Log("error_get_actor_updated", err)
				return nil, err
			}
			commentResp.UpdatedBy = actorUpdated
			userPost.LastComment = commentResp
		}
		if v.CreatedBy.Valid {
			user, err := p.repoPost.GetActor(ctx, v.CreatedBy.Int64)
			if err != nil {
				level.Error(logger).Log("error_get_actor", err)
				return nil, err
			}
			userPost.Actor = user
		}
		// TODO get isLiked by who is user login. Get the user login from context
		// isLiked, err := p.repo.GetIsLikedByUser()
	}

	total, err := p.repoPost.GetMetadataPost(ctx, req)
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

func (p *Post) GetDetailPost(ctx context.Context, id int64) (*model.UserPostResponse, error) {
	logger := kitlog.With(p.logger, "method", "GetDetailPost")
	resp, err := p.repoPost.GetDetailPost(ctx, id)
	if err != nil {
		level.Error(logger).Log("error_get_detail", err)
		return nil, err
	}

	userPost := &model.UserPostResponse{
		ID:            resp.ID,
		Title:         resp.Title,
		Tag:           helper.SetPointerString(resp.Tag.String),
		ImagePath:     resp.ImagePath.String,
		Images:        resp.Images.String,
		LikesCount:    resp.LikesCount,
		CommentCounts: resp.CommentCounts,
		Status:        resp.Status,
		CreatedAt:     resp.CreatedAt,
		UpdatedAt:     resp.UpdatedAt,
	}
	if resp.LastUserPostCommentID.Valid {
		comment, err := p.repoComment.GetLastComment(ctx, resp.LastUserPostCommentID.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_last_comment", err)
			return nil, err
		}
		userPost.LastUserPostCommentID = helper.SetPointerInt64(resp.LastUserPostCommentID.Int64)
		commentResp := &model.Comment{
			ID:         comment.ID,
			UserPostID: comment.UserPostID,
			Text:       comment.Comment,
			CreatedAt:  comment.CreatedAt,
			UpdatedAt:  comment.UpdatedAt,
		}
		actorCreated, err := p.repoPost.GetActor(ctx, comment.CreatedBy)
		if err != nil {
			level.Error(logger).Log("error_get_actor_created", err)
			return nil, err
		}
		commentResp.CreatedBy = actorCreated
		actorUpdated, err := p.repoPost.GetActor(ctx, comment.UpdatedBy)
		if err != nil {
			level.Error(logger).Log("error_get_actor_updated", err)
			return nil, err
		}
		commentResp.UpdatedBy = actorUpdated
	}
	if resp.CreatedBy.Valid {
		user, err := p.repoPost.GetActor(ctx, resp.CreatedBy.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_actor", err)
			return nil, err
		}
		userPost.Actor = user
	}
	// TODO get isLiked by who is user login. Get the user login from context
	// isLiked, err := p.repo.GetIsLikedByUser()
	return userPost, nil
}
