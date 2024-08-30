package adjustservice

import (
	"context"
	"log"
	"user-service/internal/interface/service"
	"user-service/models"
	"user-service/pkg/proto/user"
)

type Adjust struct {
	S *service.Service
}

func (u *Adjust) LogIn(ctx context.Context, req *user.LogInRequest) (*user.LogInResposne, error) {
	var newreq = models.LogInRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	res, err := u.S.LogIn(ctx, &newreq)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user.LogInResposne{Status: res.Status}, nil
}

func (u *Adjust) AddUser(ctx context.Context, req *user.RegisterUserRequest) (*user.GeneralResponse, error) {
	var newreq = models.RegisterUserRequest{
		Username: req.Username,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
	}
	res, err := u.S.CreateUser(ctx, &newreq)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user.GeneralResponse{Message: res.Message}, nil
}

func (u *Adjust) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	res, err := u.S.GetUser(ctx, &models.GetUserRequest{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user.GetUserResponse{Id: res.ID, Username: res.Username, Age: res.Age, Email: res.Email, Logout: res.LogOut}, nil

}

func (u *Adjust) LastOne(ctx context.Context, req *user.LastInsertedUser) (*user.GetUserResponse, error) {
	res, err := u.S.LastInserted(ctx, &models.LastInsertedUser{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user.GetUserResponse{Id: res.ID, Username: res.Username, Age: res.Age, Email: res.Email, Logout: res.LogOut}, nil
}

func (u *Adjust) Update(ctx context.Context, req *user.UpdateUserRequest) (*user.GeneralResponse, error) {
	var newreq = models.UpdateUserRequest{
		ID:       req.Id,
		Username: req.Username,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
	}
	res, err := u.S.UpdateUser(ctx, &newreq)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user.GeneralResponse{Message: res.Message}, nil
}

func (u *Adjust) Logout(ctx context.Context, req *user.GetUserRequest) (*user.GeneralResponse, error) {
	res, err := u.S.LogOut(ctx, &models.GetUserRequest{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user.GeneralResponse{Message: res.Message}, nil
}

func (u *Adjust) Delete(ctx context.Context, req *user.GetUserRequest) (*user.GeneralResponse, error) {
	res, err := u.S.DeletUser(ctx, &models.GetUserRequest{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user.GeneralResponse{Message: res.Message}, nil
}
