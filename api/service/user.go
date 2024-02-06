package service

import (
	"context"
	"database/sql"
	"fmt"
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

func (s *UserService) Create(ctx context.Context, loginId, password string) error {
	// create user table
	statement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?)", tableUser, nameCol)
	prep, err := s.db.Prepare(statement)
	if err != nil {
		return err
	}

	defer prep.Close()
	// first, we use login id as username
	row, err := prep.ExecContext(ctx, loginId)
	if err != nil {
		return err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return err
	}

	// create auth table
	statement = fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES (?, ?, ?)", tableAuth, loginIdCol, userIdCol, passwordCol)
	prep, err = s.db.Prepare(statement)
	if err != nil {
		return err
	}

	defer prep.Close()
	_, err = prep.ExecContext(ctx, loginId, id, password)
	if err != nil {
		return err
	}
	return nil
}
