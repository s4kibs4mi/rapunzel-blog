package api

import (
	"testing"
	"google.golang.org/grpc"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"context"
	"fmt"
)

func TestUserServer_Register(t *testing.T) {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	client := protos.NewUserServiceClient(conn)
	resp, e := client.Register(context.Background(), &protos.ReqRegistration{
		Name: "Sakib",
	})
	if e != nil {
		t.Error(e)
	}
	fmt.Println(resp)
}
