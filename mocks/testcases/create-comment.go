package testcases

import (
	"errors"

	"github.com/sapawarga/userpost-service/lib/convert"
	"github.com/sapawarga/userpost-service/model"
)

var newComment = &model.CreateCommentRequest{
	UserPostID: 1,
	Text:       "this is comment",
}

var newCommentRepository = &model.CreateCommentRequestRepository{
	UserPostID: 1,
	Text:       "this is comment",
	ActorID:    1,
}

type ResponseCreateComment struct {
	ID    int64
	Error error
}

var updateTotalCommentPost = &model.UpdatePostRequest{
	ID:            1,
	LastCommentID: convert.SetPointerInt64(1),
}

type CreateCommentOnAPost struct {
	Description       string
	UsecaseRequest    *model.CreateCommentRequest
	UpdatePostRequest *model.UpdatePostRequest
	RepositoryRequest *model.CreateCommentRequestRepository
	MockRepository    ResponseCreateComment
	MockUpdatePost    error
	MockUsecase       error
}

var CreateCommentOnAPostData = []CreateCommentOnAPost{
	{
		Description:       "success_create_comment",
		UsecaseRequest:    newComment,
		RepositoryRequest: newCommentRepository,
		UpdatePostRequest: updateTotalCommentPost,
		MockUpdatePost:    nil,
		MockRepository: ResponseCreateComment{
			ID:    1,
			Error: nil,
		},
		MockUsecase: nil,
	}, {
		Description:       "failed_create_comment",
		UsecaseRequest:    newComment,
		RepositoryRequest: newCommentRepository,
		UpdatePostRequest: updateTotalCommentPost,
		MockUpdatePost:    nil,
		MockRepository: ResponseCreateComment{
			ID:    0,
			Error: errors.New("something_went_wrong"),
		},
		MockUsecase: errors.New("something_went_wrong"),
	}, {
		Description:       "failed_update_post",
		UsecaseRequest:    newComment,
		RepositoryRequest: newCommentRepository,
		UpdatePostRequest: updateTotalCommentPost,
		MockUpdatePost:    errors.New("something_went_wrong"),
		MockRepository: ResponseCreateComment{
			ID:    1,
			Error: nil,
		},
		MockUsecase: errors.New("something_went_wrong"),
	},
}

func CreateCommentDescription() []string {
	var arr = []string{}
	for _, data := range CreateCommentOnAPostData {
		arr = append(arr, data.Description)
	}
	return arr
}
