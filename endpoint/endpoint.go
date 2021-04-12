package endpoint

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/endpoint"
	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
	"github.com/sapawarga/userpost-service/usecase"
)

func MakeGetListUserPost(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetListUserPostRequest)
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
		return &UserPostWithMetadata{
			Data: resp.Data,
			Metadata: &Metadata{
				Page:      resp.Metadata.Page,
				TotalPage: resp.Metadata.TotalPage,
				Total:     resp.Metadata.Total,
			},
		}, nil
	}
}

func MakeGetDetailUserPost(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetByID)
		resp, err := usecase.GetDetailPost(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return resp, nil
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
			Code:    helper.STATUSCREATED,
			Message: "a_post_has_been_created",
		}, nil
	}
}

func MakeUpdateStatusOrTitle(ctx context.Context, usecase usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*UpdateStatusOrTitle)
		if err := Validate(req); err != nil {
			return nil, err
		}

		if err = usecase.UpdateTitleOrStatus(ctx, &model.UpdatePostRequest{
			ID:     req.ID,
			Status: req.Status,
			Title:  req.Title,
		}); err != nil {
			return nil, err
		}
		return &StatusResponse{
			Code:    helper.STATUSUPDATED,
			Message: "post_has_been_updated",
		}, nil
	}
}
