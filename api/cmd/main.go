package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
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

	// graceful shutdown
	mux := router.NewRouter(db)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctxTimeout)
	}()

	err = srv.ListenAndServe()
	wg.Wait()
	if !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
