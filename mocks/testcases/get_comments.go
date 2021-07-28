package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/userpost-service/lib/convert"
	"github.com/sapawarga/userpost-service/model"
)

type ResponseGetCommentsRepository struct {
	Result []*model.CommentResponse
	Error  error
}

type ResponseGetCommentsUsecase struct {
	Result *model.CommentWithMetadata
	Error  error
}

var amountComments = convert.SetPointerInt64(2)

type GetComments struct {
	Description                      string
	UsecaseParams                    *model.GetCommentRequest
	GetCommentsByIDRequestRepository *model.GetComment
	GetActorParams                   int64
	GetTotalComment                  int64
	ResponseGetTotalComment          ResponseGetTotalComment
	ResponseGetComments              ResponseGetCommentsRepository
	ResponseUsecase                  ResponseGetCommentsUsecase
	ResponseGetActor                 ResponseGetActor
}

var (
	commentsRepository = []*model.CommentResponse{
		{
			ID:         1,
			Comment:    "comment",
			UserPostID: sql.NullInt64{Int64: 1, Valid: true},
			CreatedAt:  current,
			UpdatedAt:  current,
			CreatedBy:  sql.NullInt64{Int64: 1, Valid: true},
			UpdatedBy:  sql.NullInt64{Int64: 1, Valid: true},
		}, {
			ID:         2,
			Comment:    "ini juga comment",
			UserPostID: sql.NullInt64{Int64: 1, Valid: true},
			CreatedAt:  current,
			UpdatedAt:  current,
			CreatedBy:  sql.NullInt64{Int64: 1, Valid: true},
			UpdatedBy:  sql.NullInt64{Int64: 1, Valid: true},
		},
	}
	commentsUsecase = []*model.Comment{
		{
			ID:         1,
			UserPostID: 1,
			Text:       "comment",
			CreatedAt:  current,
			UpdatedAt:  current,
			User:       actor,
			CreatedBy:  1,
			UpdatedBy:  1,
		}, {
			ID:         2,
			UserPostID: 1,
			Text:       "ini juga comment",
			User:       actor,
			CreatedAt:  current,
			UpdatedAt:  current,
			CreatedBy:  1,
			UpdatedBy:  1,
		},
	}

	responseUsecase = &model.CommentWithMetadata{
		Data: commentsUsecase,
		Metadata: &model.Metadata{
			Page:  1,
			Total: 2,
		},
	}
)

var reqUsecase = &model.GetCommentRequest{
	ID:   1,
	Page: 1,
}

var reqRepository = &model.GetComment{
	ID:     1,
	Limit:  20,
	Offset: 0,
}

var GetCommentsData = []GetComments{
	{
		Description:                      "success_get_list_comments",
		UsecaseParams:                    reqUsecase,
		GetCommentsByIDRequestRepository: reqRepository,
		GetTotalComment:                  1,
		GetActorParams:                   1,
		ResponseGetComments: ResponseGetCommentsRepository{
			Result: commentsRepository,
			Error:  nil,
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: amountComments,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
			Error:  nil,
		},
		ResponseUsecase: ResponseGetCommentsUsecase{
			Result: responseUsecase,
			Error:  nil,
		},
	}, {
		Description:                      "success_get_list_comments_even_nil_comment",
		UsecaseParams:                    reqUsecase,
		GetCommentsByIDRequestRepository: reqRepository,
		GetActorParams:                   1,
		GetTotalComment:                  1,
		ResponseGetComments: ResponseGetCommentsRepository{
			Result: nil,
			Error:  nil,
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: amountComments,
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
		UsecaseParams:                    reqUsecase,
		GetCommentsByIDRequestRepository: reqRepository,
		GetTotalComment:                  1,
		GetActorParams:                   1,
		ResponseGetComments: ResponseGetCommentsRepository{
			Result: nil,
			Error:  errors.New("failed_get_comments"),
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: nil,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: nil,
			Error:  errors.New("failed_get_comments"),
		},
		ResponseUsecase: ResponseGetCommentsUsecase{
			Result: nil,
			Error:  errors.New("failed_get_comments"),
		},
	}, {
		Description:                      "failed_get_metadata_comments",
		UsecaseParams:                    reqUsecase,
		GetCommentsByIDRequestRepository: reqRepository,
		GetTotalComment:                  1,
		GetActorParams:                   1,
		ResponseGetComments: ResponseGetCommentsRepository{
			Result: commentsRepository,
			Error:  nil,
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
		},
		ResponseGetActor: ResponseGetActor{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
		},
		ResponseUsecase: ResponseGetCommentsUsecase{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
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
