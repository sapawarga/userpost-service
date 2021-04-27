package repository

import (
	"context"

	"github.com/sapawarga/userpost-service/model"
)

type PostI interface {
	// query for get userpost
	GetListPost(ctx context.Context, request *model.UserPostRequest) ([]*model.PostResponse, error)
	GetMetadataPost(ctx context.Context, request *model.UserPostRequest) (*int64, error)
	GetListPostByMe(ctx context.Context, request *model.UserPostByMeRequest) ([]*model.PostResponse, error)
	GetMetadataPostByMe(ctx context.Context, request *model.UserPostByMeRequest) (*int64, error)
	GetActor(ctx context.Context, id int64) (*model.UserResponse, error)
	GetIsLikedByUser(ctx context.Context, req *model.IsLikedByUser) (bool, error)
	GetDetailPost(ctx context.Context, id int64) (*model.PostResponse, error)
	// query for create post
	InsertPost(ctx context.Context, request *model.CreateNewPostRequestRepository) error
	// query for update
	UpdateStatusOrTitle(ctx context.Context, request *model.UpdatePostRequest) error
}

type CommentI interface {
	GetLastComment(ctx context.Context, id int64) (*model.CommentResponse, error)
	GetTotalComments(ctx context.Context, userPostID int64) (*int64, error)
	GetCommentsByPostID(ctx context.Context, id int64) ([]*model.CommentResponse, error)
	Create(ctx context.Context, req *model.CreateCommentRequest) error
}
