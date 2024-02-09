package model

type (
	CreateUserRequest struct {
		LoginId  string `json:"loginId"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Token string `json:"token"`
	}
	EditUserRequest struct {
		token string `json:"token"`
	}
	EditUserResponse struct {
		name string `json:"name`
	}
)
