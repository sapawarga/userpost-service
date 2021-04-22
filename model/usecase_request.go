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
	Title        string
	ImagePathURL string
	Images       string
	Tags         *string
	Status       int64
}

type UpdatePostRequest struct {
	ID     int64
	Status *int64
	Title  *string
}

type ActorFromContext struct {
	m map[string]interface{}
}

func (ac *ActorFromContext) Get(key string) interface{} {
	return ac.m[key]
}
