package test

import (
	"api/data"
	"testing"
)

func TestUser(t *testing.T) {
	id := 123
	name := "name"
	user := data.User{Id: id, Name: name}

	// test for creating user
	if err := user.Create(); err != nil {
		t.Error(err)
		return
	}

	// test for getting user
	if test_user, err := data.Get(id); err != nil {
		t.Error(err)
		return
	}

	//test for updating user

	//test for deleting user
	if err = data.Delete(id); err != nil {
		t.Error(err)
		return
	}
}
