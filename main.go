package main

import (
	"log"
	"net/http"

	"github.com/utkarsh352/go-chat/models"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	hub := models.NewHub()
	go hub.Run()

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		models.ServeWs(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
