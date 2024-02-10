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

func GenerateToken(uid int64) string {
	claims := jwt.MapClaims{
		"user_id": uid,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	// generate payload and header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, _ := token.SignedString([]byte("ACCESS_SECRET_KEY"))

	return accessToken
}

func TokenToId(tkStr string) (int64, error) {
	tk, err := jwt.Parse(tkStr, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tk.Header["alg"])
		}
		return []byte("ACCESS_SECRET_KEY"), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		return int64(claims["user_id"].(float64)), nil
	} else {
		return 0, fmt.Errorf("wrong bearer token")
	}
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
	prep, err = s.db.Prepare(statement)
	if err != nil {
		return "", err
	}

	defer prep.Close()
	_, err = prep.ExecContext(ctx, loginId, id, password)
	if err != nil {
		return "", err
	}
	token := GenerateToken(id)
	return token, nil
}

func (s *UserService) ReadUser(ctx context.Context, token string) (string, error) {
	// read user
	statement := fmt.Sprintf("SELECT %s from %s WHERE id = ?", nameCol, tableUser)
	prep, err := s.db.Prepare(statement)
	if err != nil {
		return "", err
	}
	defer prep.Close()

	id, err := TokenToId(token)
	if err != nil {
		return "", err
	}

	var name string
	err = prep.QueryRowContext(ctx, id).Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}

func (s *UserService) EditUser(ctx context.Context, token, name string) (string, error) {
	// read user
	statement := fmt.Sprintf("UPDATE %s SET %s=? WHERE id=?", tableUser, nameCol)
	prep, err := s.db.Prepare(statement)
	if err != nil {
		return "", err
	}
	defer prep.Close()

	id, err := TokenToId(token)
	if err != nil {
		return "", err
	}

	_, err = prep.ExecContext(ctx, name, id)
	if err != nil {
		return "", err
	}

	return name, nil
}
