package data

import (
	"fmt"
	"time"
)

var tableName = "user"
var idCol = "id"
var nameCol = "name"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func (user *User) Create() error {
	statement := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES (?, ?)", tableName, idCol, nameCol)
	stmt, err := mydb.Prepare(statement)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Name)
	if err != nil {
		return err
	}
	return nil
}
