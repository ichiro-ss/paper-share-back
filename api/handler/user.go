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

// NewUserHandler returns UserHandler based http.Handler.
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// POST
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
		}
	}
	// GET
	if r.Method == http.MethodGet {
		var readUserReq model.ReadUserRequest
		if err := json.NewDecoder(r.Body).Decode(&readUserReq); err != nil {
			log.Println(err)
			return
		} else {
			readUserRes, err := h.Read(r.Context(), &readUserReq)
			if err != nil {
				log.Println(err)
				return
			}
			if err := json.NewEncoder(w).Encode(*readUserRes); err != nil {
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

func (h *UserHandler) Read(ctx context.Context, req *model.ReadUserRequest) (*model.ReadUserResponse, error) {
	n, err := h.svc.ReadUser(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	var readUserRes model.ReadUserResponse
	readUserRes.Name = n
	return &readUserRes, nil
}
