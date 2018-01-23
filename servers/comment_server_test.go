package servers

import (
	"testing"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"context"
)

func TestCommentServer_CreateComment(t *testing.T) {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	client := protos.NewCommentServiceClient(conn)
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.CreateComment(ctx, &protos.ReqCommentCreate{
		Title:  "Hello",
		Body:   "Test Comment",
		PostId: "5a662802b34db604fb5dbc89",
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
