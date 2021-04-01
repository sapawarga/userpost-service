package model

type UserPostRequest struct {
	ActivityName *string
	Username     *string
	Category     *string
	Status       *int64
	Offset       *int64
	Limit        *int64
	SortBy       *string
	OrderBy      *string
}

type IsLikedByUser struct {
	Type   string
	UserID int64
}
