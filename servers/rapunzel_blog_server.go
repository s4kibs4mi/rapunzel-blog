package servers

import (
	"context"
	pb "github.com/s4kibs4mi/rapunzel-blog/protos"
	"github.com/s4kibs4mi/rapunzel-blog/api"
)

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

type RapunzelBlogServer struct {
}

func NewRapunzelBlogServer() *RapunzelBlogServer {
	return &RapunzelBlogServer{}
}

func (s *RapunzelBlogServer) Register(ctx context.Context, params *pb.ReqRegistration) (*pb.ResRegistration, error) {
	return api.Register(ctx, params)
}

func (s *RapunzelBlogServer) Login(ctx context.Context, params *pb.ReqLogin) (*pb.ResLogin, error) {
	return api.Login(ctx, params)
}

func (s *RapunzelBlogServer) Profile(ctx context.Context, params *pb.ReqProfile) (*pb.ResProfile, error) {
	return api.GetProfile(ctx, params)
}

func (s *RapunzelBlogServer) Update(ctx context.Context, params *pb.ReqUpdateUser) (*pb.ResUpdateUser, error) {
	return &pb.ResUpdateUser{}, nil
}

func (s *RapunzelBlogServer) ChangePassword(ctx context.Context, params *pb.ReqChangePassword) (*pb.ResChangePassword, error) {
	return &pb.ResChangePassword{}, nil
}

func (s *RapunzelBlogServer) ChangeStatus(ctx context.Context, params *pb.ReqChangeUserStatus) (*pb.ResChangeUserStatus, error) {
	return &pb.ResChangeUserStatus{}, nil
}

func (s *RapunzelBlogServer) ChangeType(ctx context.Context, params *pb.ReqChangeUserType) (*pb.ResChangeUserType, error) {
	return &pb.ResChangeUserType{}, nil
}

func (s *RapunzelBlogServer) Logout(ctx context.Context, params *pb.ReqUserLogout) (*pb.ResUserLogout, error) {
	return &pb.ResUserLogout{}, nil
}

func (s *RapunzelBlogServer) ResetPasswordRequest(ctx context.Context, params *pb.ReqResetPasswordRequest) (*pb.ResResetPasswordRequest, error) {
	return &pb.ResResetPasswordRequest{}, nil
}

func (s *RapunzelBlogServer) ResetPassword(ctx context.Context, params *pb.ReqResetPassword) (*pb.ResResetPassword, error) {
	return &pb.ResResetPassword{}, nil
}

func (s *RapunzelBlogServer) CreatePost(ctx context.Context, params *pb.ReqPostCreate) (*pb.ResPost, error) {
	return api.CreatePost(ctx, params)
}

func (s *RapunzelBlogServer) UpdatePost(ctx context.Context, params *pb.ReqPostUpdate) (*pb.ResPost, error) {
	return api.UpdatePost(ctx, params)
}

func (s *RapunzelBlogServer) DeletePost(ctx context.Context, params *pb.GetByID) (*pb.ResPostSuccess, error) {
	return api.DeletePost(ctx, params)
}

func (s *RapunzelBlogServer) ChangePostStatus(ctx context.Context, params *pb.ReqPostChangeStatus) (*pb.ResPost, error) {
	return api.ChangePostStatus(ctx, params)
}

func (s *RapunzelBlogServer) GetPost(ctx context.Context, params *pb.GetByID) (*pb.ResPost, error) {
	return api.GetPost(ctx, params)
}

func (s *RapunzelBlogServer) FavouritePost(ctx context.Context, params *pb.GetByID) (*pb.ResPost, error) {
	return api.FavouritePost(ctx, params)
}

func (s *RapunzelBlogServer) GetPosts(ctx context.Context, params *pb.GetByQuery) (*pb.ResPostList, error) {
	return api.ListPosts(ctx, params)
}

func (s *RapunzelBlogServer) CreateComment(ctx context.Context, params *pb.ReqCommentCreate) (*pb.ResComment, error) {
	return api.CreateComment(ctx, params)
}

func (s *RapunzelBlogServer) UpdateComment(ctx context.Context, params *pb.ReqCommentUpdate) (*pb.ResComment, error) {
	return nil, nil
}

func (s *RapunzelBlogServer) DeleteComment(ctx context.Context, params *pb.GetByID) (*pb.ResCommentSuccess, error) {
	return nil, nil
}

func (s *RapunzelBlogServer) GetComment(ctx context.Context, params *pb.GetByID) (*pb.ResComment, error) {
	return api.GetComment(ctx, params)
}

func (s *RapunzelBlogServer) GetComments(ctx context.Context, params *pb.GetByQuery) (*pb.ResCommentList, error) {
	return api.GetComments(ctx, params)
}
