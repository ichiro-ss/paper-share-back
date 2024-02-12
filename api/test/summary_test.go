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

func TestSummary(t *testing.T) {
	t.Run("creating summary", func(t *testing.T) {
		tkStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc4MTA4ODIsInVzZXJfaWQiOjE3fQ.uuswOlREnkOL7GD4pyRt5Deg-klNrvH9hMIC_l92P4w"
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
}