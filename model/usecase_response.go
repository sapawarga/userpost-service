package model

import (
	"time"
)

type UserPostResponse struct {
	ID                    int64
	Title                 string
	Tag                   *string
	ImagePath             string
	Images                string
	LastUserPostCommentID *int64
	LastComment           *CommentResponse
	LikesCount            int64
	IsLiked               bool
	CommentCounts         int64
	Status                int64
	Actor                 *UserResponse
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type UserPostWithMetadata struct {
	Data     []*UserPostResponse
	Metadata *Metadata
}

type Metadata struct {
	Page      int64
	TotalPage int64
	Total     int64
}
