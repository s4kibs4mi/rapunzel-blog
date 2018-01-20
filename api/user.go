package api

import (
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"context"
	pb "github.com/s4kibs4mi/rapunzel-blog/protos"
	"github.com/s4kibs4mi/rapunzel-blog/storage"
)

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func Register(ctx context.Context, params *pb.ReqRegistration) (*protos.ResRegistration, error) {
	data := storage.NewUserStorage()
	u := data.FindByUsername(params.Username)
	if u == nil {
		errDetails := &pb.ErrorDetails{
			Field:   "username",
			Details: []string{"username already exists."},
		}
		err := &pb.Error{
			ErrorDetails: []*pb.ErrorDetails{errDetails},
		}
		return &pb.ResRegistration{
			User:   nil,
			Errors: []*pb.Error{err},
		}, nil
	}
	return &pb.ResRegistration{}, nil
}
