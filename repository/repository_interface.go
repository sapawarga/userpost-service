package repository

import (
	"context"

	"github.com/sapawarga/userpost-service/model"
)

type PostI interface {
	GetListPost(ctx context.Context, request *model.UserPostRequest) ([]*model.PostResponse, error)
	GetMetadataPost(ctx context.Context, request *model.UserPostRequest) (*int64, error)
	GetActor(ctx context.Context, id int64) (*model.UserResponse, error)
	GetLastComment(ctx context.Context, id int64) (*model.CommentResponse, error)
	GetTotalComments(ctx context.Context, userPostID int64) (*int64, error)
	GetIsLikedByUser(ctx context.Context, req *model.IsLikedByUser) (bool, error)
}
