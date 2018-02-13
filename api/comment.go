package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/s4kibs4mi/rapunzel-blog/models"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	pb "github.com/s4kibs4mi/rapunzel-blog/protos"
	"github.com/s4kibs4mi/rapunzel-blog/security"
	"github.com/s4kibs4mi/rapunzel-blog/storage"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func CreateComment(ctx context.Context, params *protos.ReqCommentCreate) (*protos.ResComment, error) {
	if !security.IsAuthenticated(ctx) {
		return nil, security.GetUnauthenticatedError()
	}
	commentStorage := storage.NewCommentStorage()
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
		return &pb.ResComment{
			Comment: nil,
			Errors: []*pb.Error{
				{
					ID:           uuid.NewV4().String(),
					Code:         http.StatusUnprocessableEntity,
					Title:        "Comment Validation Error",
					Details:      "Required fields are not valid",
					ErrorDetails: validationErrorDetails,
				},
			},
		}, nil
	}

	if bson.IsObjectIdHex(params.PostId) && postStorage.FindPostByID(params.PostId) != nil {
		c := models.Comment{}
		c.ID = bson.NewObjectId()
		c.Title = params.Title
		c.Body = params.Body
		c.PostID = bson.ObjectIdHex(params.PostId)
		c.UserID = bson.ObjectIdHex(security.ReadUserIDFromContext(ctx))
		c.Favourites = 0
		c.CreatedAt = time.Now()
		c.UpdatedAt = time.Now()
		if commentStorage.SaveComment(&c) {
			return &pb.ResComment{
				Comment: &pb.Comment{
					Id:        c.ID.Hex(),
					PostId:    c.PostID.Hex(),
					UserId:    c.UserID.Hex(),
					Title:     c.Title,
					Body:      c.Body,
					CreatedAt: c.CreatedAt.String(),
					UpdatedAt: c.UpdatedAt.String(),
				},
				Errors: nil,
			}, nil
		}
		return &pb.ResComment{
			Comment: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusInternalServerError,
					Title:   "Something went wrong",
					Details: "Couldn't process the request",
				},
			},
		}, nil
	}
	return &pb.ResComment{
		Comment: nil,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusNotFound,
				Title:   "Post not found",
				Details: fmt.Sprintf("Post with ID %s not found", params.PostId),
			},
		},
	}, nil
}

func UpdateComment(ctx context.Context, params *protos.ReqCommentCreate) (*protos.ResComment, error) {
	if !security.IsAuthenticated(ctx) {
		return nil, security.GetUnauthenticatedError()
	}
	commentStorage := storage.NewCommentStorage()
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
		return &pb.ResComment{
			Comment: nil,
			Errors: []*pb.Error{
				{
					ID:           uuid.NewV4().String(),
					Code:         http.StatusUnprocessableEntity,
					Title:        "Comment Validation Error",
					Details:      "Required fields are not valid",
					ErrorDetails: validationErrorDetails,
				},
			},
		}, nil
	}

	if bson.IsObjectIdHex(params.PostId) && postStorage.FindPostByID(params.PostId) != nil {
		c := models.Comment{}
		c.ID = bson.NewObjectId()
		c.Title = params.Title
		c.Body = params.Body
		c.PostID = bson.ObjectIdHex(params.PostId)
		c.UserID = bson.ObjectIdHex(security.ReadUserIDFromContext(ctx))
		c.Favourites = 0
		c.CreatedAt = time.Now()
		c.UpdatedAt = time.Now()
		if commentStorage.SaveComment(&c) {
			return &pb.ResComment{
				Comment: &pb.Comment{
					Id:        c.ID.Hex(),
					PostId:    c.PostID.Hex(),
					UserId:    c.UserID.Hex(),
					Title:     c.Title,
					Body:      c.Body,
					CreatedAt: c.CreatedAt.String(),
					UpdatedAt: c.UpdatedAt.String(),
				},
				Errors: nil,
			}, nil
		}
		return &pb.ResComment{
			Comment: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusInternalServerError,
					Title:   "Something went wrong",
					Details: "Couldn't process the request",
				},
			},
		}, nil
	}
	return &pb.ResComment{
		Comment: nil,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusNotFound,
				Title:   "Post not found",
				Details: fmt.Sprintf("Post with ID %s not found", params.PostId),
			},
		},
	}, nil
}

func GetComments(ctx context.Context, params *protos.GetByQuery) (*protos.ResCommentList, error) {
	commentStorage := storage.NewCommentStorage()
	comments := commentStorage.FindCommentsByQuery(params.Query)
	var convertedComments []*pb.Comment
	for _, c := range comments {
		convertedComments = append(convertedComments, &pb.Comment{
			Id:        c.ID.Hex(),
			Title:     c.Title,
			Body:      c.Body,
			PostId:    c.PostID.Hex(),
			UserId:    c.UserID.Hex(),
			UpdatedAt: c.UpdatedAt.String(),
			CreatedAt: c.CreatedAt.String(),
		})
	}
	return &pb.ResCommentList{
		Comments: convertedComments,
	}, nil
}

func GetComment(ctx context.Context, params *protos.GetByID) (*protos.ResComment, error) {
	commentStorage := storage.NewCommentStorage()
	if !bson.IsObjectIdHex(params.Id) {
		return &pb.ResComment{
			Comment: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid ID",
					Details: fmt.Sprintf("Comment ID %s is not valid", params.Id),
				},
			},
		}, nil
	}
	comment := commentStorage.FindCommentByID(params.Id)
	if comment != nil {
		return &pb.ResComment{
			Comment: &pb.Comment{
				Id:        comment.ID.Hex(),
				Title:     comment.Title,
				Body:      comment.Body,
				UpdatedAt: comment.UpdatedAt.String(),
				CreatedAt: comment.CreatedAt.String(),
			},
		}, nil
	}
	return &pb.ResComment{
		Comment: nil,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusNotFound,
				Title:   "Not found",
				Details: fmt.Sprintf("Comment with ID %s not found", params.Id),
			},
		},
	}, nil
}
