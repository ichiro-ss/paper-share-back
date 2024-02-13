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
	// GET
	if r.Method == http.MethodGet {
		var readSummaryReq model.ReadSummaryRequest
		readSummaryReq.Token = strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		q := r.URL.Query()
		id, isExist := q["id"]
		if isExist && len(id) != 0 {
			readSummaryReq.Id, _ = strconv.Atoi(r.URL.Query().Get("id"))
		} else {
			readSummaryReq.Id = 0
		}
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
	// EDIT
	if r.Method == http.MethodPut {
		var editSummaryReq model.EditSummaryRequest
		if err := json.NewDecoder(r.Body).Decode(&editSummaryReq); err != nil {
			log.Println(err)
			return
		} else {
			editSummaryReq.Token = strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			summaryId, err := strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil || summaryId == 0 {
				log.Println(err)
				return
			}
			editSummaryReq.Id = summaryId
			editSummaryRes, err := h.Edit(r.Context(), &editSummaryReq)
			if err != nil {
				log.Println(err)
				return
			}
			if err := json.NewEncoder(w).Encode(*editSummaryRes); err != nil {
				log.Println(err)
				return
			}
		}
	}
	// UPDATE
	if r.Method == http.MethodDelete {
		var deleteSummaryReq model.DeleteSummaryRequest
		deleteSummaryReq.Token = strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		summaryId, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || summaryId == 0 {
			log.Println(err)
			return
		}
		deleteSummaryReq.Id = summaryId
		err = h.Delete(r.Context(), &deleteSummaryReq)
		if err != nil {
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
	summaries, err := h.svc.ReadSummary(ctx, req.Token, req.Id)
	if err != nil {
		return nil, err
	}
	return &model.ReadSummaryResponse{Summaries: summaries}, nil
}

func (h *SummaryHandler) Edit(ctx context.Context, req *model.EditSummaryRequest) (*model.EditSummaryResponse, error) {
	editSummaryRes, err := h.svc.EditSummary(ctx, req.Token, req.Title, req.Markdown, req.Id)
	if err != nil {
		return nil, err
	}
	return editSummaryRes, nil
}

func (h *SummaryHandler) Delete(ctx context.Context, req *model.DeleteSummaryRequest) error {
	err := h.svc.DeleteSummary(ctx, req.Token, req.Id)
	if err != nil {
		return err
	}

	return nil
}
