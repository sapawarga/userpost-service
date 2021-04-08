package testcases

import (
	"errors"

	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
)

var newPostRequest = &model.CreateNewPostRequest{
	Title:        "test",
	ImagePathURL: "http://localhost",
	Images:       "[{\"path\":\"http://localhost\"}]",
	Tags:         helper.SetPointerString("categories"),
	Status:       helper.ACTIVED,
}

type CreateNewUserPost struct {
	Description       string
	UsecaseRequest    *model.CreateNewPostRequest
	RepositoryRequest *model.CreateNewPostRequest
	MockRepository    error
	MockUsecase       error
}

var CreateNewUserPostData = []CreateNewUserPost{
	{
		Description:       "succes_insert_new_post",
		UsecaseRequest:    newPostRequest,
		RepositoryRequest: newPostRequest,
		MockRepository:    nil,
		MockUsecase:       nil,
	}, {
		Description:       "failed_insert_new_post",
		UsecaseRequest:    newPostRequest,
		RepositoryRequest: newPostRequest,
		MockRepository:    errors.New("something_went_wrong"),
		MockUsecase:       errors.New("something_went_wrong"),
	},
}

func CreateNewUserPostDescription() []string {
	var arr = []string{}
	for _, data := range CreateNewUserPostData {
		arr = append(arr, data.Description)
	}
	return arr
}
