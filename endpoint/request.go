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

func Validate(in *CreateNewPostRequest) error {
	return validation.ValidateStruct(in,
		validation.Field(&in.Title, validation.Required, validation.Length(0, 10)),
		validation.Field(in.Images, validation.Required, validation.By(validationImages(in.Images))),
		validation.Field(&in.Tags, validation.Required),
		validation.Field(&in.Status, validation.Required, validation.In(helper.ACTIVED, helper.DELETED, helper.INACTIVED)),
	)
}

func validationImages(in []*Image) validation.RuleFunc {
	return func(value interface{}) error {
		var err error
		images := value.([]*Image)
		for _, v := range images {
			err = validation.ValidateStruct(v,
				validation.Field(v.Path, is.URL),
			)
		}

		return err
	}
}
