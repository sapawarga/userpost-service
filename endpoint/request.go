package endpoint

type GetListUserPostRequest struct {
	ActivityName *string `httpquery:"activity_name"`
	Username     *string `httpquery:"username"`
	Category     *string `httpquery:"category"`
	Status       *int64  `httpquery:"status"`
	Page         *int64  `httpquery:"page"`
	Limit        *int64  `httpquery:"limit"`
	SortBy       *string `httpquery:"sort_by"`
	OrderBy      *string `httpquery:"order_by"`
}

type GetByID struct {
	ID int64 `json:"id" httpquery:"id"`
}
