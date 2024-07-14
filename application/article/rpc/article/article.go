// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package article

import (
	"context"

	"AlumniCircle/application/article/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ArticleDeleteRequest  = pb.ArticleDeleteRequest
	ArticleDeleteResponse = pb.ArticleDeleteResponse
	ArticleDetailRequest  = pb.ArticleDetailRequest
	ArticleDetailResponse = pb.ArticleDetailResponse
	ArticleItem           = pb.ArticleItem
	ArticlesRequest       = pb.ArticlesRequest
	ArticlesResponse      = pb.ArticlesResponse
	PublishRequest        = pb.PublishRequest
	PublishResponse       = pb.PublishResponse

	Article interface {
		Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
		Articles(ctx context.Context, in *ArticlesRequest, opts ...grpc.CallOption) (*ArticlesResponse, error)
		ArticleDelete(ctx context.Context, in *ArticleDeleteRequest, opts ...grpc.CallOption) (*ArticleDeleteResponse, error)
		ArticleDetail(ctx context.Context, in *ArticleDetailRequest, opts ...grpc.CallOption) (*ArticleDetailResponse, error)
	}

	defaultArticle struct {
		cli zrpc.Client
	}
)

func NewArticle(cli zrpc.Client) Article {
	return &defaultArticle{
		cli: cli,
	}
}

func (m *defaultArticle) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.Publish(ctx, in, opts...)
}

func (m *defaultArticle) Articles(ctx context.Context, in *ArticlesRequest, opts ...grpc.CallOption) (*ArticlesResponse, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.Articles(ctx, in, opts...)
}

func (m *defaultArticle) ArticleDelete(ctx context.Context, in *ArticleDeleteRequest, opts ...grpc.CallOption) (*ArticleDeleteResponse, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.ArticleDelete(ctx, in, opts...)
}

func (m *defaultArticle) ArticleDetail(ctx context.Context, in *ArticleDetailRequest, opts ...grpc.CallOption) (*ArticleDetailResponse, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.ArticleDetail(ctx, in, opts...)
}
