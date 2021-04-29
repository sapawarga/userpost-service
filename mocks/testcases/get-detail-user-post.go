package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
)

type ResponseGetDetailUserPost struct {
	Result *model.PostResponse
	Error  error
}

type ResponseGetDetailUsecase struct {
	Result *model.UserPostResponse
	Error  error
}

var (
	userpostResponse = &model.UserPostResponse{
		ID:                    1,
		Title:                 "title",
		Tag:                   helper.SetPointerString("tag"),
		ImagePath:             "test",
		Images:                "test",
		LastUserPostCommentID: helper.SetPointerInt64(1),
		LastComment:           comment,
		LikesCount:            0,
		IsLiked:               true,
		CommentCounts:         1,
		Status:                10,
		Actor:                 actorResponse,
		CreatedAt:             current,
		UpdatedAt:             current,
	}
	postDetail = &model.PostResponse{
		ID:                    1,
		Title:                 "title",
		Tag:                   sql.NullString{String: "Category", Valid: true},
		ImagePath:             sql.NullString{String: "test", Valid: true},
		Images:                sql.NullString{String: "test", Valid: true},
		LastUserPostCommentID: sql.NullInt64{Int64: 1, Valid: true},
		LikesCount:            1,
		CommentCounts:         1,
		Status:                10,
		CreatedBy:             sql.NullInt64{Int64: 1, Valid: true},
		UpdatedBy:             sql.NullInt64{Int64: 1, Valid: true},
		CreatedAt:             current,
		UpdatedAt:             current,
	}
)

type GetDetailUserPost struct {
	Description            string
	UsecaseParams          int64
	GetUserPostParams      int64
	GetActorParams         int64
	GetLastCommentParams   int64
	GetTotalCommentsParams int64
	IsLikedRequest         *model.AddOrRemoveLikeOnPostRequest
	CheckIsLikedResponse
	ResponseGetDetailUserPost
	ResponseGetActor
	ResponseGetLastComment
	ResponseGetTotalComment
	ResponseGetDetailUsecase
}

var GetDetailUserPostData = []GetDetailUserPost{
	{
		Description:            "success_get_detail_user_post",
		UsecaseParams:          1,
		GetUserPostParams:      1,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		ResponseGetDetailUserPost: ResponseGetDetailUserPost{
			Result: postDetail,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
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
		ResponseGetDetailUsecase: ResponseGetDetailUsecase{
			Result: userpostResponse,
			Error:  nil,
		},
	}, {
		Description:            "failed_get_detail",
		UsecaseParams:          1,
		GetUserPostParams:      1,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		ResponseGetDetailUserPost: ResponseGetDetailUserPost{
			Result: nil,
			Error:  errors.New("failed_get_detail"),
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
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
		ResponseGetDetailUsecase: ResponseGetDetailUsecase{
			Result: nil,
			Error:  errors.New("failed_get_detail"),
		},
	}, {
		Description:            "failed_get_comment",
		UsecaseParams:          1,
		GetUserPostParams:      1,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		ResponseGetDetailUserPost: ResponseGetDetailUserPost{
			Result: postDetail,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
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
		ResponseGetDetailUsecase: ResponseGetDetailUsecase{
			Result: nil,
			Error:  errors.New("failed_get_comment"),
		},
	}, {
		Description:            "failed_get_actor",
		UsecaseParams:          1,
		GetUserPostParams:      1,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: true,
			Error:  nil,
		},
		ResponseGetDetailUserPost: ResponseGetDetailUserPost{
			Result: postDetail,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: nil,
			Error:  errors.New("failed_get_actor"),
		},
		ResponseGetLastComment: ResponseGetLastComment{
			Result: commentResponse,
			Error:  nil,
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: totalComment,
			Error:  nil,
		},
		ResponseGetDetailUsecase: ResponseGetDetailUsecase{
			Result: nil,
			Error:  errors.New("failed_get_actor"),
		},
	}, {
		Description:            "failed_get_is_liked_post",
		UsecaseParams:          1,
		GetUserPostParams:      1,
		GetActorParams:         1,
		GetLastCommentParams:   1,
		GetTotalCommentsParams: 1,
		IsLikedRequest:         requestLikeOnPost,
		CheckIsLikedResponse: CheckIsLikedResponse{
			Result: false,
			Error:  errors.New("something_went_wrong"),
		},
		ResponseGetDetailUserPost: ResponseGetDetailUserPost{
			Result: postDetail,
			Error:  nil,
		},
		ResponseGetActor: ResponseGetActor{
			Result: actorResponse,
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
		ResponseGetDetailUsecase: ResponseGetDetailUsecase{
			Result: nil,
			Error:  errors.New("something_went_wrong"),
		},
	},
}

func ListUserPostDetailDescription() []string {
	var arr = []string{}
	for _, data := range GetDetailUserPostData {
		arr = append(arr, data.Description)
	}
	return arr
}
