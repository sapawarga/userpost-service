package model

import (
	"database/sql"
)

type PostResponse struct {
	ID                    int64          `db:"id"`
	Title                 string         `db:"text"`
	Tag                   sql.NullString `db:"tags"`
	ImagePath             sql.NullString `db:"image_path"`
	Images                sql.NullString `db:"images"`
	LastUserPostCommentID sql.NullInt64  `db:"last_user_post_comment_id"`
	LikesCount            int64          `db:"likes_count"`
	CommentCounts         int64          `db:"comments_count"`
	Status                int64          `db:"status"`
	CreatedBy             sql.NullInt64  `db:"created_by"`
	UpdatedBy             sql.NullInt64  `db:"updated_by"`
	CreatedAt             int64          `db:"created_at"`
	UpdatedAt             int64          `db:"updated_at"`
}

type UserResponse struct {
	ID       int64          `db:"id"`
	Name     sql.NullString `db:"name"`
	PhotoURL sql.NullString `db:"photo_url"`
	Role     sql.NullInt64  `db:"role"`
	Regency  sql.NullString `db:"regency_name"`
	District sql.NullString `db:"district_name"`
	Village  sql.NullString `db:"village_name"`
	RW       sql.NullString `db:"rw"`
	Status   int64          `db:"status"`
}

type LikeResponse struct {
	ID       int64  `db:"id"`
	Type     string `db:"type"`
	EntityID int64  `db:"entity_id"`
}

type CommentResponse struct {
	ID         int64         `db:"id"`
	UserPostID sql.NullInt64 `db:"user_post_id"`
	Comment    string        `db:"text"`
	Status     int64         `db:"status"`
	CreatedAt  int64         `db:"created_at"`
	UpdatedAt  int64         `db:"updated_at"`
	CreatedBy  sql.NullInt64 `db:"created_by"`
	UpdatedBy  sql.NullInt64 `db:"updated_by"`
}
