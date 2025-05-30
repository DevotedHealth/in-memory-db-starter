package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

// CommandRequest represents a command sent by the frontend UI.
// For example: { "command": "SET a foo" }
type CommandRequest struct {
  Command string `json:"command"`
}

// CommandResponse represents the result of executing a command.
// For example: { "output": "OK" }
type CommandResponse struct {
  Output string `json:"output"`
}

// commandHandler receives a command string from the frontend,
// processes it (parsing and executing), and returns the result.
func commandHandler(w http.ResponseWriter, r *http.Request) {
  var req CommandRequest
  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
  }

  fmt.Printf("Received command: %s\n", req.Command)

  // TODO: Parse the command string and update your in-memory state.
  // This is just a placeholder response:
  resp := CommandResponse{Output: fmt.Sprintf("Executed: %s", req.Command)}

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(resp)
}

func main() {
  // You can change this path if your frontend proxy setup differs
  http.HandleFunc("/api/command", commandHandler)

  fmt.Println("Backend listening on http://localhost:3001")
  log.Fatal(http.ListenAndServe(":3001", nil))
}
