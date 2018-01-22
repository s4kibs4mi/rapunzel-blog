package servers

import (
	"testing"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"context"
)

func TestPostServer_CreatePost(t *testing.T) {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	client := protos.NewPostServiceClient(conn)
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.CreatePost(ctx, &protos.ReqPostCreate{
		Title:      "Hello",
		Body:       "Test body",
		Categories: []string{"test", "boom", "blog"},
		Tags:       []string{"test"},
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
