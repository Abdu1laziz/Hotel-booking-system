package grpcmethods

import (
	"booking-service/internal/brokers/producer"
	interfaceservices "booking-service/internal/interface/services"
	"booking-service/pkg/database/methods"
	"booking-service/pkg/protos/booking"
	"context"
	"log"
)

type Grpc struct {
	booking.UnimplementedBookHotelServer
	A *interfaceservices.AdjustDatabase
}

func (u *Grpc) CancelWaiting(ctx context.Context, req *booking.Bytes) (*booking.GeneralResponse, error) {
	if err := producer.Producer("deleteW", "booking", req.All); err != nil {
		return nil, err
	}
	return &booking.GeneralResponse{Message: "User is deleting will get notification when it's cancelled"}, nil
}
func (u *Grpc) Create(ctx context.Context, req *booking.Bytes) (*booking.GeneralResponse, error) {
	if err := producer.Producer("create", "booking", req.All); err != nil {
		return nil, err
	}
	return &booking.GeneralResponse{Message: "Creating your request,you will get notification when it's created"}, nil
}
func (u *Grpc) CreateWaiting(ctx context.Context, req *booking.Bytes) (*booking.GeneralResponse, error) {
	if err := producer.Producer("createW", "booking", req.All); err != nil {
		return nil, err
	}
	return &booking.GeneralResponse{Message: "Creating your request,you will get notification when it's created"}, nil
}
func (u *Grpc) Delete(ctx context.Context, req *booking.Bytes) (*booking.GeneralResponse, error) {
	if err := producer.Producer("delete", "booking", req.All); err != nil {
		return nil, err
	}
	return &booking.GeneralResponse{Message: "Cancelling your book request,you will get notification when it's cancelled"}, nil
}
func (u *Grpc) Get(ctx context.Context, req *booking.GetUsersBookRequest) (*booking.GetUsersBookResponse, error) {
	res, err := u.A.Get(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
func (u *Grpc) GetWaitinglist(ctx context.Context, req *booking.GetWaitinglistRequest) (*booking.GetWaitinglistResponse, error) {
	res, err := u.A.GetW(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
func (u *Grpc) Update(ctx context.Context, req *booking.Bytes) (*booking.GeneralResponse, error) {
	if err := producer.Producer("update", "booking", req.All); err != nil {
		return nil, err
	}
	return &booking.GeneralResponse{Message: "Updating your request,you will get notification when it's updated"}, nil
}
func (u *Grpc) UpdateWaiting(ctx context.Context, req *booking.Bytes) (*booking.GeneralResponse, error) {
	if err := producer.Producer("updateW", "booking", req.All); err != nil {
		return nil, err
	}
	return &booking.GeneralResponse{Message: "Updating your request,you will get notification when it's updated"}, nil
}

func (u *Grpc) Getall(context.Context, *booking.Request) (*booking.Response, error) {
	res, err := methods.GetAllWAitingUSers()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
