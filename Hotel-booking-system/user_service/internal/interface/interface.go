package interface17

import (
	"context"
	"user-service/models"
	"user-service/pkg/proto/user"
)

type User interface {
	CreateUser(ctx context.Context, req *models.RegisterUserRequest) (*models.GeneralResponse, error)
	LogIn(ctx context.Context, req *models.LogInRequest) (*models.LogInResponse, error)
	GetUser(ctx context.Context, req *models.GetUserRequest) (*models.GetUserResponse, error)
	LastInserted(ctx context.Context, req *models.LastInsertedUser) (*models.GetUserResponse, error)
	UpdateUser(ctx context.Context, req *models.UpdateUserRequest) (*models.GeneralResponse, error)
	LogOut(ctx context.Context, req *models.GetUserRequest) (*models.GeneralResponse, error)
	DeletUser(ctx context.Context, req *models.GetUserRequest) (*models.GeneralResponse, error)
}

type Adjust interface {
	AddUser(ctx context.Context, req *user.RegisterUserRequest) (*user.GeneralResponse, error)
	LogIn(ctx context.Context, req *user.LogInRequest) (*user.LogInResposne, error)
	GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error)
	LastOne(ctx context.Context, req *user.LastInsertedUser) (*user.GetUserResponse, error)
	Update(ctx context.Context, req *user.UpdateUserRequest) (*user.GeneralResponse, error)
	Logout(ctx context.Context, req *user.GetUserRequest) (*user.GeneralResponse, error)
	Delete(ctx context.Context, req *user.GetUserRequest) (*user.GeneralResponse, error)
}
