package testcases

import (
	"errors"

	"github.com/sapawarga/userpost-service/model"
)

var newComment = &model.CreateCommentRequest{
	UserPostID: 1,
	Text:       "this is comment",
	CreatedAt:  current,
	UpdatedAt:  current,
}

type CreateCommentOnAPost struct {
	Description       string
	UsecaseRequest    *model.CreateCommentRequest
	RepositoryRequest *model.CreateCommentRequest
	MockRepository    error
	MockUsecase       error
}

var CreateCommentOnAPostData = []CreateCommentOnAPost{
	{
		Description:       "success_create_comment",
		UsecaseRequest:    newComment,
		RepositoryRequest: newComment,
		MockRepository:    nil,
		MockUsecase:       nil,
	}, {
		Description:       "failed_create_comment",
		UsecaseRequest:    newComment,
		RepositoryRequest: newComment,
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
