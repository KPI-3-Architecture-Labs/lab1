package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := map[string]string{
			"time": time.Now().Format(time.RFC3339),
		}
		jsonResponse, err := json.Marshal(currentTime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonResponse)
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":8795", nil)
	if err != nil {
		return
	}
}
