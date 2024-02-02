package test

import (
	"api/data"
	"fmt"
	"log"
)

func TestCreateUser() {
	id := 123
	name := "name"
	user := data.User{Id: id, Name: name}

	fmt.Println("id:", user.Id, ", pass:", user.Name)
	fmt.Println("Userの作成開始")
	if err := user.Create(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Userの作成に成功")
}
