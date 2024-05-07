package service

import (
	"context"
	"database/sql"
	"fmt"
)

type LoginService struct {
	db *sql.DB
}

func NewLoginService(db *sql.DB) *LoginService {
	return &LoginService{
		db: db,
	}
}

func (s *LoginService) Login(ctx context.Context, loginId, password string) (string, error) {
	statement := fmt.Sprintf("SELECT %s, %s from %s WHERE loginId = ?", userIdCol, passwordCol, tableAuth)
	prep, err := s.db.Prepare(statement)
	if err != nil {
		return "", err
	}
	defer prep.Close()

	var userId int64
	var passData string
	err = prep.QueryRowContext(ctx, loginId).Scan(&userId, &passData)
	if err != nil {
		return "", err
	}
	c := make(chan string)
	go GenerateToken(userId, c)
	token := <-c
	close(c)

	if password != passData {
		return "", fmt.Errorf("wrong password")
	}
	return token, nil
}
