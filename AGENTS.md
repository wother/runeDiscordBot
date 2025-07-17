# Agent Guidelines for Rune Discord Bot (Go Rewrite)

## Commands
- **Build All**: `go build ./cmd/...`
- **Build Specific**: `go build -o runeDiscordBot-api ./cmd/api`, `go build -o runeDiscordBot-discord ./cmd/discord`, `go build -o runeDiscordBot-tui ./cmd/tui`
- **Run**: `./runeDiscordBot-api`, `./runeDiscordBot-discord`, `./runeDiscordBot-tui` (after building)
- **Test**: `go test ./...` (all packages) or `go test ./internal/workers` (specific package)
- **Format**: `go fmt ./...` (format all code)
- **Vet**: `go vet ./...` (static analysis)
- **Tidy Modules**: `go mod tidy` (important after dependency changes or refactoring)

## Code Style
- **Package Structure**: Standard Go package layout with `cmd/` for executables and `internal/` for shared, internal libraries.
- **Naming**: PascalCase for exported, camelCase for unexported
- **Functions**: Public functions start with capital letter, document with comments
- **Error Handling**: Return error as last value, check errors explicitly
- **Structs**: Use struct tags for JSON marshaling (`json:"field_name"`)
- **Variables**: Use `var` for zero values, `:=` for assignment with inference
- **Constants**: Group related constants in blocks with `const ()`

## Architecture
- Multiple application entry points (API, Discord Bot, TUI) are located under the `cmd/` directory.
- Shared business logic and utility functions are encapsulated in packages under the `internal/` directory (e.g., `internal/workers`, `internal/api`).
- Go modules for dependency management.
- Concurrent API server using goroutines.

## Dependencies
- `github.com/bwmarrin/discordgo` for Discord API
- `github.com/joho/godotenv` for environment variables
- `github.com/charmbracelet/bubbletea` for TUI development
- Go 1.23+ required (due to `bubbletea` and updated `math/rand` usage).
- Standard library for HTTP server and JSON handling.

## Lessons Learned (for future agents)
- **Go Project Structure**: For projects with multiple executables, always favor the `cmd/` and `internal/` directory structure. It provides clear separation and adheres to Go best practices.
- **`replace` Tool Precision**: When using the `replace` tool, the `old_string` must be an *exact* match, including all whitespace, newlines, and surrounding context. Even a single character difference will cause it to fail. For large, complex blocks, consider reading the file, modifying content in memory, and then writing the entire file back.
- **Basic Syntax Errors**: Be vigilant for fundamental syntax errors (e.g., missing quotes in import paths). These can sometimes lead to misleading compiler errors (like "string literal not terminated") that point to the wrong line or obscure the true cause.
- **Empty Go Files**: Newly created `.go` files, even if empty, *must* have a `package` declaration at the top. Otherwise, `go build` will fail with an "expected 'package', found 'EOF'" error.
- **Build Output Naming**: When building executables, explicitly specify the output name using `-o` to avoid conflicts with existing directories (e.g., `go build -o myapp ./cmd/myapp`).
- **Random Number Generation**: Be aware of changes in Go's standard library. `rand.Seed` is deprecated in Go 1.20+. Use `rand.NewSource` and `rand.New` for proper random number generation, and ensure the `*rand.Rand` instance is passed to functions that require randomness.
- **Verify All Interfaces After Refactor**: After significant structural changes (like moving packages), always perform a full build and, if possible, run tests for *all* application interfaces (API, Discord bot, TUI) to ensure nothing was broken in the process. Don't assume functionality based on partial builds.
