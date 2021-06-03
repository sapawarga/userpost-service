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

type UserPostByMeRequest struct {
	ActorID int64
	*UserPostRequest
}

type CreateCommentRequestRepository struct {
	UserPostID int64
	Text       string
	ActorID    int64
	Status     int64
}

type CreateNewPostRequestRepository struct {
	Title        string
	ImagePathURL string
	Images       string
	Tags         *string
	Status       int64
	ActorID      int64
}

type AddOrRemoveLikeOnPostRequest struct {
	UserPostID int64
	ActorID    int64
	TypeEntity string
}
