package usecase

import (
	"context"

	"github.com/sapawarga/userpost-service/model"
)

type UsecaseI interface {
	GetListPost(ctx context.Context, params *model.GetListRequest) (*model.UserPostWithMetadata, error)
	GetListPostByMe(ctx context.Context, params *model.GetListRequest) (*model.UserPostWithMetadata, error)
	GetDetailPost(ctx context.Context, id int64) (*model.UserPostResponse, error)
	CreateNewPost(ctx context.Context, requestBody *model.CreateNewPostRequest) error
	UpdateTitleOrStatus(ctx context.Context, requestBody *model.UpdatePostRequest) error
	GetCommentsByPostID(ctx context.Context, id int64) ([]*model.Comment, error)
	CreateCommentOnPost(ctx context.Context, req *model.CreateCommentRequest) error
}
