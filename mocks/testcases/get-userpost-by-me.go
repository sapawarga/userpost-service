package testcases

import (
	"errors"

	"github.com/sapawarga/userpost-service/model"
)

type GetListUserPostByMe struct {
	Description            string
	UsecaseParams          *model.GetListRequest
	GetUserPostParams      *model.UserPostByMeRequest
	GetMetadataParams      *model.UserPostByMeRequest
	GetActorParams         int64
	GetLastCommentParams   int64
	GetTotalCommentsParams int64
	IsLikedRequest         *model.AddOrRemoveLikeOnPostRequest
	CheckIsLikedResponse
	ResponseGetList
	ResponseGetActor
	ResponseMetadata
	ResponseGetLastComment
	ResponseGetTotalComment
	ResponseUsecase
}

var userpostbyme = &model.UserPostByMeRequest{
	ActorID:         1,
	UserPostRequest: &userPostParams,
}

var GetListUserPostByMeData = []GetListUserPostByMe{
	{
		Description:            "success_get_list_user_post_by_me",
		UsecaseParams:          &requestUsecase,
		GetUserPostParams:      userpostbyme,
		GetMetadataParams:      userpostbyme,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		ResponseGetList: ResponseGetList{
			Result: postResponse,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
			Error:  nil,
		},
		ResponseMetadata: ResponseMetadata{
			Result: metadataResponse,
			Error:  nil,
		},
		ResponseGetLastComment: ResponseGetLastComment{
			Result: commentResponse,
			Error:  nil,
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: totalComment,
			Error:  nil,
		},
		ResponseUsecase: ResponseUsecase{
			Result: usecaseResponse,
			Error:  nil,
		},
	}, {
		Description:            "failed_get_list_user_post_by_me",
		UsecaseParams:          &requestUsecase,
		GetUserPostParams:      userpostbyme,
		GetMetadataParams:      userpostbyme,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		ResponseGetList: ResponseGetList{
			Result: nil,
			Error:  errors.New("failed_get_user_posts_by_me"),
		},
		ResponseGetActor: ResponseGetActor{
			Result: nil,
			Error:  nil,
		},
		ResponseMetadata: ResponseMetadata{
			Result: nil,
			Error:  nil,
		},
		ResponseGetLastComment: ResponseGetLastComment{
			Result: nil,
			Error:  nil,
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: nil,
			Error:  nil,
		},
		ResponseUsecase: ResponseUsecase{
			Result: nil,
			Error:  errors.New("failed_get_user_posts_by_me"),
		},
	}, {
		Description:            "failed_get_actor",
		UsecaseParams:          &requestUsecase,
		GetUserPostParams:      userpostbyme,
		GetMetadataParams:      userpostbyme,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		ResponseGetList: ResponseGetList{
			Result: postResponse,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: nil,
			Error:  errors.New("failed_get_actor"),
		},
		ResponseMetadata: ResponseMetadata{
			Result: metadataResponse,
			Error:  nil,
		},
		ResponseGetLastComment: ResponseGetLastComment{
			Result: nil,
			Error:  errors.New("failed_get_actor"),
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: nil,
			Error:  errors.New("failed_get_actor"),
		},
		ResponseUsecase: ResponseUsecase{
			Result: nil,
			Error:  errors.New("failed_get_actor"),
		},
	}, {
		Description:            "failed_get_comment",
		UsecaseParams:          &requestUsecase,
		GetUserPostParams:      userpostbyme,
		GetMetadataParams:      userpostbyme,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		ResponseGetList: ResponseGetList{
			Result: postResponse,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
			Error:  nil,
		},
		ResponseMetadata: ResponseMetadata{
			Result: metadataResponse,
			Error:  nil,
		},
		ResponseGetLastComment: ResponseGetLastComment{
			Result: nil,
			Error:  errors.New("failed_get_comment"),
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: totalComment,
			Error:  nil,
		},
		ResponseUsecase: ResponseUsecase{
			Result: nil,
			Error:  errors.New("failed_get_comment"),
		},
	}, {
		Description:            "failed_get_total_comment",
		UsecaseParams:          &requestUsecase,
		GetUserPostParams:      userpostbyme,
		GetMetadataParams:      userpostbyme,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		ResponseGetList: ResponseGetList{
			Result: postResponse,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
			Error:  nil,
		},
		ResponseMetadata: ResponseMetadata{
			Result: metadataResponse,
			Error:  nil,
		},
		ResponseGetLastComment: ResponseGetLastComment{
			Result: nil,
			Error:  errors.New("failed_get_total_comment"),
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: nil,
			Error:  errors.New("failed_get_total_comment"),
		},
		ResponseUsecase: ResponseUsecase{
			Result: nil,
			Error:  errors.New("failed_get_total_comment"),
		},
	}, {
		Description:            "failed_get_metadata",
		UsecaseParams:          &requestUsecase,
		GetUserPostParams:      userpostbyme,
		GetMetadataParams:      userpostbyme,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		ResponseGetList: ResponseGetList{
			Result: postResponse,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
			Error:  nil,
		},
		ResponseMetadata: ResponseMetadata{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
		},
		ResponseGetLastComment: ResponseGetLastComment{
			Result: commentResponse,
			Error:  nil,
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: totalComment,
			Error:  nil,
		},
		ResponseUsecase: ResponseUsecase{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
		},
	}, {
		Description:            "failed_get_metadata",
		UsecaseParams:          &requestUsecase,
		GetUserPostParams:      userpostbyme,
		GetMetadataParams:      userpostbyme,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: false,
			Error:  errors.New("something_went_wrong"),
		},
		ResponseGetList: ResponseGetList{
			Result: postResponse,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
			Error:  nil,
		},
		ResponseMetadata: ResponseMetadata{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
		},
		ResponseGetLastComment: ResponseGetLastComment{
			Result: commentResponse,
			Error:  nil,
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: totalComment,
			Error:  nil,
		},
		ResponseUsecase: ResponseUsecase{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
		},
	},
}

func ListUserPostByMeDescription() []string {
	var arr = []string{}
	for _, data := range GetListUserPostByMeData {
		arr = append(arr, data.Description)
	}
	return arr
}
