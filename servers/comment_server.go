package servers

import (
	"context"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"github.com/s4kibs4mi/rapunzel-blog/api"
)

type CommentServer struct {
}

func NewCommentServer() *CommentServer {
	return &CommentServer{}
}

func (s *CommentServer) CreateComment(ctx context.Context, params *protos.ReqCommentCreate) (*protos.ResComment, error) {
	return api.CreateComment(ctx, params)
}

func (s *CommentServer) UpdateComment(ctx context.Context, params *protos.ReqCommentUpdate) (*protos.ResComment, error) {
	return nil, nil
}

func (s *CommentServer) DeleteComment(ctx context.Context, params *protos.GetByID) (*protos.ResCommentSuccess, error) {
	return nil, nil
}

func (s *CommentServer) GetComment(ctx context.Context, params *protos.GetByID) (*protos.ResComment, error) {
	return api.GetComment(ctx, params)
}

func (s *CommentServer) GetComments(ctx context.Context, params *protos.GetByQuery) (*protos.ResCommentList, error) {
	return api.GetComments(ctx, params)
}
