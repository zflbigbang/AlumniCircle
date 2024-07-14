// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package follow

import (
	"context"

	"AlumniCircle/application/follow/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FansItem           = pb.FansItem
	FansListRequest    = pb.FansListRequest
	FansListResponse   = pb.FansListResponse
	FollowItem         = pb.FollowItem
	FollowListRequest  = pb.FollowListRequest
	FollowListResponse = pb.FollowListResponse
	FollowRequest      = pb.FollowRequest
	FollowResponse     = pb.FollowResponse
	UnFollowRequest    = pb.UnFollowRequest
	UnFollowResponse   = pb.UnFollowResponse

	Follow interface {
		// 关注
		Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error)
		// 取消关注
		UnFollow(ctx context.Context, in *UnFollowRequest, opts ...grpc.CallOption) (*UnFollowResponse, error)
		// 关注列表
		FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error)
		// 粉丝列表
		FansList(ctx context.Context, in *FansListRequest, opts ...grpc.CallOption) (*FansListResponse, error)
	}

	defaultFollow struct {
		cli zrpc.Client
	}
)

func NewFollow(cli zrpc.Client) Follow {
	return &defaultFollow{
		cli: cli,
	}
}

// 关注
func (m *defaultFollow) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error) {
	client := pb.NewFollowClient(m.cli.Conn())
	return client.Follow(ctx, in, opts...)
}

// 取消关注
func (m *defaultFollow) UnFollow(ctx context.Context, in *UnFollowRequest, opts ...grpc.CallOption) (*UnFollowResponse, error) {
	client := pb.NewFollowClient(m.cli.Conn())
	return client.UnFollow(ctx, in, opts...)
}

// 关注列表
func (m *defaultFollow) FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error) {
	client := pb.NewFollowClient(m.cli.Conn())
	return client.FollowList(ctx, in, opts...)
}

// 粉丝列表
func (m *defaultFollow) FansList(ctx context.Context, in *FansListRequest, opts ...grpc.CallOption) (*FansListResponse, error) {
	client := pb.NewFollowClient(m.cli.Conn())
	return client.FansList(ctx, in, opts...)
}
