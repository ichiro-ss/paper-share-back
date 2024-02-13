package test

import (
	"api/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"testing"
)

const tkStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc4MTA4ODIsInVzZXJfaWQiOjE3fQ.uuswOlREnkOL7GD4pyRt5Deg-klNrvH9hMIC_l92P4w"

func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var (
	loginId  = RandomString(20)
	password = "pass"
)

func TestConnectDB(t *testing.T) {
	t.Run("creating user", func(t *testing.T) {
		createUserReq := model.CreateUserRequest{
			LoginId:  loginId,
			Password: password,
		}
		userJson, err := json.Marshal(createUserReq)
		if err != nil {
			t.Error(err)
		}
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://localhost:8080/users", bytes.NewBuffer(userJson))
		if err != nil {
			t.Error(err)
		}

		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil || body == nil {
			t.Error(err)
		}
		fmt.Println(string(body))
	})

	t.Run("get user info(name)", func(t *testing.T) {
		// readUserReq := model.UserHeader{
		// 	Token: tkStr,
		// }
		// readJson, err := json.Marshal(readUserReq)
		// if err != nil {
		// 	t.Error(err)
		// }
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://localhost:8080/users", nil)
		if err != nil {
			t.Error(err)
		}
		bearer := fmt.Sprintf("Bearer %s", tkStr)
		req.Header.Set("Authorization", bearer)
		// dump, _ := httputil.DumpRequestOut(req, true)
		// fmt.Printf("%s", dump)

		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil || body == nil {
			t.Error(err)
		}
		fmt.Println(string(body))
	})
	t.Run("edit user info(name)", func(t *testing.T) {
		editUserReq := model.EditUserRequest{
			Name: "new",
		}
		userJson, err := json.Marshal(editUserReq)
		if err != nil {
			t.Error(err)
		}
		client := &http.Client{}
		req, err := http.NewRequest("PUT", "http://localhost:8080/users", bytes.NewBuffer(userJson))
		if err != nil {
			t.Error(err)
		}
		bearer := fmt.Sprintf("Bearer %s", tkStr)
		req.Header.Set("Authorization", bearer)
		// dump, _ := httputil.DumpRequestOut(req, true)
		// fmt.Printf("%s", dump)

		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil || body == nil {
			t.Error(err)
		}
		fmt.Println(string(body))
	})

	t.Run("login", func(t *testing.T) {
		loginReq := model.CreateUserRequest{
			LoginId:  loginId,
			Password: password,
		}
		userJson, err := json.Marshal(loginReq)
		if err != nil {
			t.Error(err)
		}
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewBuffer(userJson))
		if err != nil {
			t.Error(err)
		}

		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil || body == nil {
			t.Error(err)
		}
		fmt.Println(string(body))
	})
}
