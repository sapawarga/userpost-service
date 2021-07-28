package testcases

import (
	"errors"

	"github.com/sapawarga/userpost-service/lib/constant"
	"github.com/sapawarga/userpost-service/lib/convert"
	"github.com/sapawarga/userpost-service/model"
)

var updateUserPost = &model.UpdatePostRequest{
	ID:     1,
	Status: convert.SetPointerInt64(constant.ACTIVED),
	Title:  convert.SetPointerString("Update Description"),
}

type UpdateUserPostDetail struct {
	Description                string
	UsecaseRequest             *model.UpdatePostRequest
	GetDetailParam             int64
	UpdateUserPostParam        *model.UpdatePostRequest
	ResponseGetDetailRepo      ResponseGetDetailUserPost
	ResponseUpdateUserPostRepo error
	ResponseUsecase            error
}

var UpdateUserPostDetailData = []UpdateUserPostDetail{
	{
		Description:         "succes_update_detail",
		UsecaseRequest:      updateUserPost,
		GetDetailParam:      1,
		UpdateUserPostParam: updateUserPost,
		ResponseGetDetailRepo: ResponseGetDetailUserPost{
			Error: nil,
		},
		ResponseUpdateUserPostRepo: nil,
		ResponseUsecase:            nil,
	}, {
		Description:         "failed_get_detail",
		UsecaseRequest:      updateUserPost,
		GetDetailParam:      1,
		UpdateUserPostParam: updateUserPost,
		ResponseGetDetailRepo: ResponseGetDetailUserPost{
			Error: errors.New("failed_get_detail"),
		},
		ResponseUpdateUserPostRepo: nil,
		ResponseUsecase:            errors.New("failed_get_detail"),
	}, {
		Description:         "failed_update_detail",
		UsecaseRequest:      updateUserPost,
		GetDetailParam:      1,
		UpdateUserPostParam: updateUserPost,
		ResponseGetDetailRepo: ResponseGetDetailUserPost{
			Error: nil,
		},
		ResponseUpdateUserPostRepo: errors.New("failed_update_detail"),
		ResponseUsecase:            errors.New("failed_update_detail"),
	},
}

func UpdateUserPostDetailDescription() []string {
	var arr = []string{}
	for _, data := range UpdateUserPostDetailData {
		arr = append(arr, data.Description)
	}
	return arr
}
