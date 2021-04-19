package endpoint

import "github.com/sapawarga/userpost-service/model"

type UserPostWithMetadata struct {
	Data     []*model.UserPostResponse `json:"data"`
	Metadata *Metadata                 `json:"metadata"`
}

type Metadata struct {
	Page      int64 `json:"page"`
	TotalPage int64 `json:"total_page"`
	Total     int64 `json:"total"`
}

type UserPostDetail struct {
	*model.UserPostResponse
}

type StatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CommentsResponse struct {
	Data []*model.Comment `json:"data"`
}
