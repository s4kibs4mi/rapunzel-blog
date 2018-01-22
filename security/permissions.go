package security

import (
	"github.com/s4kibs4mi/rapunzel-blog/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"context"
)

func HasLoginPermissions(u *models.User) bool {
	if u.UserType == models.UserTypeGhost ||
		u.UserStatus == models.UserStatusRegistered ||
		u.UserStatus == models.UserStatusBlocked {
		return false
	}
	return true
}

func IsAuthenticated(ctx context.Context) bool {
	if ctx == nil {
		return false
	}
	return true
}

func GetUnauthenticatedError() error {
	return status.Errorf(codes.Unauthenticated, "Authentication required")
}
