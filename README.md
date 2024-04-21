# ideaboard-go-svelte

Ideaboard application built using Gorilla and Svelte.

## Setting up in local:

1. Start Svelte app inside the client folder:
   ```
   cd client
   yarn dev
   ```
2. Setup `.env` file at project root (refer `.env.example`)
3. Start Golang server at project root:
   ```
   DEBUG=true go run main.go
   ```
