package main

import (
	"log"
	"lru-cache/cache"
	"lru-cache/routers"
	"lru-cache/websocket"
	"net/http"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	r := mux.NewRouter()

	cache.InitCache(100) // Initialize cache with capacity 100
	routers.RegisterRoutes(r)

	go websocket.HandleMessages()

	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", enableCORS(r)); err != nil {
		log.Fatal(err)
	}
}
