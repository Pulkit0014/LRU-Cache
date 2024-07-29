package routers

import (
	"encoding/json"
	"fmt"
	"lru-cache/cache"
	"lru-cache/websocket"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside getHandlewr")
	key := mux.Vars(r)["key"]
	fmt.Println("key is ", key)
	value, msg, found := cache.Get(key)
	fmt.Println("The value is 0", value)
	if !found {
		if msg == "Key not found" {
			http.NotFound(w, r)
		} else {
			http.Error(w, msg, http.StatusGone)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"key": key, "value": value})
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	var request map[string]interface{}
	json.NewDecoder(r.Body).Decode(&request)
	key := request["key"].(string)
	value := request["value"]
	duration := time.Duration(request["duration"].(float64)) * time.Second
	cache.Set(key, value, duration)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Key set successfully"})

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	cache.Delete(key)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Key deleted successfully"})
}
func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	allItems := cache.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allItems)
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/cache/{key}", GetHandler).Methods("GET")
	r.HandleFunc("/cache", SetHandler).Methods("POST")
	r.HandleFunc("/cache/{key}", DeleteHandler).Methods("DELETE")
	r.HandleFunc("/cache", GetAllHandler).Methods("GET")
	r.HandleFunc("/ws", websocket.HandleConnections)
}
