package servers

import (
	"context"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"github.com/s4kibs4mi/rapunzel-blog/api"
)

type PostServer struct {
}

func NewPostServer() *PostServer {
	return &PostServer{}
}

func (s *PostServer) CreatePost(ctx context.Context, params *protos.ReqPostCreate) (*protos.ResPost, error) {
	return api.CreatePost(ctx, params)
}

func (s *PostServer) UpdatePost(ctx context.Context, params *protos.ReqPostUpdate) (*protos.ResPost, error) {
	return api.UpdatePost(ctx, params)
}

func (s *PostServer) DeletePost(ctx context.Context, params *protos.GetByID) (*protos.ResPostSuccess, error) {
	return api.DeletePost(ctx, params)
}

func (s *PostServer) ChangeStatus(ctx context.Context, params *protos.ReqPostChangeStatus) (*protos.ResPost, error) {
	return api.ChangePostStatus(ctx, params)
}

func (s *PostServer) GetPost(ctx context.Context, params *protos.GetByID) (*protos.ResPost, error) {
	return api.GetPost(ctx, params)
}

func (s *PostServer) FavouritePost(ctx context.Context, params *protos.GetByID) (*protos.ResPost, error) {
	return api.FavouritePost(ctx, params)
}

func (s *PostServer) GetPosts(ctx context.Context, params *protos.GetByQuery) (*protos.ResPostList, error) {
	return api.ListPosts(ctx, params)
}
