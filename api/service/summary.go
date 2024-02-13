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

func (s *SummaryService) ReadSummary(ctx context.Context, token string, id int) ([]*model.Summary, error) {
	var summaries []*model.Summary
	readAll := fmt.Sprintf("SELECT * from %s ORDER BY id", tableSummary)
	readWID := fmt.Sprintf("SELECT * from %s WHERE id = ?", tableSummary)

	var rows *sql.Rows
	var err error
	if id == 0 {
		rows, err = s.db.QueryContext(ctx, readAll)
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = s.db.QueryContext(ctx, readWID, id)
		if err != nil {
			return nil, err
		}
	}

	userId, err := TokenToId(token)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var summary model.Summary
		err := rows.Scan(&summary.Id, &summary.UserId, &summary.Title, &summary.Markdown)
		if err != nil {
			return nil, err
		}
		if userId == int64(summary.UserId) {
			summary.IsMine = true
		} else {
			summary.IsMine = false
		}
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

func (s *SummaryService) EditSummary(ctx context.Context, token, title, markdown string, id int) (*model.EditSummaryResponse, error) {
	var editSummaryRes model.EditSummaryResponse
	statement := fmt.Sprintf("UPDATE %s SET %s=?, %s=? WHERE id=?", tableSummary, titleCol, markdownCol)
	prep, err := s.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer prep.Close()

	readRes, err := s.ReadSummary(ctx, token, id)
	if err != nil {
		return nil, err
	}
	if readRes[0].IsMine {
		_, err = prep.ExecContext(ctx, title, markdown, id)
		if err != nil {
			return nil, err
		}
		editSummaryRes.Id = readRes[0].Id
		editSummaryRes.UserId = readRes[0].UserId
		editSummaryRes.Title = title
		editSummaryRes.Markdown = markdown
		editSummaryRes.IsMine = true
	} else {
		return nil, fmt.Errorf("this token doesn't match")
	}

	return &editSummaryRes, nil
}

func (s *SummaryService) DeleteSummary(ctx context.Context, token string, id int) error {
	statement := fmt.Sprintf("DELETE FROM %s WHERE id=?", tableSummary)
	prep, err := s.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer prep.Close()

	readRes, err := s.ReadSummary(ctx, token, id)
	if err != nil {
		return err
	}
	if readRes[0].IsMine {
		_, err = prep.ExecContext(ctx, id)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("this token doesn't match")
	}
	return nil
}
