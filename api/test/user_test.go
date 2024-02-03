package test

import (
	"api/data"
	"testing"
)

func TestUser(t *testing.T) {
	id := 123
	name := "name"
	user := data.User{Id: id, Name: name}

	if err := user.Create(); err != nil {
		t.Error(err)
		return
	}
}
