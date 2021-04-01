package model

import (
	"database/sql"
	"time"
)

type PostResponse struct {
	ID                    int64          `db:"id"`
	Title                 string         `db:"text"`
	Tag                   sql.NullString `db:"tags"`
	ImagePath             sql.NullString `db:"image_path"`
	Images                sql.NullString `db:"images"`
	LastUserPostCommentID sql.NullInt64  `db:"last_user_post_comment_id"`
	LikesCount            int64          `db:"likes_count"`
	CommentCounts         int64          `db:"comment_counts"`
	Status                int64          `db:"status"`
	CreatedBy             sql.NullInt64  `db:"created_by"`
	UpdatedBy             sql.NullInt64  `db:"updated_by"`
	CreatedAt             time.Time      `db:"created_at"`
	UpdatedAt             time.Time      `db:"updated_at"`
}

type UserResponse struct {
	ID       int64          `db:"id"`
	Name     sql.NullString `db:"name"`
	PhotoURL sql.NullString `db:"photo_url"`
	Role     sql.NullInt64  `db:"role"`
	Regency  string         `db:"regency_name"`
	District string         `db:"district_name"`
	Village  string         `db:"village_name"`
	RW       sql.NullString `db:"rw"`
}

type LikeResponse struct {
	ID       int64  `db:"id"`
	Type     string `db:"type"`
	EntityID int64  `db:"entity_id"`
}

type CommentResponse struct {
	ID            int64  `db:"id"`
	Comment       string `db:"comment"`
	ActorID       int64  `db:"actor_id"`
	ActorName     string `db:"actor_name"`
	ActorPhotoURL string `db:"actor_photo_url"`
	RegencyName   string `db:"regency_name"`
	DistrictName  string `db:"district_name"`
	VillageName   string `db:"village_name"`
	RW            string `db:"rw"`
}
