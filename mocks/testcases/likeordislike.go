package testcases

import (
	"errors"

	"github.com/sapawarga/userpost-service/model"
)

var requestLikeOnPost = &model.AddOrRemoveLikeOnPostRequest{
	UserPostID: 1,
	// ActorID:    1,
	TypeEntity: "user_post",
}

type CheckIsLikedResponse struct {
	Result bool
	Error  error
}

type LikeOrDislikePost struct {
	Description             string
	UsecaseRequest          int64
	CheckIsLikedRequest     *model.AddOrRemoveLikeOnPostRequest
	AddLikeOnPostRequest    *model.AddOrRemoveLikeOnPostRequest
	RemoveLikeOnPostRequest *model.AddOrRemoveLikeOnPostRequest
	MockCheckIsLiked        CheckIsLikedResponse
	MockAddLikeOnPost       error
	MockRemoveLikeOnPost    error
	MockUsecase             error
}

var LikeOrDislikePostData = []LikeOrDislikePost{
	{
		Description:             "success_like_a_post",
		UsecaseRequest:          1,
		CheckIsLikedRequest:     requestLikeOnPost,
		AddLikeOnPostRequest:    requestLikeOnPost,
		RemoveLikeOnPostRequest: requestLikeOnPost,
		MockCheckIsLiked: CheckIsLikedResponse{
			Result: false,
			Error:  nil,
		},
		MockAddLikeOnPost:    nil,
		MockRemoveLikeOnPost: nil,
		MockUsecase:          nil,
	}, {
		Description:             "success_dislike_a_post",
		UsecaseRequest:          1,
		CheckIsLikedRequest:     requestLikeOnPost,
		AddLikeOnPostRequest:    requestLikeOnPost,
		RemoveLikeOnPostRequest: requestLikeOnPost,
		MockCheckIsLiked: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		MockAddLikeOnPost:    nil,
		MockRemoveLikeOnPost: nil,
		MockUsecase:          nil,
	}, {
		Description:             "failed_check_is_liked",
		UsecaseRequest:          1,
		CheckIsLikedRequest:     requestLikeOnPost,
		AddLikeOnPostRequest:    requestLikeOnPost,
		RemoveLikeOnPostRequest: requestLikeOnPost,
		MockCheckIsLiked: CheckIsLikedResponse{
			Result: false,
			Error:  errors.New("failed_check_is_liked"),
		},
		MockAddLikeOnPost:    errors.New("failed_check_is_liked"),
		MockRemoveLikeOnPost: errors.New("failed_check_is_liked"),
		MockUsecase:          errors.New("failed_check_is_liked"),
	}, {
		Description:             "failed_add_liked",
		UsecaseRequest:          1,
		CheckIsLikedRequest:     requestLikeOnPost,
		AddLikeOnPostRequest:    requestLikeOnPost,
		RemoveLikeOnPostRequest: requestLikeOnPost,
		MockCheckIsLiked: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		MockAddLikeOnPost:    errors.New("failed_add_liked"),
		MockRemoveLikeOnPost: nil,
		MockUsecase:          errors.New("failed_add_liked"),
	}, {
		Description:             "failed_remove_liked",
		UsecaseRequest:          1,
		CheckIsLikedRequest:     requestLikeOnPost,
		AddLikeOnPostRequest:    requestLikeOnPost,
		RemoveLikeOnPostRequest: requestLikeOnPost,
		MockCheckIsLiked: CheckIsLikedResponse{
			Result: false,
			Error:  nil,
		},
		MockAddLikeOnPost:    nil,
		MockRemoveLikeOnPost: errors.New("failed_remove_liked"),
		MockUsecase:          errors.New("failed_remove_liked"),
	},
}

func LikeOrDislikePostDescription() []string {
	var arr = []string{}
	for _, data := range LikeOrDislikePostData {
		arr = append(arr, data.Description)
	}
	return arr
}
