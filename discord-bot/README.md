# Discord Bot

This project is a Discord bot built using Go (Golang). It serves as an entry point for interacting with the Discord API and handling various bot functionalities.

## Project Structure

- **main.go**: The entry point of the application. Initializes the server and sets up routes and handlers for the Discord bot.
- **internal/discord/bot.go**: Contains the implementation of the Discord bot logic, including request handling and bot behavior management.
- **models/discord.go**: Defines data structures and models used by the Discord bot, such as message formats and user data.
- **pkg/utils/logger.go**: Provides utility functions for logging within the application, standardizing log output and managing log levels.
- **go.mod**: Module definition for the Go application, specifying the module's name and dependencies.
- **go.sum**: Contains checksums for the module's dependencies to ensure correct versions are used.

## Setup Instructions

1. **Clone the repository**:
   ```
   git clone <repository-url>
   cd discord-bot
   ```

2. **Install dependencies**:
   ```
   go mod tidy
   ```

3. **Run the application**:
   ```
   go run main.go
   ```

## Usage Guidelines

- Ensure you have a valid Discord bot token and set it in your environment variables.
- Customize the bot's behavior by modifying the logic in `internal/discord/bot.go`.
- Use the logger utility in `pkg/utils/logger.go` for standardized logging throughout the application.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.