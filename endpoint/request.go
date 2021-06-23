package endpoint

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sapawarga/userpost-service/helper"
)

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

type CreateNewPostRequest struct {
	Title  *string  `json:"title"`
	Images []*Image `json:"images"`
	Tags   *string  `json:"tags"`
	Status *int64   `json:"status"`
}

type Image struct {
	Path string `json:"path"`
}

type CreateCommentRequest struct {
	UserPostID int64  `json:"user_post_id"`
	Text       string `json:"text"`
	Status     *int64 `json:"status"`
}

func Validate(in interface{}) error {
	var err error
	if obj, ok := in.(*CreateNewPostRequest); ok {
		err = validation.ValidateStruct(in,
			validation.Field(&obj.Title, validation.Required, validation.Length(10, 0)),
			validation.Field(&obj.Images, validation.Required, validation.By(validationImages(obj.Images))),
			validation.Field(&obj.Tags, validation.Required),
			validation.Field(&obj.Status, validation.Required, validation.In(helper.ACTIVED, helper.DELETED, helper.INACTIVED)),
		)
	} else if obj, ok := in.(*CreateCommentRequest); ok {
		err = validation.ValidateStruct(in,
			validation.Field(&obj.UserPostID, validation.Required),
			validation.Field(&obj.Text, validation.Required, validation.Length(10, 0)),
			validation.Field(&obj.Status, validation.Required, validation.In(helper.ACTIVED, helper.DELETED, helper.INACTIVED)),
		)
	}
	return err
}

func validationImages(in []*Image) validation.RuleFunc {
	return func(value interface{}) error {
		var err error
		images := value.([]*Image)
		for _, v := range images {
			err = validation.ValidateStruct(v,
				validation.Field(&v.Path, is.URL),
			)
		}
		return err
	}
}

var ordering = map[int]string{
	0: "ASC",
	1: "DESC",
}

func isOrderValid(val string) bool {
	for i := range ordering {
		if ordering[i] == val {
			return true
		}
	}

	return false
}
