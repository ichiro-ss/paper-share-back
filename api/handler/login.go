package handler

import (
	"api/model"
	"api/service"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type LoginHandler struct {
	svc *service.LoginService
}

func NewLoginHandler(svc *service.LoginService) *LoginHandler {
	return &LoginHandler{
		svc: svc,
	}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// POST
	if r.Method == http.MethodPost {
		var loginReq model.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
			log.Println(err)
			return
		}
		if loginReq.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			loginRes, err := h.Create(r.Context(), &loginReq)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				log.Println(err)
				return
			}
			if err := json.NewEncoder(w).Encode(*loginRes); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
func (h *LoginHandler) Create(ctx context.Context, req *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	tk, err := h.svc.Login(ctx, req.LoginId, req.Password)
	if err != nil {
		return nil, err
	}
	var loginRes model.CreateUserResponse
	loginRes.Token = tk
	return &loginRes, nil
}
