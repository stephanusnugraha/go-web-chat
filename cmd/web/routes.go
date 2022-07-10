package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"stephanusnugraha/go-webchat/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	return mux
}