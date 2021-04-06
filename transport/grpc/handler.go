package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	transportUserPost "github.com/sapawarga/proto-file/userpost"
)

type grpcServer struct {
	userPostGetList kitgrpc.Handler
}

func (g *grpcServer) GetListUserPost(ctx context.Context, in *transportUserPost.GetListUserPostRequest) (*transportUserPost.GetListUserPostResponse, error) {
	_, resp, err := g.userPostGetList.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*transportUserPost.GetListUserPostResponse), nil
}
