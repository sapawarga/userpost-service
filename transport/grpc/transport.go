package grpc

import (
	"context"

	"github.com/sapawarga/userpost-service/endpoint"
	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
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

	userPostUpdateHandler := kitgrpc.NewServer(
		endpoint.MakeUpdateStatusOrTitle(ctx, fs),
		decodeUpdateUserPost,
		encodeStatusResponse,
	)

	userPostGetCommentsHandler := kitgrpc.NewServer(
		endpoint.MakeGetCommentsByID(ctx, fs),
		decodeByIDRequest,
		encodeGetCommentsByIDResponse,
	)

	userPostCreateCommentHandler := kitgrpc.NewServer(
		endpoint.MakeCreateComment(ctx, fs),
		decodeCreateCommentRequest,
		encodeStatusResponse,
	)

	return &grpcServer{
		userPostGetListHandler,
		userPostGetDetailHandler,
		userPostCreateNewPostHandler,
		userPostUpdateHandler,
		userPostGetCommentsHandler,
		userPostCreateCommentHandler,
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
		result = appendDetailUserPost(ctx, v, result)
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

func appendDetailUserPost(_ context.Context, r *model.UserPostResponse, data *transportUserPost.UserPost) *transportUserPost.UserPost {
	if r.Actor != nil {
		actor := &transportUserPost.Actor{
			Id:       r.Actor.ID,
			Name:     r.Actor.Name.String,
			PhotoUrl: r.Actor.PhotoURL.String,
			Role:     r.Actor.Role.Int64,
			Regency:  r.Actor.Regency,
			District: r.Actor.District,
			Village:  r.Actor.Village,
			Rw:       r.Actor.RW.String,
		}
		data.Actor = actor
	}
	if r.LastComment != nil {
		comment := &transportUserPost.Comment{
			Id:         r.LastComment.ID,
			UserPostId: r.LastComment.UserPostID,
			Comment:    r.LastComment.Text,
			CreatedAt:  r.LastComment.CreatedAt.String(),
			UpdatedAt:  r.LastComment.UpdatedAt.String(),
		}
		actorCreated := &transportUserPost.Actor{
			Id:       r.LastComment.CreatedBy.ID,
			Name:     r.LastComment.CreatedBy.Name.String,
			PhotoUrl: r.LastComment.CreatedBy.PhotoURL.String,
			Role:     r.LastComment.CreatedBy.Role.Int64,
			Regency:  r.LastComment.CreatedBy.Regency,
			District: r.LastComment.CreatedBy.District,
			Village:  r.LastComment.CreatedBy.Village,
			Rw:       r.LastComment.CreatedBy.RW.String,
		}
		actorUpdated := &transportUserPost.Actor{
			Id:       r.LastComment.UpdatedBy.ID,
			Name:     r.LastComment.UpdatedBy.Name.String,
			PhotoUrl: r.LastComment.UpdatedBy.PhotoURL.String,
			Role:     r.LastComment.UpdatedBy.Role.Int64,
			Regency:  r.LastComment.UpdatedBy.Regency,
			District: r.LastComment.UpdatedBy.District,
			Village:  r.LastComment.UpdatedBy.Village,
			Rw:       r.LastComment.UpdatedBy.RW.String,
		}
		comment.CreatedBy = actorCreated
		comment.UpdatedBy = actorUpdated
		data.LastComment = comment
	}
	return data
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
	lastComment := &transportUserPost.Comment{comment.ID, comment.UserPostID, comment.Text,
		comment.CreatedAt.String(), comment.UpdatedAt.String(),
	}
	lastCommentActorCreated := &transportUserPost.Actor{comment.CreatedBy.ID,
		comment.CreatedBy.Name.String, comment.CreatedBy.PhotoURL.String, comment.CreatedBy.Role.Int64,
		comment.CreatedBy.Regency, comment.CreatedBy.District, comment.CreatedBy.Village, comment.CreatedBy.RW.String,
	}
	lastCommentActorUpdated := &transportUserPost.Actor{comment.UpdatedBy.ID, comment.UpdatedBy.Name.String,
		comment.UpdatedBy.PhotoURL.String, comment.UpdatedBy.Role.Int64, comment.UpdatedBy.Regency, comment.UpdatedBy.District,
		comment.UpdatedBy.Village, comment.UpdatedBy.RW.String,
	}
	lastComment.CreatedBy = lastCommentActorCreated
	lastComment.UpdatedBy = lastCommentActorUpdated
	actorUserPost := &transportUserPost.Actor{actor.ID, actor.Name.String, actor.PhotoURL.String,
		actor.Role.Int64, actor.Regency, actor.District, actor.Village, actor.RW.String,
	}

	return &transportUserPost.UserPost{
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
	}, nil
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

func decodeUpdateUserPost(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportUserPost.UpdateUserPostRequest)

	return &endpoint.UpdateStatusOrTitle{
		ID:     req.GetId(),
		Status: helper.SetPointerInt64(req.GetStatus()),
		Title:  helper.SetPointerString(req.GetTitle()),
	}, nil
}

func encodeGetCommentsByIDResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.CommentsResponse)

	response := make([]*transportUserPost.Comment, 0)
	for _, v := range resp.Data {
		created := &transportUserPost.Actor{
			Id:       v.CreatedBy.ID,
			Name:     v.CreatedBy.Name.String,
			PhotoUrl: v.CreatedBy.PhotoURL.String,
			Role:     v.CreatedBy.Role.Int64,
			Regency:  v.CreatedBy.Regency,
			District: v.CreatedBy.District,
			Village:  v.CreatedBy.Village,
			Rw:       v.CreatedBy.RW.String,
		}
		updated := &transportUserPost.Actor{
			Id:       v.UpdatedBy.ID,
			Name:     v.UpdatedBy.Name.String,
			PhotoUrl: v.UpdatedBy.PhotoURL.String,
			Role:     v.UpdatedBy.Role.Int64,
			Regency:  v.UpdatedBy.Regency,
			District: v.UpdatedBy.District,
			Village:  v.UpdatedBy.Village,
			Rw:       v.UpdatedBy.RW.String,
		}
		comment := &transportUserPost.Comment{
			Id:         v.ID,
			UserPostId: v.UserPostID,
			Comment:    v.Text,
			CreatedAt:  v.CreatedAt.String(),
			UpdatedAt:  v.UpdatedAt.String(),
			CreatedBy:  created,
			UpdatedBy:  updated,
		}
		response = append(response, comment)
	}

	return &transportUserPost.CommentsResponse{
		Comments: response,
	}, nil
}

func decodeCreateCommentRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportUserPost.CreateCommentRequest)

	return &endpoint.CreateCommentRequest{
		UserPostID: req.GetUserPostId(),
		Comment:    req.GetComment(),
		Status:     helper.SetPointerInt64(req.GetStatus()),
	}, nil
}
