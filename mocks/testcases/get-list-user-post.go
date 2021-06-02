package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
)

type UsecaseResponse struct {
	Result *model.UserPostWithMetadata
	Error  error
}

type ResponseGetList struct {
	Result []*model.PostResponse
	Error  error
}

type ResponseGetLastComment struct {
	Result *model.CommentResponse
	Error  error
}

type ResponseGetTotalComment struct {
	Result *int64
	Error  error
}

type ResponseGetActor struct {
	Result *model.UserResponse
	Error  error
}

type ResponseMetadata struct {
	Result *int64
	Error  error
}

type ResponseUsecase struct {
	Result *model.UserPostWithMetadata
	Error  error
}

type GetListUserPost struct {
	Description            string
	UsecaseParams          model.GetListRequest
	GetUserPostParams      model.UserPostRequest
	GetMetadataParams      model.UserPostRequest
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

var (
	testRandomString = helper.SetPointerString("random-test")
	statusNumber     = helper.SetPointerInt64(10)
	requestUsecase   = model.GetListRequest{
		ActivityName: testRandomString,
		Username:     testRandomString,
		Category:     testRandomString,
		Status:       statusNumber,
		Page:         helper.SetPointerInt64(1),
		Limit:        helper.SetPointerInt64(10),
		SortBy:       nil,
		OrderBy:      nil,
	}
	userPostParams = model.UserPostRequest{
		ActivityName: requestUsecase.ActivityName,
		Username:     requestUsecase.Username,
		Category:     requestUsecase.Category,
		Status:       requestUsecase.Status,
		Offset:       helper.SetPointerInt64(0),
		Limit:        requestUsecase.Limit,
		SortBy:       nil,
		OrderBy:      nil,
	}
	current, _   = helper.GetCurrentTimeUTC()
	postResponse = []*model.PostResponse{
		{
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
		}, {
			ID:                    2,
			Title:                 "test title",
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
		},
	}
	actorResponse = &model.UserResponse{
		ID:       1,
		Name:     sql.NullString{String: "John Doe", Valid: true},
		PhotoURL: sql.NullString{String: "sample", Valid: true},
		Role:     sql.NullInt64{Int64: 99, Valid: true},
		Regency:  sql.NullString{String: "regency", Valid: true},
		District: sql.NullString{String: "district", Valid: true},
		Village:  sql.NullString{String: "village", Valid: true},
		RW:       sql.NullString{String: "rw", Valid: true},
	}
	metadataResponse = helper.SetPointerInt64(2)
	commentResponse  = &model.CommentResponse{
		ID:         1,
		Comment:    "comment",
		UserPostID: 1,
		CreatedAt:  current,
		UpdatedAt:  current,
		CreatedBy:  1,
		UpdatedBy:  1,
	}
	comment = &model.Comment{
		ID:         commentResponse.ID,
		UserPostID: commentResponse.UserPostID,
		Text:       commentResponse.Comment,
		CreatedAt:  commentResponse.CreatedAt,
		UpdatedAt:  commentResponse.UpdatedAt,
		CreatedBy:  actorResponse,
		UpdatedBy:  actorResponse,
	}
	totalComment     = helper.SetPointerInt64(1)
	userPostResponse = []*model.UserPostResponse{
		{
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
		}, {
			ID:                    2,
			Title:                 "test title",
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
		},
	}
	usecaseResponse = &model.UserPostWithMetadata{
		Data: userPostResponse,
		Metadata: &model.Metadata{
			Page:      1,
			TotalPage: 1,
			Total:     2,
		},
	}
)

var GetListUserPostData = []GetListUserPost{
	{
		Description:            "success_get_list_user_post",
		UsecaseParams:          requestUsecase,
		GetUserPostParams:      userPostParams,
		GetMetadataParams:      userPostParams,
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
		Description:            "failed_get_user_post",
		UsecaseParams:          requestUsecase,
		GetUserPostParams:      userPostParams,
		GetMetadataParams:      userPostParams,
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
			Error:  errors.New("failed_get_user_posts"),
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
			Error:  errors.New("failed_get_user_posts"),
		},
	}, {
		Description:            "failed_get_actor_created",
		UsecaseParams:          requestUsecase,
		GetUserPostParams:      userPostParams,
		GetMetadataParams:      userPostParams,
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
			Result: commentResponse,
			Error:  nil,
		},
		ResponseGetTotalComment: ResponseGetTotalComment{
			Result: totalComment,
			Error:  nil,
		},
		ResponseUsecase: ResponseUsecase{
			Result: nil,
			Error:  errors.New("failed_get_actor"),
		},
	}, {
		Description:            "failed_get_comment",
		UsecaseParams:          requestUsecase,
		GetUserPostParams:      userPostParams,
		GetMetadataParams:      userPostParams,
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
		Description:            "failed_get_metadata",
		UsecaseParams:          requestUsecase,
		GetUserPostParams:      userPostParams,
		GetMetadataParams:      userPostParams,
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
		Description:            "failed_check_is_liked",
		UsecaseParams:          requestUsecase,
		GetUserPostParams:      userPostParams,
		GetMetadataParams:      userPostParams,
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
			Result: nil,
			Error:  errors.New("something_went_wrong"),
		},
	},
}

func ListUserPostDescription() []string {
	var arr = []string{}
	for _, data := range GetListUserPostData {
		arr = append(arr, data.Description)
	}
	return arr
}
