package servers

import (
	"testing"
	"google.golang.org/grpc"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"context"
	"fmt"
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
