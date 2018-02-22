package api

import (
	"context"
	"github.com/s4kibs4mi/rapunzel-blog/security"
	pb "github.com/s4kibs4mi/rapunzel-blog/proto/defs"
	"github.com/satori/go.uuid"
	"net/http"
	"github.com/s4kibs4mi/rapunzel-blog/models"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/s4kibs4mi/rapunzel-blog/storage"
	"fmt"
)

func CreatePost(ctx context.Context, params *pb.ReqPostCreate) (*pb.ResPost, error) {
	if !security.IsAuthenticated(ctx) {
		return nil, security.GetUnauthenticatedError()
	}
	postStorage := storage.NewPostStorage()
	var validationErrorDetails []*pb.ErrorDetails

	// Validating Title
	var titleErrors []string
	titleLen := len(params.Title)
	if titleLen < 1 || titleLen > 100 {
		titleErrors = append(titleErrors, "Title length must be between 1 to 100")
	}
	if len(titleErrors) != 0 {
		tErr := pb.ErrorDetails{
			Field:   "Title",
			Details: titleErrors,
		}
		validationErrorDetails = append(validationErrorDetails, &tErr)
	}
	// Validating Body
	var bodyErrors []string
	bodyLen := len(params.Body)
	if bodyLen < 1 || bodyLen > 100000 {
		bodyErrors = append(bodyErrors, "Body length must be between 1 to 100000")
	}
	if len(bodyErrors) != 0 {
		bErr := pb.ErrorDetails{
			Field:   "Body",
			Details: bodyErrors,
		}
		validationErrorDetails = append(validationErrorDetails, &bErr)
	}
	if len(validationErrorDetails) != 0 {
		return &pb.ResPost{
			Post: nil,
			Errors: []*pb.Error{
				{
					ID:           uuid.NewV4().String(),
					Code:         http.StatusUnprocessableEntity,
					Title:        "Post Validation Error",
					Details:      "Required fields are not valid",
					ErrorDetails: validationErrorDetails,
				},
			},
		}, nil
	}

	p := models.Post{}
	p.ID = bson.NewObjectId()
	p.UserID = bson.ObjectIdHex(security.ReadUserIDFromContext(ctx))
	p.Title = params.Title
	p.Body = params.Body
	p.Categories = params.Categories
	p.Tags = params.Tags
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.Views = 0
	p.Favourites = 0
	p.Status = models.PostStatusSaved
	if postStorage.SavePost(&p) {
		return &pb.ResPost{
			Post: &pb.Post{
				Id:         p.ID.Hex(),
				Title:      p.Title,
				Body:       p.Body,
				Categories: p.Categories,
				Tags:       p.Tags,
				Status:     string(p.Status),
				Favourites: p.Favourites,
				Views:      p.Views,
				UpdatedAt:  p.UpdatedAt.String(),
				CreatedAt:  p.CreatedAt.String(),
			},
		}, nil
	}
	return &pb.ResPost{
		Post: nil,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusInternalServerError,
				Title:   "Something went wrong",
				Details: "Couldn't store the post",
			},
		},
	}, nil
}

func ListPosts(ctx context.Context, params *pb.GetByQuery) (*pb.ResPostList, error) {
	postStorage := storage.NewPostStorage()
	posts := postStorage.FindPostsByQuery(params.Query)
	var convertedPosts []*pb.Post
	for _, p := range posts {
		convertedPosts = append(convertedPosts, &pb.Post{
			Id:         p.ID.Hex(),
			Title:      p.Title,
			Body:       p.Body,
			Categories: p.Categories,
			Tags:       p.Tags,
			Status:     string(p.Status),
			Favourites: p.Favourites,
			Views:      p.Views,
			UpdatedAt:  p.UpdatedAt.String(),
			CreatedAt:  p.CreatedAt.String(),
		})
	}
	return &pb.ResPostList{
		Posts: convertedPosts,
	}, nil
}

func ChangePostStatus(ctx context.Context, params *pb.ReqPostChangeStatus) (*pb.ResPost, error) {
	if !security.IsAuthenticated(ctx) {
		return nil, security.GetUnauthenticatedError()
	}
	postStorage := storage.NewPostStorage()
	if !bson.IsObjectIdHex(params.Id) {
		return &pb.ResPost{
			Post: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid ID",
					Details: fmt.Sprintf("Post ID %s is not valid", params.Id),
				},
			},
		}, nil
	}
	post := postStorage.FindPostByID(params.Id)
	if post != nil {
		if !security.HasPostWritePermission(ctx, *post) {
			return nil, security.GetUnauthorisedError()
		}
		if post.ValidateStatus(params.NewStatus) {
			post.Status = post.ToPostStatus(params.NewStatus)
			post.UpdatedAt = time.Now()
			if postStorage.UpdatePost(post) {
				return &pb.ResPost{
					Post: &pb.Post{
						Id:         post.ID.Hex(),
						Title:      post.Title,
						Body:       post.Body,
						Categories: post.Categories,
						Tags:       post.Tags,
						Status:     string(post.Status),
						Favourites: post.Favourites,
						Views:      post.Views,
						UpdatedAt:  post.UpdatedAt.String(),
						CreatedAt:  post.CreatedAt.String(),
					},
				}, nil
			}
			return &pb.ResPost{
				Post: nil,
				Errors: []*pb.Error{
					{
						ID:      uuid.NewV4().String(),
						Code:    http.StatusInternalServerError,
						Title:   "Internal server error",
						Details: "Couldn't process the request",
					},
				},
			}, nil
		}
		return &pb.ResPost{
			Post: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Bad request",
					Details: fmt.Sprintf("Post status %s is not valid", params.NewStatus),
				},
			},
		}, nil
	}
	return &pb.ResPost{
		Post: nil,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusNotFound,
				Title:   "Not found",
				Details: fmt.Sprintf("Post with ID %s not found", params.Id),
			},
		},
	}, nil

}

