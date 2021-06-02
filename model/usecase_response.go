package model

import (
	"time"
)

type UserPostResponse struct {
	ID                    int64     `json:"id"`
	Title                 string    `json:"title"`
	Tag                   *string   `json:"tags,omitempty"`
	ImagePath             string    `json:"image_path,omitempty"`
	Images                string    `json:"images"`
	LastUserPostCommentID *int64    `json:"last_user_post_comment_id,omitempty"`
	LastComment           *Comment  `json:"last_comment,omitempty"`
	LikesCount            int64     `json:"likes_count"`
	IsLiked               bool      `json:"is_liked"`
	CommentCounts         int64     `json:"comment_counts"`
	Status                int64     `json:"status"`
	Actor                 *Actor    `json:"actor"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type UserPostWithMetadata struct {
	Data     []*UserPostResponse `json:"data"`
	Metadata *Metadata           `json:"metadata"`
}

type Metadata struct {
	Page      int64 `json:"page"`
	TotalPage int64 `json:"total_page"`
	Total     int64 `json:"total"`
}

type Comment struct {
	ID         int64         `json:"id"`
	UserPostID int64         `json:"user_post_id"`
	Text       string        `json:"comment"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	CreatedBy  *UserResponse `json:"created_by"`
	UpdatedBy  *UserResponse `json:"updated_by"`
}

type Actor struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url,omitempty"`
	Role     int64  `json:"role,omitempty"`
	Regency  string `json:"regency,omitempty"`
	District string `json:"district,omitempty"`
	Village  string `json:"village,omitempty"`
	RW       string `json:"rw,omitempty"`
	Status   int64  `json:"status,omitempty"`
}
