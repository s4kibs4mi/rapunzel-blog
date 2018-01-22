package security

import (
	"golang.org/x/crypto/bcrypt"
	"context"
	"google.golang.org/grpc/metadata"
	"github.com/s4kibs4mi/rapunzel-blog/auth"
	"fmt"
)

/**
 * := Coded with love by Sakib Sami on 20/1/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func NewBCryptPassword(plain string) string {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(encrypted)
}

func CheckBCryptPassword(encryptedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(plainPassword))
	return err == nil
}

func ReadUserIDFromContext(ctx context.Context) string {
	fmt.Println(ctx)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md[auth.UserID]) > 0 {
		return md[auth.UserID][0]
	}
	return ""
}
