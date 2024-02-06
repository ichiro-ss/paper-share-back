package handler

import (
	"api/model"
	"api/service"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
	svc *service.UserService
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
			createTodoRes, err := h.Create(r.Context(), &createUserReq)
			if err != nil {
				log.Println(err)
				return
			}
			if err := json.NewEncoder(w).Encode(*createTodoRes); err != nil {
				log.Println(err)
				return
			}
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
