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
		userPost, err := p.getDetailOfUserPost(ctx, v)
		if err != nil {
			level.Error(logger).Log("error_get_detail_user_post", err)
			return nil, err
		}
		userPosts = append(userPosts, userPost)
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

	userPost, err := p.getDetailOfUserPost(ctx, resp)
	if err != nil {
		level.Error(logger).Log("error_get_detail_user_post", err)
		return nil, err
	}

	return userPost, nil
}

func (p *Post) CreateNewPost(ctx context.Context, requestBody *model.CreateNewPostRequest) error {
	// TODO: add checker tags
	logger := kitlog.With(p.logger, "method", "CreateNewPost")
	if err := p.repoPost.InsertPost(ctx, requestBody); err != nil {
		level.Error(logger).Log("error_create_post", err)
		return err
	}

	return nil
}

func (p *Post) getDetailOfUserPost(ctx context.Context, post *model.PostResponse) (*model.UserPostResponse, error) {
	logger := kitlog.With(p.logger, "method", "getDetailOfUserPost")
	userPost := &model.UserPostResponse{
		ID:            post.ID,
		Title:         post.Title,
		Tag:           helper.SetPointerString(post.Tag.String),
		ImagePath:     post.ImagePath.String,
		Images:        post.Images.String,
		LikesCount:    post.LikesCount,
		CommentCounts: post.CommentCounts,
		Status:        post.Status,
		CreatedAt:     post.CreatedAt,
		UpdatedAt:     post.UpdatedAt,
	}
	if post.LastUserPostCommentID.Valid {
		comment, err := p.getLastComment(ctx, post.LastUserPostCommentID.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_comment", err)
			return nil, err
		}
		userPost.LastUserPostCommentID = helper.SetPointerInt64(post.LastUserPostCommentID.Int64)
		userPost.LastComment = comment
	}
	if post.CreatedBy.Valid {
		user, err := p.repoPost.GetActor(ctx, post.CreatedBy.Int64)
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

func (p *Post) getLastComment(ctx context.Context, id int64) (*model.Comment, error) {
	logger := kitlog.With(p.logger, "method", "getLastComment")
	comment, err := p.repoComment.GetLastComment(ctx, id)
	if err != nil {
		level.Error(logger).Log("error_get_last_comment", err)
		return nil, err
	}
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
	return commentResp, nil
}
