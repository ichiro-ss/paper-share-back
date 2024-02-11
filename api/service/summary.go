package service

import "database/sql"

type SummaryService struct {
	db *sql.DB
}

func NewSummaryService(db *sql.DB) *SummaryService {
	return &SummaryService{
		db: db,
	}
}
