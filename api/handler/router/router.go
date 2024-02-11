package router

import (
	"api/handler"
	"api/service"
	"database/sql"
	"net/http"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)
	mux.HandleFunc("/users", userHandler.ServeHTTP)

	summaryService := service.NewSummaryService(db)
	summaryHandler := handler.NewSummaryHandler(summaryService)
	mux.HandleFunc("/summaries", summaryHandler.ServeHTTP)

	return mux
}
