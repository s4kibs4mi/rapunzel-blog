package security

import (
	"github.com/s4kibs4mi/rapunzel-blog/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"context"
	"github.com/s4kibs4mi/rapunzel-blog/storage"
	"gopkg.in/mgo.v2/bson"
)

func HasLoginPermissions(u *models.User) bool {
	if u.UserType == models.UserTypeGhost ||
		u.UserStatus == models.UserStatusRegistered ||
		u.UserStatus == models.UserStatusBlocked {
		return false
	}
	return true
}

func HasPostWritePermission(ctx context.Context, p models.Post) bool {
	if ctx == nil {
		return false
	}
	if HasPermissionAsParent(ctx) {
		return true
	}
	userStorage := storage.NewUserStorage()
	u := userStorage.FindByID(bson.ObjectIdHex(ReadUserIDFromContext(ctx)))
	return u.ID == p.UserID
}

func HasPermissionAsParent(ctx context.Context) bool {
	if ctx == nil {
		return false
	}
	userStorage := storage.NewUserStorage()
	u := userStorage.FindByID(bson.ObjectIdHex(ReadUserIDFromContext(ctx)))
	return u.UserType == models.UserTypeParent
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

func GetUnauthorisedError() error {
	return status.Errorf(codes.PermissionDenied, "Unauthorized user")
}
