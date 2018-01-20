package api

import (
	"context"
	pb "github.com/s4kibs4mi/rapunzel-blog/protos"
)

type UserServer struct {
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (u *UserServer) Register(ctx context.Context, params *pb.ReqRegistration) (*pb.ResRegistration, error) {
	return &pb.ResRegistration{}, nil
}

func (u *UserServer) Login(ctx context.Context, params *pb.ReqLogin) (*pb.ResLogin, error) {
	return &pb.ResLogin{}, nil
}

func (u *UserServer) Profile(ctx context.Context, params *pb.ReqProfile) (*pb.ResProfile, error) {
	return &pb.ResProfile{}, nil
}

func (u *UserServer) Update(ctx context.Context, params *pb.ReqUpdateUser) (*pb.ResUpdateUser, error) {
	return &pb.ResUpdateUser{}, nil
}

func (u *UserServer) ChangePassword(ctx context.Context, params *pb.ReqChangePassword) (*pb.ResChangePassword, error) {
	return &pb.ResChangePassword{}, nil
}

func (u *UserServer) ChangeStatus(ctx context.Context, params *pb.ReqChangeUserStatus) (*pb.ResChangeUserStatus, error) {
	return &pb.ResChangeUserStatus{}, nil
}

func (u *UserServer) ChangeType(ctx context.Context, params *pb.ReqChangeUserType) (*pb.ResChangeUserType, error) {
	return &pb.ResChangeUserType{}, nil
}

func (u *UserServer) Logout(ctx context.Context, params *pb.ReqUserLogout) (*pb.ResUserLogout, error) {
	return &pb.ResUserLogout{}, nil
}

func (u *UserServer) ResetPasswordRequest(ctx context.Context, params *pb.ReqResetPasswordRequest) (*pb.ResResetPasswordRequest, error) {
	return &pb.ResResetPasswordRequest{}, nil
}

func (u *UserServer) ResetPassword(ctx context.Context, params *pb.ReqResetPassword) (*pb.ResResetPassword, error) {
	return &pb.ResResetPassword{}, nil
}
