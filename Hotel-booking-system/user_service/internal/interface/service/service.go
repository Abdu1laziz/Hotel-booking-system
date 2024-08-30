package service

import (
	"context"
	intr "user-service/internal/interface"
	"user-service/models"
	"user-service/pkg/proto/user"
)

type Service struct {
	D intr.User
}

type Adjust struct {
	A intr.Adjust
}

func (u *Service)  LogIn(ctx context.Context, req *models.LogInRequest) (*models.LogInResponse, error){
	return u.D.LogIn(ctx,req)
}

func (u *Service) CreateUser(ctx context.Context, req *models.RegisterUserRequest) (*models.GeneralResponse, error) {
	return u.D.CreateUser(ctx, req)
}

func (u *Service) GetUser(ctx context.Context, req *models.GetUserRequest) (*models.GetUserResponse, error) {
	return u.D.GetUser(ctx, req)
}

func (u *Service) LastInserted(ctx context.Context, req *models.LastInsertedUser) (*models.GetUserResponse, error) {
	return u.D.LastInserted(ctx, req)
}

func (u *Service) UpdateUser(ctx context.Context, req *models.UpdateUserRequest) (*models.GeneralResponse, error) {
	return u.D.UpdateUser(ctx, req)
}

func (u *Service) LogOut(ctx context.Context, req *models.GetUserRequest) (*models.GeneralResponse, error) {
	return u.D.LogOut(ctx, req)
}

func (u *Service) DeletUser(ctx context.Context, req *models.GetUserRequest) (*models.GeneralResponse, error) {
	return u.D.DeletUser(ctx, req)
}

func (u *Adjust) AddUser(ctx context.Context, req *user.RegisterUserRequest) (*user.GeneralResponse, error) {
	return u.A.AddUser(ctx, req)
}

func (u *Adjust) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	return u.A.GetUser(ctx, req)
}

func (u *Adjust) LastOne(ctx context.Context, req *user.LastInsertedUser) (*user.GetUserResponse, error) {
	return u.A.LastOne(ctx, req)
}

func (u *Adjust) Update(ctx context.Context, req *user.UpdateUserRequest) (*user.GeneralResponse, error) {
	return u.A.Update(ctx, req)
}

func (u *Adjust) Logout(ctx context.Context, req *user.GetUserRequest) (*user.GeneralResponse, error) {
	return u.A.Logout(ctx, req)
}

func (u *Adjust) Delete(ctx context.Context, req *user.GetUserRequest) (*user.GeneralResponse, error) {
	return u.A.Delete(ctx, req)
}
func (u *Adjust) LogIn(ctx context.Context, req *user.LogInRequest) (*user.LogInResposne, error){
	return u.A.LogIn(ctx,req)
}