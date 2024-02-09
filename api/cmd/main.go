package main

import (
	"errors"
	"net/http"
	"time"

	"api/data"
	"api/handler/router"
)

func main() {
	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	db := data.GetMydb()
	defer db.Close()

	mux := router.NewRouter(db)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err = srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
