package servers

import (
	"context"
	"fmt"
	"testing"

	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func TestUserServer_Register(t *testing.T) {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	client := protos.NewUserServiceClient(conn)
	resp, e := client.Register(context.Background(), &protos.ReqRegistration{
		Name:     "Sakib Sami",
		Email:    "root@sakib.ninja",
		Username: "s4kibs4mi",
		Password: "12345678",
		Details:  "Hello World",
	})
	if e != nil {
		t.Error(e)
		return
	}
	if resp.Errors != nil {
		t.Error(resp.Errors)
		return
	}
	fmt.Println(resp)
}

func TestUserServer_Login(t *testing.T) {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	client := protos.NewUserServiceClient(conn)
	resp, e := client.Login(context.Background(), &protos.ReqLogin{
		Username: "s4kibs4mi",
		Password: "12345678",
	})
	if e != nil {
		t.Error(e)
		return
	}
	if resp.Errors != nil {
		t.Error(resp.Errors)
		return
	}
	fmt.Println(resp)
}

func TestUserProfile(t *testing.T) {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	token := "c6502710-67e7-404f-8149-6f5275cf5372"
	client := protos.NewUserServiceClient(conn)
	md := metadata.Pairs("Authorization", "Bearer "+token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	req := &protos.ReqProfile{
		AccessToken: token,
	}
	p, err := client.Profile(ctx, req)
	if err != nil {
		t.Error(err)
	}
	if p.User == nil {
		t.Error("Failed to fetch profile")
	}
}
