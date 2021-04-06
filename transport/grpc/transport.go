package grpc

import (
	"context"

	"github.com/sapawarga/userpost-service/endpoint"
	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/usecase"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	transportUserPost "github.com/sapawarga/proto-file/userpost"
)

func MakeHandler(ctx context.Context, fs usecase.UsecaseI) transportUserPost.UserPostHandlerServer {
	userPostGetListHandler := kitgrpc.NewServer(
		endpoint.MakeGetListUserPost(ctx, fs),
		decodeGetListUserPost,
		encodeGetListUserPost,
	)

	return &grpcServer{
		userPostGetListHandler,
	}
}

func decodeGetListUserPost(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportUserPost.GetListUserPostRequest)

	return &endpoint.GetListUserPostRequest{
		ActivityName: helper.SetPointerString(req.GetActivityName()),
		Username:     helper.SetPointerString(req.GetUsername()),
		Category:     helper.SetPointerString(req.GetCategory()),
		Status:       helper.SetPointerInt64(req.GetStatus()),
		Page:         helper.SetPointerInt64(req.GetPage()),
		Limit:        helper.SetPointerInt64(req.GetLimit()),
		SortBy:       helper.SetPointerString(req.GetSortBy()),
		OrderBy:      helper.SetPointerString(req.GetOrderBy()),
	}, nil
}

func encodeGetListUserPost(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.UserPostWithMetadata)
	data := resp.Data
	metadata := resp.Metadata

	resultData := make([]*transportUserPost.UserPost, 0)

	for _, v := range data {
		result := &transportUserPost.UserPost{
			Id:                    v.ID,
			Title:                 v.Title,
			Tag:                   helper.GetStringFromPointer(v.Tag),
			ImagePath:             v.ImagePath,
			Images:                v.Images,
			LastUserPostCommentId: helper.GetInt64FromPointer(v.LastUserPostCommentID),
			LikesCount:            v.LikesCount,
			CommentCounts:         v.CommentCounts,
			Status:                v.Status,
			CreatedAt:             v.CreatedAt.String(),
			UpdatedAt:             v.UpdatedAt.String(),
		}
		if v.Actor != nil {
			actor := &transportUserPost.Actor{
				Id:       v.Actor.ID,
				Name:     v.Actor.Name.String,
				PhotoUrl: v.Actor.PhotoURL.String,
				Role:     v.Actor.Role.Int64,
				Regency:  v.Actor.Regency,
				District: v.Actor.District,
				Village:  v.Actor.Village,
				Rw:       v.Actor.RW.String,
			}
			result.Actor = actor
		}
		if v.LastComment != nil {
			comment := &transportUserPost.Comment{
				Id:            v.LastComment.ID,
				Comment:       v.LastComment.Comment,
				ActorId:       v.LastComment.ActorID,
				ActorName:     v.LastComment.ActorName,
				ActorPhotoUrl: v.LastComment.ActorPhotoURL,
				RegencyName:   v.LastComment.RegencyName,
				DistrictName:  v.LastComment.DistrictName,
				VillageName:   v.LastComment.DistrictName,
				Rw:            v.LastComment.RW,
			}
			result.LastComment = comment
		}

		resultData = append(resultData, result)
	}

	meta := &transportUserPost.Metadata{
		Page:      metadata.Page,
		Total:     metadata.Total,
		TotalPage: metadata.TotalPage,
	}

	return &transportUserPost.GetListUserPostResponse{
		Data:     resultData,
		Metadata: meta,
	}, nil
}
