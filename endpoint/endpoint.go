package endpoint

import (
	"context"
	"encoding/json"
	"math"

	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
	"github.com/sapawarga/userpost-service/usecase"
)

func MakeGetListUserPost(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetListUserPostRequest)
		orderBy := helper.GetStringFromPointer(req.OrderBy)
		if req.OrderBy != nil && !isOrderValid(orderBy) {
			return nil, errors.New("order_must_between_ASC_DESC")
		}
		// TODO: for get metadata from headers grpc needs to update when using authorization
		resp, err := usecase.GetListPost(ctx, &model.GetListRequest{
			ActivityName: req.ActivityName,
			Username:     req.Username,
			Category:     req.Category,
			Status:       req.Status,
			Page:         req.Page,
			Limit:        req.Limit,
			SortBy:       req.SortBy,
			OrderBy:      req.OrderBy,
		})
		if err != nil {
			return nil, err
		}

		totalPage := math.Ceil(float64(resp.Metadata.Total) / float64(*req.Limit))

		data := &UserPostWithMetadata{
			Data: resp.Data,
			Metadata: &Metadata{
				PerPage:     helper.GetInt64FromPointer(req.Limit),
				TotalPage:   totalPage,
				Total:       resp.Metadata.Total,
				CurrentPage: helper.GetInt64FromPointer(req.Page),
			},
		}
		return map[string]interface{}{
			"data": data,
		}, nil
	}
}

func MakeGetDetailUserPost(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetByID)
		// TODO: for get metadata from headers grpc needs to update when using authorization
		resp, err := usecase.GetDetailPost(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"data": resp}, nil
	}
}

func MakeGetListUserPostByMe(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetListUserPostRequest)
		// TODO: for get metadata from headers grpc needs to update when using authorization
		resp, err := usecase.GetListPostByMe(ctx, &model.GetListRequest{
			ActivityName: req.ActivityName,
			Username:     req.Username,
			Category:     req.Category,
			Status:       req.Status,
			Page:         req.Page,
			Limit:        req.Limit,
			SortBy:       req.SortBy,
			OrderBy:      req.OrderBy,
		})
		if err != nil {
			return nil, err
		}
		totalPage := math.Ceil(float64(resp.Metadata.Total) / float64(*req.Limit))

		data := &UserPostWithMetadata{
			Data: resp.Data,
			Metadata: &Metadata{
				PerPage:     helper.GetInt64FromPointer(req.Limit),
				TotalPage:   totalPage,
				Total:       resp.Metadata.Total,
				CurrentPage: helper.GetInt64FromPointer(req.Page),
			},
		}
		return map[string]interface{}{
			"data": data}, nil
	}
}

func MakeCreateNewPost(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CreateNewPostRequest)
		if err := Validate(req); err != nil {
			return nil, err
		}

		imagePathURL := req.Images[0]
		imagesFormatted, err := json.Marshal(req.Images)
		if err != nil {
			return nil, err
		}

		bodyRequest := &model.CreateNewPostRequest{
			Title:        helper.GetStringFromPointer(req.Title),
			ImagePathURL: imagePathURL.Path,
			Images:       string(imagesFormatted),
			Tags:         req.Tags,
			Status:       helper.GetInt64FromPointer(req.Status),
		}

		if err = usecase.CreateNewPost(ctx, bodyRequest); err != nil {
			return nil, err
		}

		return &StatusResponse{
			Code:    helper.STATUS_CREATED,
			Message: "a_post_has_been_created",
		}, nil
	}
}

func MakeUpdateStatusOrTitle(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CreateCommentRequest)
		if err := Validate(req); err != nil {
			return nil, err
		}

		if err = usecase.UpdateTitleOrStatus(ctx, &model.UpdatePostRequest{
			ID:     req.UserPostID,
			Status: req.Status,
			Title:  helper.SetPointerString(req.Text),
		}); err != nil {
			return nil, err
		}
		return &StatusResponse{
			Code:    helper.STATUS_UPDATED,
			Message: "post_has_been_updated",
		}, nil
	}
}

func MakeGetCommentsByID(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetComment)
		resp, err := usecase.GetCommentsByPostID(ctx, &model.GetCommentRequest{
			ID:   req.ID,
			Page: req.Page,
		})
		if err != nil {
			return nil, err
		}

		data := &CommentsResponse{
			Data: resp.Data,
			Metadata: &Metadata{
				PerPage:     20,
				Total:       resp.Metadata.Total,
				TotalPage:   math.Ceil(float64(resp.Metadata.Total) / float64(20)),
				CurrentPage: resp.Metadata.Page,
			},
		}

		return map[string]interface{}{
			"data": data}, nil
	}
}

func MakeCreateComment(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CreateCommentRequest)
		if err := Validate(req); err != nil {
			return nil, err
		}

		if err = usecase.CreateCommentOnPost(ctx, &model.CreateCommentRequest{
			UserPostID: req.UserPostID,
			Text:       req.Text,
			Status:     helper.GetInt64FromPointer(req.Status),
		}); err != nil {
			return nil, err
		}

		return &StatusResponse{
			Code:    helper.STATUS_CREATED,
			Message: "success_post_comment",
		}, nil
	}
}

func MakeLikeOrDislikePost(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetByID)
		// TODO: for get metadata from headers grpc needs to update when using authorization
		if err = usecase.LikeOrDislikePost(ctx, req.ID); err != nil {
			return nil, err
		}

		return &StatusResponse{
			Code:    helper.STATUS_CREATED,
			Message: "success_like_or_dislike_a_post",
		}, nil
	}
}

func MakeCheckHealthy(ctx context.Context) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &StatusResponse{
			Code:    helper.STATUS_OK,
			Message: "service_is_ok",
		}, nil
	}
}

func MakeCheckReadiness(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if err := usecase.CheckHealthReadiness(ctx); err != nil {
			return nil, err
		}
		return &StatusResponse{
			Code:    helper.STATUS_OK,
			Message: "service_is_ready",
		}, nil
	}
}
