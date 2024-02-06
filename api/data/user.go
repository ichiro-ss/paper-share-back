package data

import (
	"fmt"
)

var tableUser = "user"
var nameCol = "name"
var tableAuth = "authorizations"
var loginIdCol = "loginId"
var userIdCol = "userId"
var passwordCol = "password"

type User struct {
	LoginId  int    `json:"loginId"`
	Password string `json:"password"`
}

func (user *User) Create() error {
	// create user table
	statement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?)", tableUser, nameCol)
	prep, err := mydb.Prepare(statement)
	if err != nil {
		return err
	}

	defer prep.Close()
	// first, we use login id as username
	row, err := prep.Exec(user.LoginId)
	if err != nil {
		return err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return err
	}

	// create auth table
	statement = fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES (?, ?, ?)", tableAuth, loginIdCol, userIdCol, passwordCol)
	prep, err = mydb.Prepare(statement)
	if err != nil {
		return err
	}

	defer prep.Close()
	_, err = prep.Exec(user.LoginId, id, user.Password)
	if err != nil {
		return err
	}
	return nil
}
