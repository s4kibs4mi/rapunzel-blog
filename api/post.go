package api

import (
	"context"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"github.com/s4kibs4mi/rapunzel-blog/security"
)

func CreatePost(ctx context.Context, params *protos.ReqPostCreate) (*protos.ResPost, error) {
	if !security.IsAuthenticated(ctx) {
		return nil, security.GetUnauthenticatedError()
	}

	return &protos.ResPost{}, nil
}
