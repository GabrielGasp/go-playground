package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func printCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	w.Write([]byte(code))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func main() {
	r := chi.NewRouter()
	r.Get("/{code}", printCode)
	r.Get("/health", healthCheck)
	http.ListenAndServe(":3000", r)
}
