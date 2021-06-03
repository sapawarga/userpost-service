package testcases

import (
	"errors"

	"github.com/sapawarga/userpost-service/model"
)

type ResponseGetCommentsRepository struct {
	Result []*model.CommentResponse
	Error  error
}

type ResponseGetCommentsUsecase struct {
	Result []*model.Comment
	Error  error
}

type GetComments struct {
	Description                      string
	UsecaseParams                    int64
	GetCommentsByIDRequestRepository int64
	GetActorParams                   int64
	ResponseGetComments              ResponseGetCommentsRepository
	ResponseUsecase                  ResponseGetCommentsUsecase
	ResponseGetActor                 ResponseGetActor
}

var (
	commentsRepository = []*model.CommentResponse{
		{
			ID:         1,
			Comment:    "comment",
			UserPostID: 1,
			CreatedAt:  current,
			UpdatedAt:  current,
			CreatedBy:  1,
			UpdatedBy:  1,
		}, {
			ID:         2,
			Comment:    "ini juga comment",
			UserPostID: 1,
			CreatedAt:  current,
			UpdatedAt:  current,
			CreatedBy:  1,
			UpdatedBy:  1,
		},
	}
	commentsUsecase = []*model.Comment{
		{
			ID:         1,
			UserPostID: 1,
			Text:       "comment",
			CreatedAt:  current,
			UpdatedAt:  current,
			CreatedBy:  actor,
			UpdatedBy:  actor,
		}, {
			ID:         2,
			UserPostID: 1,
			Text:       "ini juga comment",
			CreatedAt:  current,
			UpdatedAt:  current,
			CreatedBy:  actor,
			UpdatedBy:  actor,
		},
	}
)

var GetCommentsData = []GetComments{
	{
		Description:                      "success_get_list_comments",
		UsecaseParams:                    1,
		GetCommentsByIDRequestRepository: 1,
		GetActorParams:                   1,
		ResponseGetComments: ResponseGetCommentsRepository{
			Result: commentsRepository,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
			Error:  nil,
		},
		ResponseUsecase: ResponseGetCommentsUsecase{
			Result: commentsUsecase,
			Error:  nil,
		},
	}, {
		Description:                      "success_get_list_comments_even_nil_comment",
		UsecaseParams:                    1,
		GetCommentsByIDRequestRepository: 1,
		GetActorParams:                   1,
		ResponseGetComments: ResponseGetCommentsRepository{
			Result: nil,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: nil,
			Error:  nil,
		},
		ResponseUsecase: ResponseGetCommentsUsecase{
			Result: nil,
			Error:  nil,
		},
	}, {
		Description:                      "failed_get_list_comments",
		UsecaseParams:                    1,
		GetCommentsByIDRequestRepository: 1,
		GetActorParams:                   1,
		ResponseGetComments: ResponseGetCommentsRepository{
			Result: nil,
			Error:  errors.New("failed_get_comments"),
		},
		ResponseGetActor: ResponseGetActor{
			Result: nil,
			Error:  errors.New("failed_get_comments"),
		},
		ResponseUsecase: ResponseGetCommentsUsecase{
			Result: nil,
			Error:  errors.New("failed_get_comments"),
		},
	},
}

func ListGetCommentsDescription() []string {
	var arr = []string{}
	for _, data := range GetCommentsData {
		arr = append(arr, data.Description)
	}
	return arr
}
