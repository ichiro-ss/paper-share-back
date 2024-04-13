package service

import (
	"api/model"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const tableSummary = "summaries"
const tableAuthor = "authors"
const tableSummaryAuthor = "summary_authors"
const titleCol = "title"
const markdownCol = "markdown"
const authorNameCol = "name"
const summaryIdCol = "summaryId"
const authorIdCol = "authorId"

type SummaryService struct {
	db *sql.DB
}

func NewSummaryService(db *sql.DB) *SummaryService {
	return &SummaryService{
		db: db,
	}
}

// CreateSummary creates a new summary (+ authors + summary_authors map).
func (s *SummaryService) CreateSummary(ctx context.Context, token, title, mk string, authors []string) error {
	id, err := TokenToId(token)
	if err != nil {
		return err
	}
	// insert summary
	statement := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES (?, ?, ?)", tableSummary, userIdCol, titleCol, markdownCol)
	prep, err := s.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer prep.Close()

	res, err := prep.ExecContext(ctx, id, title, mk)
	if err != nil {
		return err
	}
	lid, err := res.LastInsertId()
	if err != nil {
		return err
	}
	summaryId := int(lid)

	// insert authors and summary_authors map
	selectAuthor := fmt.Sprintf("SELECT id FROM %s WHERE name = ?", tableAuthor)
	insertAuthor := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?)", tableAuthor, authorNameCol)
	insertMap := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES (?, ?)", tableSummaryAuthor, summaryIdCol, authorIdCol)
	prep_select, err := s.db.Prepare(selectAuthor)
	if err != nil {
		return err
	}
	defer prep_select.Close()
	prep_in_author, err := s.db.Prepare(insertAuthor)
	if err != nil {
		return err
	}
	defer prep_in_author.Close()
	prep_in_map, err := s.db.Prepare(insertMap)
	if err != nil {
		return err
	}
	defer prep_in_map.Close()
	for _, author := range authors {
		authorId := 0
		err := prep_select.QueryRowContext(ctx, author).Scan(&authorId)
		// if author doesn't exist, insert author
		if err != nil && err == sql.ErrNoRows {
			res, err_in := prep_in_author.ExecContext(ctx, author)
			if err_in != nil {
				return err_in
			}
			lid, err_id := res.LastInsertId()
			if err_id != nil {
				return err_id
			}
			authorId = int(lid)
		} else if err != nil {
			return err
		}
		_, err = prep_in_map.ExecContext(ctx, summaryId, authorId)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadSummary reads summaries.
func (s *SummaryService) ReadSummary(ctx context.Context, token string, id int) ([]*model.Summary, error) {
	var summaries []*model.Summary

	// get summaries and authors info using JOIN
	query := fmt.Sprintf("SELECT s.id, s.userId, s.title, s.markdown, "+
		"GROUP_CONCAT(a.name SEPARATOR '$') AS authors "+
		"FROM %s s "+
		"LEFT JOIN %s sa ON s.id = sa.summaryId "+
		"LEFT JOIN %s a ON sa.authorId = a.id", tableSummary, tableSummaryAuthor, tableAuthor)

	var rows *sql.Rows
	var err error
	if id != 0 {
		query = fmt.Sprintf("%s WHERE s.id = ? GROUP BY s.id", query)
		rows, err = s.db.QueryContext(ctx, query, id)
	} else {
		query = query + " GROUP BY s.id"
		rows, err = s.db.QueryContext(ctx, query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userId, err := TokenToId(token)
	if err != nil {
		return nil, err
	}

	// split authors by '$'
	for rows.Next() {
		var summary model.Summary
		var authors sql.NullString
		err := rows.Scan(&summary.Id, &summary.UserId, &summary.Title, &summary.Markdown, &authors)
		if err != nil {
			return nil, err
		}
		summary.Authors = strings.Split(authors.String, "$")
		fmt.Println(summary)

		if userId == int64(summary.UserId) {
			summary.IsMine = true
		} else {
			summary.IsMine = false
		}
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

// EditSummary edits a summary.
func (s *SummaryService) EditSummary(ctx context.Context, token, title, markdown string, id int, authors []string) (*model.EditSummaryResponse, error) {
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
		// edit authors
		selectAuthor := fmt.Sprintf("SELECT id FROM %s WHERE name = ?", tableAuthor)
		insertAuthor := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?)", tableAuthor, authorNameCol)
		deleteMap := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", tableSummaryAuthor, summaryIdCol)
		insertMap := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES (?, ?)", tableSummaryAuthor, summaryIdCol, authorIdCol)
		prep_select, err := s.db.Prepare(selectAuthor)
		if err != nil {
			return nil, err
		}
		defer prep_select.Close()
		prep_in_author, err := s.db.Prepare(insertAuthor)
		if err != nil {
			return nil, err
		}
		defer prep_in_author.Close()
		prep_del_map, err := s.db.Prepare(deleteMap)
		if err != nil {
			return nil, err
		}
		defer prep_del_map.Close()
		prep_in_map, err := s.db.Prepare(insertMap)
		if err != nil {
			return nil, err
		}
		defer prep_in_map.Close()

		// delete summary_authors map
		_, err = prep_del_map.ExecContext(ctx, id)
		if err != nil {
			return nil, err
		}

		// change authors
		for _, author := range authors {
			authorId := 0
			err := prep_select.QueryRowContext(ctx, author).Scan(&authorId)
			// if author doesn't exist, insert author
			if err != nil && err == sql.ErrNoRows {
				res, err_in := prep_in_author.ExecContext(ctx, author)
				if err_in != nil {
					return nil, err_in
				}
				lid, err_id := res.LastInsertId()
				if err_id != nil {
					return nil, err_id
				}
				authorId = int(lid)
			} else if err != nil {
				return nil, err
			}
			_, err = prep_in_map.ExecContext(ctx, id, authorId)
			if err != nil {
				return nil, err
			}
		}

		editSummaryRes.Id = readRes[0].Id
		editSummaryRes.UserId = readRes[0].UserId
		editSummaryRes.Title = title
		editSummaryRes.Markdown = markdown
		editSummaryRes.Authors = authors
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