func GetPost(ctx context.Context, params *pb.GetByID) (*pb.ResPost, error) {
	postStorage := storage.NewPostStorage()
	if !bson.IsObjectIdHex(params.Id) {
		return &pb.ResPost{
			Post: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid ID",
					Details: fmt.Sprintf("Post ID %s is not valid", params.Id),
				},
			},
		}, nil
	}
	post := postStorage.FindPostByID(params.Id)
	if post != nil {
		post.Views++
		if !postStorage.UpdatePost(post) {
			post.Views--
		}
		return &pb.ResPost{
			Post: &pb.Post{
				Id:         post.ID.Hex(),
				Title:      post.Title,
				Body:       post.Body,
				Categories: post.Categories,
				Tags:       post.Tags,
				Status:     string(post.Status),
				Favourites: post.Favourites,
				Views:      post.Views,
				UpdatedAt:  post.UpdatedAt.String(),
				CreatedAt:  post.CreatedAt.String(),
			},
		}, nil
	}
	return &pb.ResPost{
		Post: nil,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusNotFound,
				Title:   "Not found",
				Details: fmt.Sprintf("Post with ID %s not found", params.Id),
			},
		},
	}, nil
}

func DeletePost(ctx context.Context, params *pb.GetByID) (*pb.ResPostSuccess, error) {
	postStorage := storage.NewPostStorage()
	if !bson.IsObjectIdHex(params.Id) {
		return &pb.ResPostSuccess{
			Success: false,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid ID",
					Details: fmt.Sprintf("Post ID %s is not valid", params.Id),
				},
			},
		}, nil
	}
	post := postStorage.FindPostByID(params.Id)
	if post != nil {
		if !security.HasPostWritePermission(ctx, *post) {
			return nil, security.GetUnauthorisedError()
		}
		if !postStorage.DeletePost(post) {
			return &pb.ResPostSuccess{
				Success: false,
				Errors: []*pb.Error{
					{
						ID:      uuid.NewV4().String(),
						Code:    http.StatusInternalServerError,
						Title:   "Something went wrong",
						Details: fmt.Sprintf("Couldn't delete post with ID %s", params.Id),
					},
				},
			}, nil
		}
		return &pb.ResPostSuccess{
			Success: true,
		}, nil
	}
	return &pb.ResPostSuccess{
		Success: false,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusNotFound,
				Title:   "Not found",
				Details: fmt.Sprintf("Post with ID %s not found", params.Id),
			},
		},
	}, nil
}

func UpdatePost(ctx context.Context, params *pb.ReqPostUpdate) (*pb.ResPost, error) {
	if !security.IsAuthenticated(ctx) {
		return nil, security.GetUnauthenticatedError()
	}
	postStorage := storage.NewPostStorage()
	if !bson.IsObjectIdHex(params.Id) {
		return &pb.ResPost{
			Post: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid ID",
					Details: fmt.Sprintf("Post ID %s is not valid", params.Id),
				},
			},
		}, nil
	}
	post := postStorage.FindPostByID(params.Id)
	if post != nil {
		if !security.HasPostWritePermission(ctx, *post) {
			return nil, security.GetUnauthorisedError()
		}
		post.Title = params.Title
		post.Body = params.Body
		post.Tags = params.Tags
		post.Categories = params.Categories
		post.UpdatedAt = time.Now()
		if postStorage.UpdatePost(post) {
			return &pb.ResPost{
				Post: &pb.Post{
					Id:         post.ID.Hex(),
					Title:      post.Title,
					Body:       post.Body,
					Categories: post.Categories,
					Tags:       post.Tags,
					Status:     string(post.Status),
					Favourites: post.Favourites,
					Views:      post.Views,
					UpdatedAt:  post.UpdatedAt.String(),
					CreatedAt:  post.CreatedAt.String(),
				},
			}, nil
		}
		return &pb.ResPost{
			Post: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusInternalServerError,
					Title:   "Something went wrong",
					Details: fmt.Sprintf("Couldn't update post with ID %s", params.Id),
				},
			},
		}, nil
	}
	return &pb.ResPost{
		Post: nil,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusNotFound,
				Title:   "Not found",
				Details: fmt.Sprintf("Post with ID %s not found", params.Id),
			},
		},
	}, nil
}

func FavouritePost(ctx context.Context, params *pb.GetByID) (*pb.ResPost, error) {
	postStorage := storage.NewPostStorage()
	if !bson.IsObjectIdHex(params.Id) {
		return &pb.ResPost{
			Post: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid ID",
					Details: fmt.Sprintf("Post ID %s is not valid", params.Id),
				},
			},
		}, nil
	}
	post := postStorage.FindPostByID(params.Id)
	if post != nil {
		post.Favourites++
		if !postStorage.UpdatePost(post) {
			post.Favourites--
		}
		return &pb.ResPost{
			Post: &pb.Post{
				Id:         post.ID.Hex(),
				Title:      post.Title,
				Body:       post.Body,
				Categories: post.Categories,
				Tags:       post.Tags,
				Status:     string(post.Status),
				Favourites: post.Favourites,
				Views:      post.Views,
				UpdatedAt:  post.UpdatedAt.String(),
				CreatedAt:  post.CreatedAt.String(),
			},
		}, nil
	}
	return &pb.ResPost{
		Post: nil,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusNotFound,
				Title:   "Not found",
				Details: fmt.Sprintf("Post with ID %s not found", params.Id),
			},
		},
	}, nil
}
