# Agent Guidelines for Rune Discord Bot (Go Rewrite)

## Commands
- **Build**: `go build` or `go build -o runeDiscordBot`
- **Run**: `go run main.go` (main bot) or `go run .`
- **Test**: `go test ./...` (all packages) or `go test ./workers` (specific package)
- **Format**: `go fmt ./...` (format all code)
- **Vet**: `go vet ./...` (static analysis)

## Code Style
- **Package Structure**: Standard Go package layout with clear separation
- **Naming**: PascalCase for exported, camelCase for unexported
- **Functions**: Public functions start with capital letter, document with comments
- **Error Handling**: Return error as last value, check errors explicitly
- **Structs**: Use struct tags for JSON marshaling (`json:"field_name"`)
- **Variables**: Use `var` for zero values, `:=` for assignment with inference
- **Constants**: Group related constants in blocks with `const ()`

## Architecture
- Main bot logic in `main.go` and `discord-bot/main.go`
- Worker functions in `workers/` package
- API endpoints in `api/` package  
- Go modules for dependency management
- Concurrent API server using goroutines

## Dependencies
- `github.com/bwmarrin/discordgo` for Discord API
- `github.com/joho/godotenv` for environment variables
- Go 1.22+ required
- Standard library for HTTP server and JSON handling