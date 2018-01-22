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
	md := metadata.Pairs("Authorization", "Bearer f0c93cc4-5073-4ede-bee2-d230d2b0c7b6")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.CreatePost(ctx, &protos.ReqPostCreate{
		Title:      "Hello",
		Body:       "Test body",
		Categories: []string{},
		Tags:       []string{},
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
