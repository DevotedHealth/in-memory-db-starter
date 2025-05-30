package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strings"
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

  // Basic parsing
  parts := strings.Fields(req.Command)
  if len(parts) == 0 {
    json.NewEncoder(w).Encode(CommandResponse{Output: "ERROR: Empty command"})
    return
  }

  op := strings.ToUpper(parts[0])
  args := parts[1:]

  // Command switch (stubbed)
  var result string
  switch op {
  case "SET":
    // TODO: implement SET [name] [value]
    result = "OK"
  case "GET":
    // TODO: implement GET [name]
    result = "NULL"
  case "DELETE":
    // TODO: implement DELETE [name]
    result = "OK"
  case "COUNT":
    // TODO: implement COUNT [value]
    result = "0"
  case "BEGIN":
    // TODO: begin transaction
    result = "OK"
  case "ROLLBACK":
    // TODO: roll back most recent transaction
    result = "TRANSACTION NOT FOUND"
  case "COMMIT":
    // TODO: commit all open transactions
    result = "OK"
  default:
    result = "ERROR: Unknown command"
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(CommandResponse{Output: result})
}

func main() {
  // You can change this path if your frontend proxy setup differs
  http.HandleFunc("/api/command", commandHandler)

  fmt.Println("Backend listening on http://localhost:3001")
  log.Fatal(http.ListenAndServe(":3001", nil))
}
