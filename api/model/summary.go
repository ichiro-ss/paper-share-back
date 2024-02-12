package model

type (
	CreateSummaryRequest struct {
		Token    string `json:"token"`
		Title    string `json:"title"`
		Markdown string `json:"markdown"`
	}

	ReadSummaryRequest struct {
		Token string `json:"token"`
		Id    int    `json:"id"`
	}
	ReadSummaryResponse struct {
		Id       int    `json:"id"`
		UserId   int    `json:"userId"`
		Title    string `json:"title"`
		Markdown string `json:"markdown"`
		IsMine   bool   `json:"isMine"`
	}

	EditSummaryRequest struct {
		Token    string `json:"token"`
		Id       int    `json:"Id"`
		Title    string `json:"title"`
		Markdown string `json:"markdown"`
	}
	EditSummaryResponse struct {
		Id       int    `json:"id"`
		UserId   int    `json:"userId"`
		Title    string `json:"title"`
		Markdown string `json:"markdown"`
		IsMine   bool   `json:"isMine"`
	}

	DeleteSummaryRequest struct {
		Token string `json:"token"`
		Id    int    `json:"id"`
	}
)
