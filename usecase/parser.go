package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/sapawarga/userpost-service/config"
	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
)

func (p *Post) getDetailOfUserPost(ctx context.Context, post *model.PostResponse) (*model.UserPostResponse, error) {
	logger := kitlog.With(p.logger, "method", "getDetailOfUserPost")
	cfg, _ := config.NewConfig()
	images := make([]map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(post.Images.String), &images); err != nil {
		images = nil
	}
	for _, v := range images {
		v["path"] = fmt.Sprintf("%s/%s", cfg.AppStoragePublicURL, v["path"])
	}
	userPost := &model.UserPostResponse{
		ID:            post.ID,
		Title:         post.Title,
		Tag:           helper.SetPointerString(post.Tag.String),
		ImagePath:     fmt.Sprintf("%s/%s", cfg.AppStoragePublicURL, post.ImagePath.String),
		Images:        images,
		LikesCount:    post.LikesCount,
		CommentCounts: post.CommentCounts,
		Status:        post.Status,
		CreatedAt:     post.CreatedAt,
		UpdatedAt:     post.UpdatedAt,
	}
	if post.LastUserPostCommentID.Valid {
		comment, err := p.repoComment.GetLastComment(ctx, post.LastUserPostCommentID.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_comment", err)
			return nil, err
		}
		detailComment, err := p.getDetailComment(ctx, comment)
		if err != nil {
			level.Error(logger).Log("error_get_comment", err)
			return nil, err
		}
		userPost.LastUserPostCommentID = helper.SetPointerInt64(post.LastUserPostCommentID.Int64)
		userPost.LastComment = detailComment
	}
	if post.CreatedBy.Valid {
		user, err := p.repoPost.GetActor(ctx, post.CreatedBy.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_actor", err)
			return nil, err
		}
		userPost.Actor = p.parsingUserResponse(ctx, user)
	}

	return userPost, nil
}

func (p *Post) parsingUserResponse(ctx context.Context, user *model.UserResponse) *model.Actor {
	return &model.Actor{
		ID:        user.ID,
		Name:      user.Name.String,
		PhotoURL:  user.PhotoURL.String,
		Role:      user.Role.Int64,
		RoleLabel: model.RoleLabel[user.Role.Int64],
		Regency:   user.Regency.String,
		District:  user.District.String,
		Village:   user.Village.String,
		RW:        user.RW.String,
		Status:    user.Status,
	}
}

func (p *Post) getDetailComment(ctx context.Context, comment *model.CommentResponse) (*model.Comment, error) {
	logger := kitlog.With(p.logger, "method", "getDetailComment")
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
	commentResp.CreatedBy = p.parsingUserResponse(ctx, actorCreated)
	actorUpdated, err := p.repoPost.GetActor(ctx, comment.UpdatedBy)
	if err != nil {
		level.Error(logger).Log("error_get_actor_updated", err)
		return nil, err
	}
	commentResp.UpdatedBy = p.parsingUserResponse(ctx, actorUpdated)
	return commentResp, nil
}

func (p *Post) appendListUserPost(ctx context.Context, resp []*model.PostResponse) (userPosts []*model.UserPostResponse, err error) {
	// TODO: using actor TODO

	// actorID := ctx.Value(helper.ACTORKEY).(*model.ActorFromContext).Get("id").(int64)
	for _, v := range resp {
		userPost, err := p.getDetailOfUserPost(ctx, v)
		if err != nil {
			return nil, err
		}
		isLiked, err := p.repoPost.CheckIsExistLikeOnPostBy(ctx, &model.AddOrRemoveLikeOnPostRequest{
			UserPostID: v.ID,
			// ActorID:    actorID,
			TypeEntity: helper.TYPE_USERPOST,
		})
		if err != nil {
			return nil, err
		}
		if isLiked {
			userPost.IsLiked = isLiked
		}
		userPosts = append(userPosts, userPost)
	}
	return userPosts, nil
}
