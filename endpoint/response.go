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
