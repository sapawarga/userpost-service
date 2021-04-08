package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
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
