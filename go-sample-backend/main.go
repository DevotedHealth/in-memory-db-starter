package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CommandRequest struct {
	Command string `json:"command"`
}

type CommandResponse struct {
	Output string `json:"output"`
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	var req CommandRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received command: %s\n", req.Command)
	resp := CommandResponse{Output: fmt.Sprintf("Executed: %s", req.Command)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/command", commandHandler)
	fmt.Println("Backend listening on http://localhost:3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
