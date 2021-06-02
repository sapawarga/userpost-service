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

type CreateNewPostRequest struct {
	Title        string  `json:"title"`
	ImagePathURL string  `json:"image_path_url"`
	Images       string  `json:"images"`
	Tags         *string `json:"tags,omitempty"`
	Status       int64   `json:"status"`
}

type UpdatePostRequest struct {
	ID     int64   `json:"id"`
	Status *int64  `json:"status,omitempty"`
	Title  *string `json:"title,omitempty"`
}

type CreateCommentRequest struct {
	UserPostID int64  `json:"user_post_id"`
	Text       string `json:"comment"`
}

type ActorFromContext struct {
	Data map[string]interface{}
}

func (ac *ActorFromContext) Get(key string) interface{} {
	return ac.Data[key]
}
