package grpc

import (
	"context"
	"encoding/json"

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

	userPostGetListByMeHandler := kitgrpc.NewServer(
		endpoint.MakeGetListUserPostByMe(ctx, fs),
		decodeGetListUserPost,
		encodeGetListUserPost,
	)

	userPostLikeDislikeHandler := kitgrpc.NewServer(
		endpoint.MakeLikeOrDislikePost(ctx, fs),
		decodeByIDRequest,
		encodeStatusResponse,
	)

	return &grpcServer{
		userPostGetListHandler,
		userPostGetDetailHandler,
		userPostCreateNewPostHandler,
		userPostUpdateHandler,
		userPostGetCommentsHandler,
		userPostCreateCommentHandler,
		userPostGetListByMeHandler,
		userPostLikeDislikeHandler,
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
		images, _ := json.Marshal(v.Images)
		result := &transportUserPost.UserPost{
			Id:                    v.ID,
			Title:                 v.Title,
			Tag:                   helper.GetStringFromPointer(v.Tag),
			ImagePath:             v.ImagePath,
			Images:                string(images),
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

func appendDetailUserPost(ctx context.Context, r *model.UserPostResponse, data *transportUserPost.UserPost) *transportUserPost.UserPost {
	if r.Actor != nil {
		actor := encodeActor(ctx, r.Actor)
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
		actorCreated := encodeActor(ctx, r.LastComment.CreatedBy)
		actorUpdated := encodeActor(ctx, r.LastComment.UpdatedBy)
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

	lastComment := &transportUserPost.Comment{
		Id:         comment.ID,
		UserPostId: comment.UserPostID,
		Comment:    comment.Text,
		CreatedAt:  comment.CreatedAt.String(),
		UpdatedAt:  comment.UpdatedAt.String(),
	}

	lastCommentActorCreated := encodeActor(ctx, comment.CreatedBy)
	lastCommentActorUpdated := encodeActor(ctx, comment.UpdatedBy)

	lastComment.CreatedBy = lastCommentActorCreated
	lastComment.UpdatedBy = lastCommentActorUpdated

	actorUserPost := encodeActor(ctx, resp.Actor)

	images, _ := json.Marshal(resp.Images)
	userDetail := &transportUserPost.UserPost{
		Id:                    resp.ID,
		Title:                 resp.Title,
		Tag:                   helper.GetStringFromPointer(resp.Tag),
		ImagePath:             resp.ImagePath,
		Images:                string(images),
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

func encodeActor(ctx context.Context, r interface{}) *transportUserPost.Actor {
	actorResp := r.(*model.UserResponse)
	return &transportUserPost.Actor{
		Id:       actorResp.ID,
		Name:     actorResp.Name.String,
		PhotoUrl: actorResp.PhotoURL.String,
		Role:     actorResp.Role.Int64,
		Regency:  actorResp.Regency.String,
		District: actorResp.District.String,
		Village:  actorResp.Village.String,
		Rw:       actorResp.RW.String,
	}
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

	return &endpoint.CreateCommentRequest{
		UserPostID: req.GetId(),
		Status:     helper.SetPointerInt64(req.GetStatus()),
		Text:       req.GetTitle(),
	}, nil
}

func encodeGetCommentsByIDResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.CommentsResponse)

	response := make([]*transportUserPost.Comment, 0)
	for _, v := range resp.Data {
		created := encodeActor(ctx, v.CreatedBy)
		updated := encodeActor(ctx, v.UpdatedBy)
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
		Text:       req.GetComment(),
		Status:     helper.SetPointerInt64(req.GetStatus()),
	}, nil
}
