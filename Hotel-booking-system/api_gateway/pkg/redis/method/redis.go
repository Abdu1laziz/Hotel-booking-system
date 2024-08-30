package redismethod

import (
	"api-gateway/models"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	R   *redis.Client
	Ctx context.Context
}

func (u *Redis) Register(req *models.RegisterUserRequest) error {
	byted, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := u.R.Set(u.Ctx, req.Email, byted, time.Minute*1).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Redis) RegisterGet(email string) (*models.RegisterUserRequest, error) {
	val, err := u.R.Get(u.Ctx, email).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var res models.RegisterUserRequest

	if err := json.Unmarshal([]byte(val), &res); err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}

func (u *Redis) VerifyCodeRequest(req *models.VerifyRequest) error {
	if err := u.R.Set(u.Ctx, req.Email+"code", req.Code, time.Minute*5).Err(); err != nil {
		return err
	}
	return nil
}

func (u *Redis) VerifyCodeResponse(req *models.VerifyRequest) (string, error) {
	var code string
	if err := u.R.Get(u.Ctx, req.Email+"code").Scan(&code); err != nil {
		return "", nil
	}
	return code, nil
}

func (u *Redis) SetUser(req *models.GetUserResponse) error {
	byted, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return err
	}
	id := strconv.Itoa(int(req.ID))
	if err := u.R.Set(u.Ctx, id, byted, 0).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Redis) GetUser(req *models.GetUserRequest) (*models.GetUserResponse, error) {
	var res models.GetUserResponse

	val, err := u.R.Get(u.Ctx, strconv.Itoa(int(req.ID))).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := json.Unmarshal([]byte(val), &res); err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}
