# In-Memory Database Starter

This is a minimal project scaffolded for your technical assessment.

## Getting Started

```bash
npm install
npm run dev
```

## Project Structure

Below is the structure that we have provided. You're welcome to change anything about this structure if it helps you get to a solution you're proud of!

CRK TODO

### Where Should I Put My Logic?

You have two options, depending on your preference:

1. Build the logic directly in TypeScript
You can implement the database logic in src/db/store.ts and wire it directly into the React UI.

2. Use another language
You’re welcome to build the logic in a separate service in any language (Python, Go, etc.). Just expose an HTTP endpoint like /api/command that accepts a JSON payload:
```json
{ "command": "SET foo" }
```
and returns a JSON response:
```json
{ "output": "OK" }
```
The frontend is already configured to proxy `/api` to `localhost:3001`.

### Sample Backend in Go

Here's a minimal Go server that handles incoming commands via JSON.

**File: `go-backend/main.go`**

```go
// See go-backend/main.go for full example
```

Run with:
```bash
go run go-backend/main.go
```
