package sql

import (
	"log"
	"user-service/models"

	"github.com/Masterminds/squirrel"
)

func Create(req *models.RegisterUserRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("users").
		Columns("username", "age", "email", "password").
		Values(req.Username, req.Age, req.Email, req.Password).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func Get(req *models.GetUserRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Select("id,username,age,email,logout").
		From("users").
		Where(squirrel.Eq{"id": req.ID, "logout": false}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func LastInserted() (string, []interface{}, error) {
	query, args, err := squirrel.Select("id,username,age,email,logout").
		From("users").
		OrderBy("id DESC").
		Limit(1).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func Update(req *models.UpdateUserRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Update("users").
		SetMap(map[string]interface{}{
			"username": req.Username,
			"age":      req.Age,
			"email":    req.Email,
			"password": req.Password,
		}).Where(squirrel.Eq{"id": req.ID}).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func LogIn(req *models.LogInRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Select("password").
		From("users").
		Where(squirrel.Eq{"email": req.Email}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func LogOut(req *models.GetUserRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Update("users").
		SetMap(map[string]interface{}{
			"logout": true,
		}).Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func Delete(req *models.GetUserRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("users").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}
