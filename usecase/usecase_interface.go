package usecase

import (
	"context"

	"github.com/sapawarga/userpost-service/model"
)

type UsecaseI interface {
	GetListPost(ctx context.Context, params *model.GetListRequest) (*model.UserPostWithMetadata, error)
}
