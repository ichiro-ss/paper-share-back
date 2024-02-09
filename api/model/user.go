package model

type (
	CreateUserRequest struct {
		LoginId  string `json:"loginId"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Token string `json:"token"`
	}

	ReadUserRequest struct {
		Token string `json:"token"`
	}
	ReadUserResponse struct {
		Name string `json:"name"`
	}
	EditUserRequest struct {
		Token string `json:"token"`
	}
	EditUserResponse struct {
		Name string `json:"name"`
	}
)
