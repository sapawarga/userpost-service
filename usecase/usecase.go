package usecase

import (
	"context"
	"errors"

	"github.com/sapawarga/userpost-service/lib/constant"
	"github.com/sapawarga/userpost-service/lib/convert"
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
	var limit, offset = constant.DEFAULT_LIMIT, constant.DEFAULT_OFFSET
	if params.Page != nil && params.Limit != nil {
		limit = convert.GetInt64FromPointer(params.Limit)
		offset = (convert.GetInt64FromPointer(params.Page) - 1) * limit
	}

	req := &model.UserPostRequest{
		ActivityName: params.ActivityName,
		Username:     params.Username,
		Category:     params.Category,
		Status:       params.Status,
		Offset:       convert.SetPointerInt64(offset),
		Limit:        convert.SetPointerInt64(limit),
		SortBy:       params.SortBy,
		OrderBy:      params.OrderBy,
		Search:       params.Search,
	}

	listData, err := p.repoPost.GetListPost(ctx, req)
	if err != nil {
		level.Error(logger).Log("error_get_userpost", err)
		return nil, err
	}

	userPosts, err := p.appendListUserPost(ctx, listData)
	if err != nil {
		level.Error(logger).Log("error_append_list", err)
		return nil, err
	}
	total, err := p.repoPost.GetMetadataPost(ctx, req)
	if err != nil {
		level.Error(logger).Log("error_get_metadata", err)
		return nil, err
	}

	metadata := &model.Metadata{
		Page:  convert.GetInt64FromPointer(params.Page),
		Total: convert.GetInt64FromPointer(total),
	}

	return &model.UserPostWithMetadata{Data: userPosts, Metadata: metadata}, nil
}

func (p *Post) GetListPostByMe(ctx context.Context, params *model.GetListRequest) (*model.UserPostWithMetadata, error) {
	logger := kitlog.With(p.logger, "method", "GetListPostByMe")
	var limit, offset int64 = constant.DEFAULT_LIMIT, constant.DEFAULT_OFFSET

	if params.Page != nil && params.Limit != nil {
		limit = convert.GetInt64FromPointer(params.Limit)
		offset = (convert.GetInt64FromPointer(params.Page) - 1) * limit
	}

	req := &model.UserPostByMeRequest{
		// ActorID: ctx.Value(constant.ACTORKEY).(*model.ActorFromContext).Get("id").(int64),
		UserPostRequest: &model.UserPostRequest{
			ActivityName: params.ActivityName,
			Username:     params.Username,
			Category:     params.Category,
			Status:       params.Status,
			Offset:       convert.SetPointerInt64(offset),
			Limit:        convert.SetPointerInt64(limit),
			SortBy:       params.SortBy,
			OrderBy:      params.OrderBy,
			Search:       params.Search},
	}
	resp, err := p.repoPost.GetListPostByMe(ctx, req)
	if err != nil {
		level.Error(logger).Log("error_get_list_post_by_me", err)
		return nil, err
	}

	userPosts, err := p.appendListUserPost(ctx, resp)
	if err != nil {
		level.Error(logger).Log("error_append_list", err)
		return nil, err
	}

	total, err := p.repoPost.GetMetadataPostByMe(ctx, req)
	if err != nil {
		level.Error(logger).Log("error_get_metadata", err)
		return nil, err
	}
	return &model.UserPostWithMetadata{
		Data: userPosts,
		Metadata: &model.Metadata{
			Page:  convert.GetInt64FromPointer(params.Page),
			Total: convert.GetInt64FromPointer(total),
		},
	}, nil
}

func (p *Post) GetDetailPost(ctx context.Context, id int64) (*model.UserPostResponse, error) {
	// TODO: add actor
	logger := kitlog.With(p.logger, "method", "GetDetailPost")
	// actor := ctx.Value(constant.ACTORKEY).(*model.ActorFromContext).Data
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

	isLiked, err := p.repoPost.CheckIsExistLikeOnPostBy(ctx, &model.AddOrRemoveLikeOnPostRequest{
		UserPostID: id,
		// ActorID:    actor["id"].(int64),
		TypeEntity: constant.TYPE_USERPOST,
	})
	if err != nil {
		level.Error(logger).Log("error_check_isliked", err)
		return nil, err
	}

	if isLiked {
		userPost.IsLiked = isLiked
	}

	return userPost, nil
}

func (p *Post) CreateNewPost(ctx context.Context, requestBody *model.CreateNewPostRequest) error {
	// TODO: add checker tags, Add Actor
	logger := kitlog.With(p.logger, "method", "CreateNewPost")
	// actor := ctx.Value(constant.ACTORKEY).(*model.ActorFromContext).Data
	if err := p.repoPost.InsertPost(ctx, &model.CreateNewPostRequestRepository{
		Title:        requestBody.Title,
		ImagePathURL: requestBody.ImagePathURL,
		Images:       requestBody.Images,
		Tags:         requestBody.Tags,
		Status:       requestBody.Status,
		// ActorID:      actor["id"].(int64),
		ActorID: 1, // for now use default as admin
	}); err != nil {
		level.Error(logger).Log("error_create_post", err)
		return err
	}

	level.Info(logger).Log("msg", "success_create_new_post")
	return nil
}

