package test

import (
	"api/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

const tkStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc4MTA4ODIsInVzZXJfaWQiOjE3fQ.uuswOlREnkOL7GD4pyRt5Deg-klNrvH9hMIC_l92P4w"

func TestSummary(t *testing.T) {
	t.Run("creating summary", func(t *testing.T) {
		createSummaryReq := model.CreateSummaryRequest{
			Title:    "interesting paper",
			Markdown: "## this is interesting paper",
		}
		summaryJson, err := json.Marshal(createSummaryReq)
		if err != nil {
			t.Error(err)
		}
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://localhost:8080/summaries", bytes.NewBuffer(summaryJson))
		if err != nil {
			t.Error(err)
		}
		bearer := fmt.Sprintf("Bearer %s", tkStr)
		req.Header.Set("Authorization", bearer)

		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()

		_, err = io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}
		fmt.Println("status code:", res.StatusCode)
	})

	t.Run("reading summary", func(t *testing.T) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://localhost:8080/summaries", nil)
		if err != nil {
			t.Error(err)
		}
		bearer := fmt.Sprintf("Bearer %s", tkStr)
		req.Header.Set("Authorization", bearer)

		q := req.URL.Query()
		q.Set("id", "1")
		req.URL.RawQuery = q.Encode()

		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(string(body))
	})

	t.Run("edit summary info", func(t *testing.T) {
		editSummaryReq := model.EditSummaryRequest{
			Title:    "edited",
			Markdown: "## Edited",
		}
		SummaryJson, err := json.Marshal(editSummaryReq)
		if err != nil {
			t.Error(err)
		}
		client := &http.Client{}
		req, err := http.NewRequest("PUT", "http://localhost:8080/summaries", bytes.NewBuffer(SummaryJson))
		if err != nil {
			t.Error(err)
		}
		bearer := fmt.Sprintf("Bearer %s", tkStr)
		req.Header.Set("Authorization", bearer)

		q := req.URL.Query()
		q.Set("id", "1")
		req.URL.RawQuery = q.Encode()

		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil || body == nil {
			t.Error(err)
		}
		fmt.Println(string(body))
	})
}
