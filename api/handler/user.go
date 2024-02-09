package handler

import (
	"api/model"
	"api/service"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserHandler struct {
	svc *service.UserService
}

// NewUserHandler returns UserHandler based http.Handler.
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var createUserReq model.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&createUserReq); err != nil {
			log.Println(err)
			return
		}
		if createUserReq.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			createUserRes, err := h.Create(r.Context(), &createUserReq)
			if err != nil {
				log.Println(err)
				return
			}
			if err := json.NewEncoder(w).Encode(*createUserRes); err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("type of createUserRes: %T\n", createUserRes.Token)
			fmt.Println("value:", createUserRes.Token)
			res, err := json.Marshal(createUserRes)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("type of res%T\n", res)
			fmt.Println(string(res))
			w.Write(res)
		}
	}
}

func (h *UserHandler) Create(ctx context.Context, req *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	tk, err := h.svc.CreateUser(ctx, req.LoginId, req.Password)
	if err != nil {
		return nil, err
	}
	var createUserRes model.CreateUserResponse
	createUserRes.Token = tk
	return &createUserRes, nil
}
