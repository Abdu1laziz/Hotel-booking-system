package grpcmethods

import (
	"context"
	"log"
	"user-service/internal/interface/service"
	"user-service/pkg/proto/user"
)

type Service struct {
	user.UnimplementedUserServer
	S *service.Adjust
}

func (u *Service) Register(ctx context.Context, req *user.RegisterUserRequest) (*user.GeneralResponse, error) {
	res, err := u.S.AddUser(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *Service) LogIn(ctx context.Context, req *user.LogInRequest) (*user.LogInResposne, error) {
	res, err := u.S.LogIn(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *Service) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	res, err := u.S.GetUser(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *Service) LastInserted(ctx context.Context, req *user.LastInsertedUser) (*user.GetUserResponse, error) {
	res, err := u.S.LastOne(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *Service) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.GeneralResponse, error) {
	res, err := u.S.Update(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *Service) LogOut(ctx context.Context, req *user.GetUserRequest) (*user.GeneralResponse, error) {
	res, err := u.S.Logout(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *Service) DeleteUser(ctx context.Context, req *user.GetUserRequest) (*user.GeneralResponse, error) {
	res, err := u.S.Delete(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
