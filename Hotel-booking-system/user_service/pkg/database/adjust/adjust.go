package adjust

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"user-service/models"
	db "user-service/pkg/database/sql"
	"user-service/pkg/proto/notification"

	"golang.org/x/crypto/bcrypt"
)

type Database struct {
	Db *sql.DB
	N  notification.NotificationClient
}

func (u *Database) LogIn(ctx context.Context, req *models.LogInRequest) (*models.LogInResponse, error) {
	query, args, err := db.LogIn(req)
	if err != nil {
		log.Println("Error building login query:", err)
		return nil, err
	}

	var password string
	if err := u.Db.QueryRow(query, args...).Scan(&password); err != nil {
		log.Println("Error retrieving password:", err)
		return nil, err
	}

	if u.ComparePassword(password, req.Password) {
		_, err := u.N.Email(ctx, &notification.EmailSend{Email: req.Email, Message: "You are logged in again, you have already logged in before 😉"})
		if err != nil {
			log.Println("Error sending email notification:", err)
		}
		return &models.LogInResponse{Status: true}, nil
	}
	return nil, errors.New("the password is not correct 🤨")
}

func (u *Database) CreateUser(ctx context.Context, req *models.RegisterUserRequest) (*models.GeneralResponse, error) {
	req.Password = u.Hashing(req.Password)
	query, args, err := db.Create(req)
	if err != nil {
		log.Println("Error building user creation query:", err)
		return nil, err
	}

	var id int
	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println("Error retrieving user ID:", err)
		return nil, err
	}

	if _, err = u.N.AddUser(ctx, &notification.AddnewUser{UserId: strconv.Itoa(id)}); err != nil {
		log.Println("Error notifying new user:", err)
	}

	_, err = u.N.Notification(ctx, &notification.ProduceMessage{UserId: int32(id), Message: fmt.Sprintf("Your account is created successfully with ID %v", id)})
	if err != nil {
		log.Println("Error sending notification:", err)
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("Successfully created ID %v 👍", id)}, nil
}

func (u *Database) GetUser(ctx context.Context, req *models.GetUserRequest) (*models.GetUserResponse, error) {
	query, args, err := db.Get(req)
	if err != nil {
		log.Println("Error building get user query:", err)
		return nil, errors.New("user not found 🤷‍♂️")
	}

	var res models.GetUserResponse
	if err := u.Db.QueryRow(query, args...).Scan(&res.ID, &res.Username, &res.Age, &res.Email, &res.LogOut); err != nil {
		log.Println("Error scanning user data:", err)
		return nil, err
	}
	return &res, nil
}

func (u *Database) LastInserted(ctx context.Context, req *models.LastInsertedUser) (*models.GetUserResponse, error) {
	query, args, err := db.LastInserted()
	if err != nil {
		log.Println("Error building last inserted query:", err)
		return nil, err
	}

	var res models.GetUserResponse
	if err := u.Db.QueryRow(query, args...).Scan(&res.ID, &res.Username, &res.Age, &res.Email, &res.LogOut); err != nil {
		log.Println("Error scanning last inserted user data:", err)
		return nil, err
	}
	return &res, nil
}

func (u *Database) UpdateUser(ctx context.Context, req *models.UpdateUserRequest) (*models.GeneralResponse, error) {
	req.Password = u.Hashing(req.Password)
	query, args, err := db.Update(req)
	if err != nil {
		log.Println("Error building update user query:", err)
		return nil, err
	}

	var id int
	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println("Error retrieving updated user ID:", err)
		return nil, err
	}

	if _, err = u.N.Notification(ctx, &notification.ProduceMessage{UserId: int32(id), Message: fmt.Sprintf("Your account has been updated successfully with ID %v", id)}); err != nil {
		log.Println("Error sending update notification:", err)
	}
	return &models.GeneralResponse{Message: fmt.Sprintf("Your account has been updated with ID %v 🤓", id)}, nil
}

func (u *Database) LogOut(ctx context.Context, req *models.GetUserRequest) (*models.GeneralResponse, error) {
	query, args, err := db.LogOut(req)
	if err != nil {
		log.Println("Error building logout query:", err)
		return nil, err
	}

	if _, err = u.Db.Exec(query, args...); err != nil {
		log.Println("Error executing logout query:", err)
		return nil, err
	}

	if _, err = u.N.Notification(ctx, &notification.ProduceMessage{UserId: int32(req.ID), Message: "You have successfully logged out, you can log in again! 👍"}); err != nil {
		log.Println("Error sending logout notification:", err)
	}
	return &models.GeneralResponse{Message: "Exit... 👉"}, nil
}

func (u *Database) DeleteUser(ctx context.Context, req *models.GetUserRequest) (*models.GeneralResponse, error) {
	query, args, err := db.Delete(req) // Исправлено на Delete
	if err != nil {
		log.Println("Error building delete user query:", err)
		return nil, err
	}

	if _, err = u.Db.Exec(query, args...); err != nil {
		log.Println("Error executing delete query:", err)
		return nil, err
	}

	if _, err = u.N.Notification(ctx, &notification.ProduceMessage{UserId: int32(req.ID), Message: "You have successfully deleted your account"}); err != nil {
		log.Println("Error sending delete notification:", err)
	}
	return &models.GeneralResponse{Message: "Deleting... ☠️"}, nil
}

func (u *Database) ComparePassword(hashed, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		log.Println("Error comparing passwords:", err)
		return false
	}
	return true
}

func (u *Database) Hashing(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return ""
	}
	return string(hashed)
}