func (p *Post) UpdateTitleOrStatus(ctx context.Context, requestBody *model.UpdatePostRequest) error {
	logger := kitlog.With(p.logger, "method", "UpdateTitleOrStatus")
	_, err := p.repoPost.GetDetailPost(ctx, requestBody.ID)
	if err != nil {
		level.Error(logger).Log("error_get_detail", err)
		return err
	}
	if err = p.repoPost.UpdateDetailOfUserPost(ctx, requestBody); err != nil {
		level.Error(logger).Log("error_update", err)
		return err
	}

	level.Info(logger).Log("msg", "success_update_title_or_status")
	return nil
}

func (p *Post) GetCommentsByPostID(ctx context.Context, req *model.GetCommentRequest) (*model.CommentWithMetadata, error) {
	logger := kitlog.With(p.logger, "method", "GetCommentsByPostID")
	var limit, offset = constant.DEFAULT_LIMIT, constant.DEFAULT_OFFSET

	resp, err := p.repoComment.GetCommentsByPostID(ctx, &model.GetComment{
		ID:     req.ID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		level.Error(logger).Log("error_get_comments_by_post_id", err)
		return nil, err
	}

	totalComment, err := p.repoComment.GetTotalComments(ctx, req.ID)
	if err != nil {
		level.Error(logger).Log("error_get_metadata", err)
		return nil, err
	}
	comments := make([]*model.Comment, 0)

	if len(resp) > 0 {
		for _, v := range resp {
			detailComment, err := p.getDetailComment(ctx, v)
			if err != nil {
				level.Error(logger).Log("error_get_detail_comment", err)
				return nil, err
			}
			comment := &model.Comment{
				ID:         detailComment.ID,
				UserPostID: detailComment.UserPostID,
				Text:       detailComment.Text,
				User:       detailComment.User,
				CreatedAt:  detailComment.CreatedAt,
				UpdatedAt:  detailComment.UpdatedAt,
				CreatedBy:  detailComment.CreatedBy,
				UpdatedBy:  detailComment.UpdatedBy,
			}
			comments = append(comments, comment)
		}
	}
	meta := &model.Metadata{
		Page:  req.Page,
		Total: convert.GetInt64FromPointer(totalComment),
	}

	return &model.CommentWithMetadata{
		Data:     comments,
		Metadata: meta,
	}, nil
}

func (p *Post) CreateCommentOnPost(ctx context.Context, req *model.CreateCommentRequest) error {
	// TODO: implement authorization and authenticationn
	logger := kitlog.With(p.logger, "method", "CreateCommentOnPost")
	// actor := ctx.Value(constant.ACTORKEY).(*model.ActorFromContext).Data
	id, err := p.repoComment.Create(ctx, &model.CreateCommentRequestRepository{
		UserPostID: req.UserPostID,
		Text:       req.Text,
		Status:     req.Status,
		// ActorID:    actor["id"].(int64),
		ActorID: 1, // TODO: actor not existed yet using admin as default
	})
	if err != nil {
		level.Error(logger).Log("error_create_comment", err)
		return err
	}

	if err := p.repoPost.UpdateDetailOfUserPost(ctx, &model.UpdatePostRequest{
		ID:            req.UserPostID,
		LastCommentID: convert.SetPointerInt64(id),
	}); err != nil {
		level.Error(logger).Log("error_update", err)
		return err
	}

	return nil
}

func (p *Post) LikeOrDislikePost(ctx context.Context, id int64) error {
	// TODO: context actor
	logger := kitlog.With(p.logger, "method", "LikeOrDislikePost")
	// actorID := ctx.Value(constant.ACTORKEY).(*model.ActorFromContext).Get("id").(int64)
	var err error

	request := &model.AddOrRemoveLikeOnPostRequest{
		UserPostID: id,
		// ActorID:    actorID,
		TypeEntity: constant.TYPE_USERPOST,
	}
	isExist, err := p.repoPost.CheckIsExistLikeOnPostBy(ctx, request)
	if err != nil {
		level.Error(logger).Log("error_check_is_liked", err)
		return err
	}

	if isExist {
		err = p.repoPost.AddLikeOnPost(ctx, request)
	} else {
		err = p.repoPost.RemoveLikeOnPost(ctx, request)
	}

	if err != nil {
		level.Error(logger).Log("error_add_or_remove_like", err)
		return err
	}

	return nil
}

func (p *Post) CheckHealthReadiness(ctx context.Context) error {
	logger := kitlog.With(p.logger, "method", "CheckHealthReadiness")
	if err := p.repoPost.CheckHealthReadiness(ctx); err != nil {
		level.Error(logger).Log("error", errors.New("service_not_ready"))
		return errors.New("service_not_ready")
	}
	return nil
}
