package handler

import (
	"api/model"
	"api/service"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type SummaryHandler struct {
	svc *service.SummaryService
}

func NewSummaryHandler(svc *service.SummaryService) *SummaryHandler {
	return &SummaryHandler{
		svc: svc,
	}
}

func (h *SummaryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// POST
	if r.Method == http.MethodPost {
		var createSummaryReq model.CreateSummaryRequest
		if err := json.NewDecoder(r.Body).Decode(&createSummaryReq); err != nil {
			log.Println(err)
			return
		} else {
			createSummaryReq.Token = strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			err := h.Create(r.Context(), &createSummaryReq)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
	if r.Method == http.MethodGet {
		var readSummaryReq model.ReadSummaryRequest
		readSummaryReq.Token = strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		summaryId, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Println(err)
			return
		}
		readSummaryReq.Id = summaryId
		readSummaryRes, err := h.Read(r.Context(), &readSummaryReq)
		if err != nil {
			log.Println(err)
			return
		}
		if err := json.NewEncoder(w).Encode(*readSummaryRes); err != nil {
			log.Println(err)
			return
		}

	}
}

func (h *SummaryHandler) Create(ctx context.Context, req *model.CreateSummaryRequest) error {
	err := h.svc.CreateSummary(ctx, req.Token, req.Title, req.Markdown)
	if err != nil {
		return err
	}

	return nil
}

func (h *SummaryHandler) Read(ctx context.Context, req *model.ReadSummaryRequest) (*model.ReadSummaryResponse, error) {
	readSummaryRes, err := h.svc.ReadSummary(ctx, req.Token, req.Id)
	if err != nil {
		return nil, err
	}
	return readSummaryRes, nil
}
