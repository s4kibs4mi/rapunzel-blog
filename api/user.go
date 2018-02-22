package api

import (
	"context"
	pb "github.com/s4kibs4mi/rapunzel-blog/proto/defs"
	"github.com/s4kibs4mi/rapunzel-blog/storage"
	"github.com/satori/go.uuid"
	"net/http"
	"github.com/s4kibs4mi/rapunzel-blog/models"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/s4kibs4mi/rapunzel-blog/security"
	"github.com/goware/emailx"
	"strings"
)

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func Register(ctx context.Context, params *pb.ReqRegistration) (*pb.ResRegistration, error) {
	data := storage.NewUserStorage()
	var validationErrorDetails []*pb.ErrorDetails

	// Validating Name
	var nameErrors []string
	nameLen := len(params.Name)
	if nameLen < 3 || nameLen > 50 {
		nameErrors = append(nameErrors, "Name length must be between 3 to 50")
	}
	if len(nameErrors) != 0 {
		uErr := pb.ErrorDetails{
			Field:   "Name",
			Details: nameErrors,
		}
		validationErrorDetails = append(validationErrorDetails, &uErr)
	}
	// Validating Username
	var usernameErrors []string
	usernameLen := len(params.Username)
	if usernameLen < 3 || usernameLen > 50 {
		usernameErrors = append(usernameErrors, "Username length must be between 3 to 50")
	}
	if u := data.FindByUsername(params.Username); u != nil {
		usernameErrors = append(usernameErrors, "Username already exists.")
	}
	if len(usernameErrors) != 0 {
		uErr := pb.ErrorDetails{
			Field:   "Username",
			Details: usernameErrors,
		}
		validationErrorDetails = append(validationErrorDetails, &uErr)
	}
	// Validating Password
	var passwordErrors []string
	passwordLen := len(params.Password)
	if passwordLen < 8 || passwordLen > 50 {
		passwordErrors = append(passwordErrors, "Password length must be between 8 to 50")
	}
	if len(passwordErrors) != 0 {
		pErr := pb.ErrorDetails{
			Field:   "Password",
			Details: passwordErrors,
		}
		validationErrorDetails = append(validationErrorDetails, &pErr)
	}
	// Validating Email
	var emailErrors []string
	emailValidationErr := emailx.Validate(params.Email)
	if emailValidationErr != nil {
		emailErrors = append(emailErrors, "Email must be valid.")
	}
	if u := data.FindByEmail(params.Email); u != nil {
		emailErrors = append(emailErrors, "Email already exists.")
	}
	if len(emailErrors) != 0 {
		eErr := pb.ErrorDetails{
			Field:   "Email",
			Details: emailErrors,
		}
		validationErrorDetails = append(validationErrorDetails, &eErr)
	}
	if len(validationErrorDetails) != 0 {
		return &pb.ResRegistration{
			User: nil,
			Errors: []*pb.Error{
				{
					ID:           uuid.NewV4().String(),
					Code:         http.StatusUnprocessableEntity,
					Title:        "User Validation Error",
					Details:      "Required fields are not valid",
					ErrorDetails: validationErrorDetails,
				},
			},
		}, nil
	}
	u := &models.User{}
	u.ID = bson.NewObjectId()
	u.Name = params.Name
	u.Username = params.Username
	u.Email = params.Email
	u.Password = security.NewBCryptPassword(params.Password)
	u.Details = params.Details
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	if data.Count() == 0 {
		u.UserType = models.UserTypeParent
		u.UserStatus = models.UserStatusVerified
	} else {
		u.UserType = models.UserTypeGhost
		u.UserStatus = models.UserStatusRegistered
	}
	if data.Save(*u) {
		return &pb.ResRegistration{
			User: &pb.User{
				ID:         u.ID.Hex(),
				Name:       u.Name,
				Username:   u.Username,
				Email:      u.Email,
				Details:    u.Details,
				UserStatus: string(u.UserStatus),
				UserType:   string(u.UserType),
				CreatedAt:  u.CreatedAt.String(),
				UpdatedAt:  u.UpdatedAt.String(),
			},
		}, nil
	}
	return &pb.ResRegistration{
		User: nil,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusInternalServerError,
				Title:   "Unable to process request",
				Details: "Something went wrong",
			},
		},
	}, nil
}

