package test

import (
	"api/data"
	"testing"
)

func TestUser(t *testing.T) {
	id := 123
	name := "name"
	user := data.User{Id: id, Name: name}

	t.Run("creating user", func(t *testing.T) {
		if err := user.Create(); err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("getting user", func(t *testing.T) {
		if test_user, err := data.Get(id); err != nil {
			t.Error(err)
			return
		}
	})

	//test for updating user

	t.Run("deleting user", func(t *testing.T) {
		if err = data.Delete(id); err != nil {
			t.Error(err)
			return
		}
	})
}
