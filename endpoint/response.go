package endpoint

import "github.com/sapawarga/userpost-service/model"

type UserPostWithMetadata struct {
	Data     []*model.UserPostResponse `json:"items"`
	Metadata *Metadata                 `json:"_meta"`
}

type Metadata struct {
	PerPage     int64   `json:"perPage"`
	TotalPage   float64 `json:"pageCount"`
	CurrentPage int64   `json:"currentPage"`
	Total       int64   `json:"totalCount"`
}

type UserPostDetail struct {
	*model.UserPostResponse
}

type StatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CommentsResponse struct {
	Data     []*model.Comment `json:"items"`
	Metadata *Metadata        `json:"_meta"`
}
