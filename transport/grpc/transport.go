package grpc

import (
	"context"

	"github.com/sapawarga/userpost-service/endpoint"
	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/usecase"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	transportUserPost "github.com/sapawarga/proto-file/userpost"
)

func MakeHandler(ctx context.Context, fs usecase.UsecaseI) transportUserPost.UserPostHandlerServer {
	userPostGetListHandler := kitgrpc.NewServer(
		endpoint.MakeGetListUserPost(ctx, fs),
		decodeGetListUserPost,
		encodeGetListUserPost,
	)

	userPostGetDetailHandler := kitgrpc.NewServer(
		endpoint.MakeGetDetailUserPost(ctx, fs),
		decodeByIDRequest,
		encodedUserPostDetail,
	)

	userPostCreateNewPostHandler := kitgrpc.NewServer(
		endpoint.MakeCreateNewPost(ctx, fs),
		decodeCreateNewPostRequest,
		encodeStatusResponse,
	)

	return &grpcServer{
		userPostGetListHandler,
		userPostGetDetailHandler,
		userPostCreateNewPostHandler,
	}
}

func decodeGetListUserPost(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportUserPost.GetListUserPostRequest)

	return &endpoint.GetListUserPostRequest{
		ActivityName: helper.SetPointerString(req.GetActivityName()),
		Username:     helper.SetPointerString(req.GetUsername()),
		Category:     helper.SetPointerString(req.GetCategory()),
		Status:       helper.SetPointerInt64(req.GetStatus()),
		Page:         helper.SetPointerInt64(req.GetPage()),
		Limit:        helper.SetPointerInt64(req.GetLimit()),
		SortBy:       helper.SetPointerString(req.GetSortBy()),
		OrderBy:      helper.SetPointerString(req.GetOrderBy()),
	}, nil
}

func encodeGetListUserPost(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.UserPostWithMetadata)
	data := resp.Data
	metadata := resp.Metadata

	resultData := make([]*transportUserPost.UserPost, 0)

	for _, v := range data {
		result := &transportUserPost.UserPost{
			Id:                    v.ID,
			Title:                 v.Title,
			Tag:                   helper.GetStringFromPointer(v.Tag),
			ImagePath:             v.ImagePath,
			Images:                v.Images,
			LastUserPostCommentId: helper.GetInt64FromPointer(v.LastUserPostCommentID),
			LikesCount:            v.LikesCount,
			CommentCounts:         v.CommentCounts,
			Status:                v.Status,
			CreatedAt:             v.CreatedAt.String(),
			UpdatedAt:             v.UpdatedAt.String(),
		}
		if v.Actor != nil {
			actor := &transportUserPost.Actor{
				Id:       v.Actor.ID,
				Name:     v.Actor.Name.String,
				PhotoUrl: v.Actor.PhotoURL.String,
				Role:     v.Actor.Role.Int64,
				Regency:  v.Actor.Regency,
				District: v.Actor.District,
				Village:  v.Actor.Village,
				Rw:       v.Actor.RW.String,
			}
			result.Actor = actor
		}
		if v.LastComment != nil {
			comment := &transportUserPost.Comment{
				Id:         v.LastComment.ID,
				UserPostId: v.LastComment.UserPostID,
				Comment:    v.LastComment.Text,
				CreatedAt:  v.LastComment.CreatedAt.String(),
				UpdatedAt:  v.LastComment.UpdatedAt.String(),
			}
			actorCreated := &transportUserPost.Actor{
				Id:       v.LastComment.CreatedBy.ID,
				Name:     v.LastComment.CreatedBy.Name.String,
				PhotoUrl: v.LastComment.CreatedBy.PhotoURL.String,
				Role:     v.LastComment.CreatedBy.Role.Int64,
				Regency:  v.LastComment.CreatedBy.Regency,
				District: v.LastComment.CreatedBy.District,
				Village:  v.LastComment.CreatedBy.Village,
				Rw:       v.LastComment.CreatedBy.RW.String,
			}
			actorUpdated := &transportUserPost.Actor{
				Id:       v.LastComment.UpdatedBy.ID,
				Name:     v.LastComment.UpdatedBy.Name.String,
				PhotoUrl: v.LastComment.UpdatedBy.PhotoURL.String,
				Role:     v.LastComment.UpdatedBy.Role.Int64,
				Regency:  v.LastComment.UpdatedBy.Regency,
				District: v.LastComment.UpdatedBy.District,
				Village:  v.LastComment.UpdatedBy.Village,
				Rw:       v.LastComment.UpdatedBy.RW.String,
			}
			comment.CreatedBy = actorCreated
			comment.UpdatedBy = actorUpdated
			result.LastComment = comment
		}

		resultData = append(resultData, result)
	}

	meta := &transportUserPost.Metadata{
		Page:      metadata.Page,
		Total:     metadata.Total,
		TotalPage: metadata.TotalPage,
	}

	return &transportUserPost.GetListUserPostResponse{
		Data:     resultData,
		Metadata: meta,
	}, nil
}

func decodeByIDRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportUserPost.ByID)

	return &endpoint.GetByID{
		ID: req.GetId(),
	}, nil
}

func encodedUserPostDetail(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.UserPostDetail)
	comment := resp.LastComment
	actor := resp.Actor

	lastComment := &transportUserPost.Comment{
		Id:         comment.ID,
		UserPostId: comment.UserPostID,
		Comment:    comment.Text,
		CreatedAt:  comment.CreatedAt.String(),
		UpdatedAt:  comment.UpdatedAt.String(),
	}

	lastCommentActorCreated := &transportUserPost.Actor{
		Id:       comment.CreatedBy.ID,
		Name:     comment.CreatedBy.Name.String,
		PhotoUrl: comment.CreatedBy.PhotoURL.String,
		Role:     comment.CreatedBy.Role.Int64,
		Regency:  comment.CreatedBy.Regency,
		District: comment.CreatedBy.District,
		Village:  comment.CreatedBy.Village,
		Rw:       comment.CreatedBy.RW.String,
	}

	lastCommentActorUpdated := &transportUserPost.Actor{
		Id:       comment.UpdatedBy.ID,
		Name:     comment.UpdatedBy.Name.String,
		PhotoUrl: comment.UpdatedBy.PhotoURL.String,
		Role:     comment.UpdatedBy.Role.Int64,
		Regency:  comment.UpdatedBy.Regency,
		District: comment.UpdatedBy.District,
		Village:  comment.UpdatedBy.Village,
		Rw:       comment.UpdatedBy.RW.String,
	}

	lastComment.CreatedBy = lastCommentActorCreated
	lastComment.UpdatedBy = lastCommentActorUpdated

	actorUserPost := &transportUserPost.Actor{
		Id:       actor.ID,
		Name:     actor.Name.String,
		PhotoUrl: actor.PhotoURL.String,
		Role:     actor.Role.Int64,
		Regency:  actor.Regency,
		District: actor.District,
		Village:  actor.Village,
		Rw:       actor.RW.String,
	}

	userDetail := &transportUserPost.UserPost{
		Id:                    resp.ID,
		Title:                 resp.Title,
		Tag:                   helper.GetStringFromPointer(resp.Tag),
		ImagePath:             resp.ImagePath,
		Images:                resp.Images,
		LastUserPostCommentId: helper.GetInt64FromPointer(resp.LastUserPostCommentID),
		LastComment:           lastComment,
		LikesCount:            resp.LikesCount,
		IsLiked:               resp.IsLiked,
		CommentCounts:         resp.CommentCounts,
		Status:                resp.Status,
		Actor:                 actorUserPost,
		CreatedAt:             resp.CreatedAt.String(),
		UpdatedAt:             resp.UpdatedAt.String(),
	}

	return userDetail, nil
}

func decodeCreateNewPostRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportUserPost.CreateNewPostRequest)

	images := make([]*endpoint.Image, 0)

	for _, v := range req.Images {
		image := &endpoint.Image{
			Path: v.GetPath(),
		}
		images = append(images, image)
	}

	return &endpoint.CreateNewPostRequest{
		Title:  helper.SetPointerString(req.GetTitle()),
		Images: images,
		Tags:   helper.SetPointerString(req.GetTags()),
		Status: helper.SetPointerInt64(req.GetStatus()),
	}, nil
}

func encodeStatusResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.StatusResponse)

	return &transportUserPost.StatusResponse{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}
