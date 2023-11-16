package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		var message Message
		if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		response := map[string]string{"received": message.Text}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":3000", r)
}
