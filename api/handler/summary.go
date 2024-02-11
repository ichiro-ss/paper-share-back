package handler

import (
	"api/service"
	"net/http"
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

}
