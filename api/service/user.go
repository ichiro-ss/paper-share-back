package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var tableUser = "user"
var nameCol = "name"
var tableAuth = "authorizations"
var loginIdCol = "loginId"
var userIdCol = "userId"
var passwordCol = "password"

type UserService struct {
	db *sql.DB
}

// NewUserService returns new UserService.
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func GenerateToken(uid string) string {
	claims := jwt.MapClaims{
		"user_id": uid,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	// generate payload and header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, _ := token.SignedString([]byte("ACCESS_SECRET_KEY"))

	return accessToken
}

func (s *UserService) CreateUser(ctx context.Context, loginId, password string) (string, error) {
	// create user table
	statement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?)", tableUser, nameCol)
	prep, err := s.db.Prepare(statement)
	if err != nil {
		return "", err
	}

	defer prep.Close()
	// first, we use login id as username
	row, err := prep.ExecContext(ctx, loginId)
	if err != nil {
		return "", err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return "", err
	}

	// create auth table
	statement = fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES (?, ?, ?)", tableAuth, loginIdCol, userIdCol, passwordCol)
	fmt.Println("statement:", statement)
	prep, err = s.db.Prepare(statement)
	if err != nil {
		return "", err
	}

	defer prep.Close()
	_, err = prep.ExecContext(ctx, loginId, id, password)
	if err != nil {
		return "", err
	}
	token := GenerateToken(loginId)
	return token, nil
}
