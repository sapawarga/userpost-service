package model

import (
	"time"
)

type UserPostResponse struct {
	ID                    int64                    `json:"id"`
	Title                 string                   `json:"text"`
	Tag                   *string                  `json:"tags,omitempty"`
	ImagePath             string                   `json:"image_path_full,omitempty"`
	Images                []map[string]interface{} `json:"images"`
	LastUserPostCommentID *int64                   `json:"last_user_post_comment_id,omitempty"`
	LastComment           *Comment                 `json:"last_comment,omitempty"`
	LikesCount            int64                    `json:"likes_count"`
	IsLiked               bool                     `json:"is_liked"`
	CommentCounts         int64                    `json:"comment_counts"`
	Status                int64                    `json:"status"`
	StatusLabel           string                   `json:"status_label"`
	Actor                 *Actor                   `json:"user"`
	CreatedAt             time.Time                `json:"created_at"`
	UpdatedAt             time.Time                `json:"updated_at"`
}

type UserPostWithMetadata struct {
	Data     []*UserPostResponse `json:"data"`
	Metadata *Metadata           `json:"metadata"`
}

type CommentWithMetadata struct {
	Data     []*Comment `json:"data"`
	Metadata *Metadata  `json:"metadata"`
}

type Metadata struct {
	Page  int64 `json:"page"`
	Total int64 `json:"total"`
}

type Comment struct {
	ID         int64     `json:"id"`
	UserPostID int64     `json:"user_post_id"`
	Text       string    `json:"text"`
	User       *Actor    `json:"user"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedBy  int64     `json:"created_by"`
	UpdatedBy  int64     `json:"updated_by"`
}

type Actor struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	PhotoURL  string `json:"photo_url_full,omitempty"`
	Role      int64  `json:"role,omitempty"`
	RoleLabel string `json:"role_label,omitempty"`
	Regency   string `json:"kabkota,omitempty"`
	District  string `json:"kecamatan,omitempty"`
	Village   string `json:"kelurahan,omitempty"`
	RW        string `json:"rw,omitempty"`
}

var RoleLabel = map[int64]string{
	10:  "user",
	49:  "trainer",
	50:  "staffRW",
	60:  "staffKel",
	70:  "staffKec",
	80:  "staffKabKota",
	88:  "staffOPD",
	89:  "staffSaberhoax",
	90:  "staffProv",
	91:  "pimpinan",
	99:  "admin",
	100: "service_account",
}

var StatusLabel = map[int64]map[string]string{
	-1: {"en": "Status Deleted", "id": "Dihapus"},
	0:  {"en": "Inactived", "id": "Tidak Aktif"},
	10: {"en": "Actived", "id": "Aktif"},
}
