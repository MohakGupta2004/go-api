package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello bitches"))
	})
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
