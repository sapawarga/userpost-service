package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/userpost-service/lib/convert"
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
	testRandomString = convert.SetPointerString("random-test")
	statusNumber     = convert.SetPointerInt64(10)
	requestUsecase   = model.GetListRequest{
		ActivityName: testRandomString,
		Username:     testRandomString,
		Category:     testRandomString,
		Status:       statusNumber,
		Page:         convert.SetPointerInt64(1),
		Limit:        convert.SetPointerInt64(10),
		SortBy:       nil,
		OrderBy:      nil,
	}
	userPostParams = model.UserPostRequest{
		ActivityName: requestUsecase.ActivityName,
		Username:     requestUsecase.Username,
		Category:     requestUsecase.Category,
		Status:       requestUsecase.Status,
		Offset:       convert.SetPointerInt64(0),
		Limit:        requestUsecase.Limit,
		SortBy:       nil,
		OrderBy:      nil,
	}
	_, current   = convert.GetCurrentTimeUTC()
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
		PhotoURL: sql.NullString{String: "www.instagram.com/htm-medium=?p9878y2y3", Valid: true},
		Role:     sql.NullInt64{Int64: 99, Valid: true},
		Regency:  sql.NullString{String: "regency", Valid: true},
		District: sql.NullString{String: "district", Valid: true},
		Village:  sql.NullString{String: "village", Valid: true},
		RW:       sql.NullString{String: "rw", Valid: true},
		Status:   10,
	}
	actor = &model.Actor{
		ID:        1,
		Name:      "John Doe",
		PhotoURL:  "www.instagram.com/htm-medium=?p9878y2y3",
		Role:      99,
		RoleLabel: model.RoleLabel[int64(99)],
		Regency:   convert.SetPointerString("regency"),
		District:  convert.SetPointerString("district"),
		Village:   convert.SetPointerString("village"),
		RW:        convert.SetPointerString("rw"),
	}
	metadataResponse = convert.SetPointerInt64(2)
	commentResponse  = &model.CommentResponse{
		ID:         1,
		Comment:    "comment",
		UserPostID: sql.NullInt64{Int64: 1, Valid: true},
		CreatedAt:  current,
		UpdatedAt:  current,
		CreatedBy:  sql.NullInt64{Int64: 1, Valid: true},
		UpdatedBy:  sql.NullInt64{Int64: 1, Valid: true},
	}
	comment = &model.Comment{
		ID:         commentResponse.ID,
		UserPostID: commentResponse.UserPostID.Int64,
		Text:       commentResponse.Comment,
		CreatedAt:  commentResponse.CreatedAt,
		UpdatedAt:  commentResponse.UpdatedAt,
		User:       actor,
		CreatedBy:  1,
		UpdatedBy:  1,
	}
	totalComment     = convert.SetPointerInt64(1)
	userPostResponse = []*model.UserPostResponse{
		{
			ID:                    1,
			Title:                 "title",
			Tag:                   "tag",
			ImagePath:             "test",
			Images:                images,
			LastUserPostCommentID: convert.SetPointerInt64(1),
			LastComment:           comment,
			LikesCount:            0,
			IsLiked:               true,
			CommentCounts:         1,
			Status:                10,
			Actor:                 actor,
			CreatedAt:             current,
			UpdatedAt:             current,
		}, {
			ID:                    2,
			Title:                 "test title",
			Tag:                   "tag",
			ImagePath:             "test",
			Images:                images,
			LastUserPostCommentID: convert.SetPointerInt64(1),
			LastComment:           comment,
			LikesCount:            0,
			IsLiked:               true,
			CommentCounts:         1,
			Status:                10,
			Actor:                 actor,
			CreatedAt:             current,
			UpdatedAt:             current,
		},
	}
	usecaseResponse = &model.UserPostWithMetadata{
		Data: userPostResponse,
		Metadata: &model.Metadata{
			Page:  1,
			Total: 2,
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
