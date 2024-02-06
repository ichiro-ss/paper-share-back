package model

type (
	CreateUserRequest struct {
		LoginId  string `json:"loginId"`
		Password string `json:"password"`
	}
)
