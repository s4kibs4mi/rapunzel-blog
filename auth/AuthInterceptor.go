package auth

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"strings"
	"github.com/s4kibs4mi/rapunzel-blog/storage"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

const (
	AuthorizationKey = "authorization"
	BearerKey        = "Bearer"
	UserID           = "user_id"
)

func isInAuthorizationScope(method string) bool {
	//if strings.HasSuffix(method, "Login") {
	//	return false
	//}
	if strings.HasSuffix(method, "Register") {
		return false
	}
	return true
}

func UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if isInAuthorizationScope(info.FullMethod) {
		sessionStorage := storage.NewSessionStorage()
		userStorage := storage.NewUserStorage()

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			authorization := md[AuthorizationKey]
			if len(authorization) > 0 && strings.HasPrefix(authorization[0], BearerKey) {
				token := strings.TrimSpace(strings.TrimPrefix(authorization[0], BearerKey))
				fmt.Println("Token : ", token)
				session := sessionStorage.FindSessionByAccessToken(token)
				if session == nil {
					return handler(nil, status.Errorf(codes.Unauthenticated, "authentication required"))
				}
				fmt.Println(session)
				u := userStorage.FindByID(session.UserID.String())
				if u == nil {
					return handler(nil, status.Errorf(codes.Unauthenticated, "authentication required"))
				}
				fmt.Println(u)
				nCtx := context.WithValue(ctx, UserID, u.ID)
				return handler(nCtx, req)
			}
			return handler(nil, grpc.Errorf(codes.Unauthenticated, "authentication required"))
		}
		return handler(nil, grpc.Errorf(codes.Unauthenticated, "authentication required"))
	}
	return handler(ctx, req)
}
