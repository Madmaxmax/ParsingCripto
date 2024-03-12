package main

import (
	"net/http"
	"os"
	"stepik_go/handlers"
	"stepik_go/money"
)

func main() {
	go money.GetMoneyInfo()
	go money.UpdateMoneyInfo()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets", http.StripPrefix("assets", fs))

	mux.HandleFunc("/search", handlers.SearchHandler)
	mux.HandleFunc("/", handlers.IndexHandler)
	http.ListenAndServe(":"+port, mux)
}
