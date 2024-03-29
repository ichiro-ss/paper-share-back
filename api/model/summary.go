package model

type (
	Summary struct {
		Id       int      `json:"id"`
		UserId   int      `json:"userId"`
		Title    string   `json:"title"`
		Markdown string   `json:"markdown"`
		Authors  []string `json:"authors"`
		IsMine   bool     `json:"isMine"`
	}

	CreateSummaryRequest struct {
		Token    string   `json:"token"`
		Title    string   `json:"title"`
		Markdown string   `json:"markdown"`
		Authors  []string `json:"authors"`
	}

	ReadSummaryRequest struct {
		Token string `json:"token"`
		Id    int    `json:"id"`
	}
	ReadSummaryResponse struct {
		Summaries []*Summary `json:"summaries"`
	}

	EditSummaryRequest struct {
		Token    string   `json:"token"`
		Id       int      `json:"Id"`
		Title    string   `json:"title"`
		Markdown string   `json:"markdown"`
		Authors  []string `json:"authors"`
	}
	EditSummaryResponse struct {
		Id       int      `json:"id"`
		UserId   int      `json:"userId"`
		Title    string   `json:"title"`
		Markdown string   `json:"markdown"`
		Authors  []string `json:"authors"`
		IsMine   bool     `json:"isMine"`
	}

	DeleteSummaryRequest struct {
		Token string `json:"token"`
		Id    int    `json:"id"`
	}
)
