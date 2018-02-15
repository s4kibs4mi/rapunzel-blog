package servers

import (
	"testing"
	"google.golang.org/grpc"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"os"
)

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

var conn *grpc.ClientConn
var client protos.RapunzelBlogServiceClient
var err error

func init() {
	conn, err = grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Couldn't initialize connection, ", err)
		os.Exit(-1)
	}
	client = protos.NewRapunzelBlogServiceClient(conn)
}

func TestRapunzelBlogServer_Register(t *testing.T) {
	resp, e := client.Register(context.Background(), &protos.ReqRegistration{
		Name:     "Nur",
		Email:    "nur@sakib.ninja",
		Username: "nur",
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

func TestRapunzelBlogServer_Login(t *testing.T) {
	resp, e := client.Login(context.Background(), &protos.ReqLogin{
		Username: "s4kibs4mi",
		Password: "123456789",
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

func TestRapunzelBlogServer_CreatePost(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.CreatePost(ctx, &protos.ReqPostCreate{
		Title:      "Hello",
		Body:       "Test body",
		Categories: []string{"test", "jally", "blog"},
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

func TestRapunzelBlogServer_UpdatePost(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.UpdatePost(ctx, &protos.ReqPostUpdate{
		Id:         "5a662858b34db60518737db1",
		Title:      "Hello",
		Body:       "Test body",
		Categories: []string{"test", "jally", "blog"},
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

func TestRapunzelBlogServer_GetPosts(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.GetPosts(ctx, &protos.GetByQuery{
		Query: []*protos.Query{
			{
				Field: "status",
				Value: "saved",
			},
			{
				Field: "categories",
				Value: "jally",
			},
		},
	})
	if e != nil {
		t.Error(e)
		return
	}
	for _, p := range resp.Posts {
		fmt.Println(p)
	}
}

func TestRapunzelBlogServer_GetPost(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.GetPost(ctx, &protos.GetByID{
		Id: "5a662802b34db604fb5dbc89",
	})
	if e != nil {
		t.Error(e)
		return
	}
	if resp.Post == nil {
		t.Error(resp.Errors)
		return
	}
	fmt.Println(resp.Post)
}

func TestRapunzelBlogServer_FavouritePost(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.FavouritePost(ctx, &protos.GetByID{
		Id: "5a662802b34db604fb5dbc89",
	})
	if e != nil {
		t.Error(e)
		return
	}
	if resp.Post == nil {
		t.Error(resp.Errors)
		return
	}
	fmt.Println(resp.Post)
}

func TestRapunzelBlogServer_DeletePost(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.DeletePost(ctx, &protos.GetByID{
		Id: "5a662858b34db60518737db1",
	})
	if e != nil {
		t.Error(e)
		return
	}
	if !resp.Success {
		t.Error(resp.Errors)
		return
	}
	fmt.Println(resp.Success)
}

func TestRapunzelBlogServer_ChangeStatus(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.ChangePostStatus(ctx, &protos.ReqPostChangeStatus{
		Id:        "5a662802b34db604fb5dbc89",
		NewStatus: "published",
	})
	if e != nil {
		t.Error(e)
		return
	}
	if resp.Post == nil {
		t.Error(resp.Errors)
		return
	}
	fmt.Println(resp.Post)
}

func TestRapunzelBlogServer_CreateComment(t *testing.T) {
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

func TestRapunzelBlogServer_GetComments(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.GetComments(ctx, &protos.GetByQuery{

	})
	if e != nil {
		t.Error(e)
		return
	}
	for _, c := range resp.Comments {
		fmt.Println(c)
	}
}

func TestRapunzelBlogServer_GetComment(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.GetComment(ctx, &protos.GetByID{
		Id: "5a67227e29c4463ab740dce8",
	})
	if e != nil {
		t.Error(e)
		return
	}
	if resp.Comment == nil {
		t.Error(resp.Errors)
		return
	}
	fmt.Println(resp.Comment)
}

func TestRapunzelBlogServer_GetProfile(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.Profile(ctx, &protos.ReqProfile{
		UserID: "5a854a6329c4467ceb3fc892",
	})
	if e != nil {
		t.Error(e)
		return
	}
	if resp.User == nil {
		t.Error(resp.Errors)
		return
	}
	fmt.Println(resp.User)
}

func TestRapunzelBlogServer_ChangeUserStatus(t *testing.T) {
	md := metadata.Pairs("Authorization", "Bearer 13ca3c5f-ec6d-4914-a0a8-98b3d681a05b")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, e := client.ChangeStatus(ctx, &protos.ReqChangeUserStatus{
		UserID:    "5a854a6329c4467ceb3fc892",
		NewStatus: "verified",
	})
	if e != nil {
		t.Error(e)
		return
	}
	if resp.Success == false {
		t.Error(resp.Errors)
		return
	}
	fmt.Println(resp.Success)
}
