package router

import (
	"api/handler"
	"api/handler/middleware"
	"api/service"
	"database/sql"
	"net/http"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	userService := service.NewUserService(db)
	userHandler := middleware.CORS(handler.NewUserHandler(userService))
	mux.HandleFunc("/users", userHandler.ServeHTTP)

	summaryService := service.NewSummaryService(db)
	summaryHandler := middleware.CORS(handler.NewSummaryHandler(summaryService))
	mux.HandleFunc("/summaries", summaryHandler.ServeHTTP)

	loginService := service.NewLoginService(db)
	loginHandler := middleware.CORS(handler.NewLoginHandler(loginService))
	mux.HandleFunc("/login", loginHandler.ServeHTTP)

	return mux
}
