package model

type (
	CreateUserRequest struct {
		LoginId  int    `json:"loginId"`
		Password string `json:"password"`
	}
)
