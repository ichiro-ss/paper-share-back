package service

import (
	"context"
	"database/sql"
	"fmt"
)

const tableSummary = "summaries"
const titleCol = "title"
const markdownCol = "markdown"

type SummaryService struct {
	db *sql.DB
}

func NewSummaryService(db *sql.DB) *SummaryService {
	return &SummaryService{
		db: db,
	}
}

func (s *SummaryService) CreateSummary(ctx context.Context, token, title, mk string) error {
	id, err := TokenToId(token)
	if err != nil {
		return err
	}
	statement := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES (?, ?, ?)", tableSummary, userIdCol, titleCol, markdownCol)
	prep, err := s.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer prep.Close()

	_, err = prep.ExecContext(ctx, id, title, mk)
	if err != nil {
		return err
	}
	return nil
}
