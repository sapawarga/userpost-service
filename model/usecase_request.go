package model

type GetListRequest struct {
	ActivityName *string
	Username     *string
	Category     *string
	Status       *int64
	Page         *int64
	Limit        *int64
	SortBy       *string
	OrderBy      *string
}
