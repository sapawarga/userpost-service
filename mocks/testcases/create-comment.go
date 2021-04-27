package testcases

import (
	"errors"

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

type CreateCommentOnAPost struct {
	Description       string
	UsecaseRequest    *model.CreateCommentRequest
	RepositoryRequest *model.CreateCommentRequestRepository
	MockRepository    error
	MockUsecase       error
}

var CreateCommentOnAPostData = []CreateCommentOnAPost{
	{
		Description:       "success_create_comment",
		UsecaseRequest:    newComment,
		RepositoryRequest: newCommentRepository,
		MockRepository:    nil,
		MockUsecase:       nil,
	}, {
		Description:       "failed_create_comment",
		UsecaseRequest:    newComment,
		RepositoryRequest: newCommentRepository,
		MockRepository:    errors.New("something_went_wrong"),
		MockUsecase:       errors.New("something_went_wrong"),
	},
}

func CreateCommentDescription() []string {
	var arr = []string{}
	for _, data := range CreateCommentOnAPostData {
		arr = append(arr, data.Description)
	}
	return arr
}
