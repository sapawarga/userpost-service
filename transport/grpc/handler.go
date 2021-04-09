package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	transportUserPost "github.com/sapawarga/proto-file/userpost"
)

type grpcServer struct {
	userPostGetList   kitgrpc.Handler
	userPostGetDetail kitgrpc.Handler
	userPostCreateNew kitgrpc.Handler
}

func (g *grpcServer) GetListUserPost(ctx context.Context, in *transportUserPost.GetListUserPostRequest) (*transportUserPost.GetListUserPostResponse, error) {
	_, resp, err := g.userPostGetList.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*transportUserPost.GetListUserPostResponse), nil
}

func (g *grpcServer) GetDetailUserPost(ctx context.Context, in *transportUserPost.ByID) (*transportUserPost.UserPost, error) {
	_, resp, err := g.userPostGetDetail.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*transportUserPost.UserPost), nil
}

func (g *grpcServer) CreateNewPost(ctx context.Context, in *transportUserPost.CreateNewPostRequest) (*transportUserPost.StatusResponse, error) {
	_, resp, err := g.userPostCreateNew.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*transportUserPost.StatusResponse), nil
}
