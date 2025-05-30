# In-Memory Database Starter

This is a minimal project scaffolded for your technical assessment. The details of the assessment can be seen in [PROMPT.md](https://github.com/DevotedHealth/in-memory-db-starter/blob/master/PROMPT.md).

## Getting Started

```bash
npm install
npm run dev
```

This will start the frontend on `http://localhost:5173`.

## Project Structure

Below is the structure that we have provided. You're welcome to change anything about this structure if it helps you get to a solution you're proud of!

```
in-memory-db-starter/
├── go-backend/
│   └── main.go                  # Sample Go server
├── src/
│   ├── components/
│   │   ├── CommandInput.tsx     # User input field for commands
│   │   └── OutputDisplay.tsx    # Shows output or history
│   ├── db/
│   │   └── store.ts             # In-memory DB logic (if using TypeScript)
│   ├── rpc/
│   │   └── handlers.ts          # UI-to-logic bridge (call store or backend)
│   ├── App.tsx                  # Main app layout
│   ├── main.tsx                 # React entry point
│   └── index.css                # Styles
├── .gitignore
├── index.html
├── package.json
├── package-lock.json
├── PROMPT.md                    # Assessment instructions for candidates
├── README.md                    # Project guide
├── tsconfig.json
└── vite.config.ts
```

### Where Should I Put My DB Logic?

You can implement your database logic in one of two ways:

#### Option 1: TypeScript (in-browser)
Implement your logic directly in `src/db/store.ts` and wire it into the UI using `handleCommand()` (stub is in `src/rpc/handlers.ts`).

#### Option 2: HTTP Backend (any language)
Implement your logic in a separate backend (like Go, Python, etc.) that listens on `localhost:3001` and exposes a `POST /command` endpoint. The frontend is already configured to proxy to that port. A sample Go backend is included in `go-sample-backend/main.go`

> Note: This Go backend is just an example. You're welcome to implement your backend logic in any language as long as it listens on `localhost:3001` and responds to `POST /command` with the expected JSON structure.

```go
// See go-sample-backend/main.go for full example
```

Run with:
```bash
go run go-sample-backend/main.go
```

The frontend expects to send:

```json
{
  "command": "SET a foo"
}
```

And receive:

```json
{
  "output": "OK"
}
```