func Login(ctx context.Context, params *pb.ReqLogin) (*pb.ResLogin, error) {
	userData := storage.NewUserStorage()
	var validationErrorDetails []*pb.ErrorDetails
	// Validating Username
	var usernameErrors []string
	usernameLen := len(params.Username)
	if usernameLen < 3 || usernameLen > 50 {
		usernameErrors = append(usernameErrors, "Username length must be between 3 to 50")
	}
	if len(usernameErrors) != 0 {
		uErr := pb.ErrorDetails{
			Field:   "Username",
			Details: usernameErrors,
		}
		validationErrorDetails = append(validationErrorDetails, &uErr)
	}
	// Validating Password
	var passwordErrors []string
	passwordLen := len(params.Password)
	if passwordLen < 8 || passwordLen > 50 {
		passwordErrors = append(passwordErrors, "Password length must be between 8 to 50")
	}
	if len(passwordErrors) != 0 {
		pErr := pb.ErrorDetails{
			Field:   "Password",
			Details: passwordErrors,
		}
		validationErrorDetails = append(validationErrorDetails, &pErr)
	}
	if len(validationErrorDetails) != 0 {
		return &pb.ResLogin{
			Session: nil,
			Errors: []*pb.Error{
				{
					ID:           uuid.NewV4().String(),
					Code:         http.StatusUnprocessableEntity,
					Title:        "User Validation Error",
					Details:      "Required fields are not valid",
					ErrorDetails: validationErrorDetails,
				},
			},
		}, nil
	}

	u := userData.FindByUsername(params.Username)
	if u == nil {
		return &pb.ResLogin{
			Session: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusNotFound,
					Title:   "Authorization failed",
					Details: "Username is not registered",
				},
			},
		}, nil
	}
	if !security.CheckBCryptPassword(u.Password, params.Password) {
		return &pb.ResLogin{
			Session: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusUnauthorized,
					Title:   "Authorization failed",
					Details: "Username & password mismatch",
				},
			},
		}, nil
	}
	if !security.HasLoginPermissions(u) {
		return &pb.ResLogin{
			Session: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusForbidden,
					Title:   "Unverified",
					Details: "User is not verified to login",
				},
			},
		}, nil
	}

	sessionData := storage.NewSessionStorage()
	session := &models.Session{
		ID:           bson.NewObjectId(),
		UserID:       u.ID,
		AccessToken:  uuid.NewV4().String(),
		RefreshToken: uuid.NewV4().String(),
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(24 * time.Hour),
	}
	if sessionData.SaveSession(session) {
		return &pb.ResLogin{
			Session: &pb.Session{
				AccessToken:  session.AccessToken,
				RefreshToken: session.RefreshToken,
				CreatedTime:  session.CreatedAt.String(),
				ExpireTime:   session.ExpiredAt.String(),
			},
		}, nil
	}
	return &pb.ResLogin{
		Session: nil,
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

func GetProfile(ctx context.Context, params *pb.ReqProfile) (*pb.ResProfile, error) {
	if !security.IsAuthenticated(ctx) {
		return nil, security.GetUnauthenticatedError()
	}

	userData := storage.NewUserStorage()

	params.UserID = strings.TrimSpace(params.UserID)
	if params.UserID != "" && !bson.IsObjectIdHex(params.UserID) {
		return &pb.ResProfile{
			User: nil,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid data",
					Details: "Invalid userID",
				},
			},
		}, nil
	}
	if params.UserID == "" {
		params.UserID = security.ReadUserIDFromContext(ctx)
	}
	user := userData.FindByID(bson.ObjectIdHex(params.UserID))

	if user != nil {
		return &pb.ResProfile{
			User: &pb.User{
				ID:         user.ID.Hex(),
				Name:       user.Name,
				Username:   user.Username,
				Email:      user.Email,
				Details:    user.Details,
				UserStatus: string(user.UserStatus),
				UserType:   string(user.UserType),
				CreatedAt:  user.CreatedAt.String(),
				UpdatedAt:  user.UpdatedAt.String(),
			},
		}, nil
	}
	return &pb.ResProfile{
		User: nil,
		Errors: []*pb.Error{
			{
				ID:      uuid.NewV4().String(),
				Code:    http.StatusNotFound,
				Title:   "User not found",
				Details: "User not found",
			},
		},
	}, nil
}

func ChangeUserStatus(ctx context.Context, params *pb.ReqChangeUserStatus) (*pb.ResChangeUserStatus, error) {
	if !security.IsAuthenticated(ctx) {
		return nil, security.GetUnauthenticatedError()
	}
	if !security.HasPermissionAsParent(ctx) {
		return nil, security.GetUnauthorisedError()
	}
	if !isUserStatusValid(params.NewStatus) {
		return &pb.ResChangeUserStatus{
			Success: false,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid data",
					Details: "Invalid user status",
				},
			},
		}, nil
	}
	if !bson.IsObjectIdHex(params.UserID) {
		return &pb.ResChangeUserStatus{
			Success: false,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid data",
					Details: "Invalid userID",
				},
			},
		}, nil
	}

	userStorage := storage.NewUserStorage()
	u := userStorage.FindByID(bson.ObjectIdHex(params.UserID))
	if u == nil {
		return &pb.ResChangeUserStatus{
			Success: false,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusNotFound,
					Title:   "User not found",
					Details: "User not found",
				},
			},
		}, nil
	}

	u.UserStatus = models.UserStatus(params.NewStatus)
	if userStorage.Update(u) {
		return &pb.ResChangeUserStatus{
			Success: true,
			Errors:  nil,
		}, nil
	}
	return &pb.ResChangeUserStatus{
		Success: false,
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

func ChangeUserType(ctx context.Context, params *pb.ReqChangeUserType) (*pb.ResChangeUserType, error) {
	if !security.IsAuthenticated(ctx) {
		return nil, security.GetUnauthenticatedError()
	}
	if !security.HasPermissionAsParent(ctx) {
		return nil, security.GetUnauthorisedError()
	}
	if !isUserTypeValid(params.NewType) {
		return &pb.ResChangeUserType{
			Success: false,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid data",
					Details: "Invalid user type",
				},
			},
		}, nil
	}
	if !bson.IsObjectIdHex(params.UserID) {
		return &pb.ResChangeUserType{
			Success: false,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusBadRequest,
					Title:   "Invalid data",
					Details: "Invalid userID",
				},
			},
		}, nil
	}

	userStorage := storage.NewUserStorage()
	u := userStorage.FindByID(bson.ObjectIdHex(params.UserID))
	if u == nil {
		return &pb.ResChangeUserType{
			Success: false,
			Errors: []*pb.Error{
				{
					ID:      uuid.NewV4().String(),
					Code:    http.StatusNotFound,
					Title:   "User not found",
					Details: "User not found",
				},
			},
		}, nil
	}

	u.UserType = models.UserType(params.NewType)
	if userStorage.Update(u) {
		return &pb.ResChangeUserType{
			Success: true,
			Errors:  nil,
		}, nil
	}
	return &pb.ResChangeUserType{
		Success: false,
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
