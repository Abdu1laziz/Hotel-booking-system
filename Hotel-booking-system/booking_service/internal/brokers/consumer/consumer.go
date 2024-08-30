package kafkaconsumer

import (
	interfaceservices "booking-service/internal/interface/services"
	"booking-service/models"
	"booking-service/pkg/protos/booking"
	"context"
	"encoding/json"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Consumer17 struct {
	A   *interfaceservices.AdjustDatabase
	Ctx context.Context
}

func (u *Consumer17) Consumer() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumeTopics("booking"),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	for {
		fetches := client.PollFetches(ctx)
		if err := fetches.Errors(); len(err) > 0 {
			log.Fatal(err)
		}
		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			for _, record := range ftp.Records {
				if err := u.Adjust(record); err != nil {
					log.Println(err)
				}
			}
		})
	}
}

func (u *Consumer17) Adjust(record *kgo.Record) error {
	switch string(record.Key) {
	case "create":
		if err := u.Create(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	case "update":
		if err := u.Update(record.Value); err != nil {
			log.Println(err)
			return err
		}
	case "delete":
		if err := u.Delete(record.Value); err != nil {
			log.Println(err)
			return err
		}
	case "createW":
		if err := u.CreateW(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	case "updateW":
		if err := u.UpdateW(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	case "deleteW":
		if err := u.DeleteW(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	}
	return nil
}

func (u *Consumer17) Create(req []byte) error {
	var req1 models.BookHotelRequest

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = booking.BookHotelRequest{
		UserID:       req1.UserID,
		HotelID:      req1.HotelID,
		RoomId:       req1.RoomID,
		RoomType:     req1.RoomType,
		CheckInDate:  timestamppb.New(req1.CheckInDate),
		CheckOutDate: timestamppb.New(req1.CheckOutDate),
	}
	_, err := u.A.Create(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer17) Update(req []byte) error {
	var req1 models.BookHotelUpdateRequest

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = booking.BookHotelUpdateRequest{
		Id:           req1.ID,
		RoomId:       req1.RoomID,
		RoomType:     req1.RoomType,
		CheckInDate:  req1.CheckInDate,
		CheckOutDate: req1.CheckOutDate,
	}
	_, err := u.A.Update(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer17) Delete(req []byte) error {
	var req1 models.CancelRoomRequest
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = booking.CancelROomRequest{
		Id: req1.ID,
	}

	_, err := u.A.Cancel(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer17) CreateW(req []byte) error {
	var req1 models.CreateWaitingList

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = booking.CreateWaitingList{
		UserId:       req1.UserID,
		UserEmail:    req1.UserEmail,
		RoomType:     req1.RoomType,
		HotelId:      req1.HotelID,
		CheckInDate:  timestamppb.New(req1.CheckInDate),
		CheckOutDate: timestamppb.New(req1.CheckOutDate),
	}
	_, err := u.A.CreateW(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer17) UpdateW(req []byte) error {
	var req1 models.UpdateWaitingListRequest
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = booking.UpdateWaitingListRequest{
		Id:           req1.ID,
		UserId:       req1.UserID,
		RoomType:     req1.RoomType,
		HotelId:      req1.HotelID,
		CheckInDate:  timestamppb.New(req1.CheckInDate),
		CheckOutDate: timestamppb.New(req1.CheckOutDate),
	}
	_, err := u.A.UpdateW(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer17) DeleteW(req []byte) error {
	var req1 models.DeleteWaitingList
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = booking.DeleteWaitingList{
		Id: req1.ID,
	}

	_, err := u.A.DeleteW(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
