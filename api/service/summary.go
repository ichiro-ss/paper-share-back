package service

import (
	"api/model"
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

func (s *SummaryService) ReadSummary(ctx context.Context, token string, id int) (*model.ReadSummaryResponse, error) {
	var readSummaryRes model.ReadSummaryResponse
	statement := fmt.Sprintf("SELECT * from %s WHERE id = ?", tableSummary)
	prep, err := s.db.Prepare(statement)
	if err != nil {
		return &readSummaryRes, err
	}
	defer prep.Close()

	err = prep.QueryRowContext(ctx, id).Scan(&readSummaryRes.Id, &readSummaryRes.UserId, &readSummaryRes.Title, &readSummaryRes.Markdown)
	if err != nil {
		return &readSummaryRes, err
	}

	userId, err := TokenToId(token)
	if err != nil {
		return &model.ReadSummaryResponse{}, err
	}
	if userId == int64(readSummaryRes.UserId) {
		readSummaryRes.IsMine = true
	} else {
		readSummaryRes.IsMine = false
	}

	return &readSummaryRes, nil
}
