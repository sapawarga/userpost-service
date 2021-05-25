package testcases

import (
	"errors"

	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
)

var tags = helper.SetPointerString("categories")

var newPostRequest = &model.CreateNewPostRequest{
	Title:        "test",
	ImagePathURL: "http://localhost",
	Images:       "[{\"path\":\"http://localhost\"}]",
	Tags:         tags,
	Status:       helper.ACTIVED,
}

var newRepositoryRequest = &model.CreateNewPostRequestRepository{
	Title:        "test",
	ImagePathURL: "http://localhost",
	Images:       "[{\"path\":\"http://localhost\"}]",
	Tags:         tags,
	Status:       helper.ACTIVED,
	// ActorID:      1,
}

type CreateNewUserPost struct {
	Description       string
	UsecaseRequest    *model.CreateNewPostRequest
	RepositoryRequest *model.CreateNewPostRequestRepository
	MockRepository    error
	MockUsecase       error
}

var CreateNewUserPostData = []CreateNewUserPost{
	{
		Description:       "succes_insert_new_post",
		UsecaseRequest:    newPostRequest,
		RepositoryRequest: newRepositoryRequest,
		MockRepository:    nil,
		MockUsecase:       nil,
	}, {
		Description:       "failed_insert_new_post",
		UsecaseRequest:    newPostRequest,
		RepositoryRequest: newRepositoryRequest,
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
