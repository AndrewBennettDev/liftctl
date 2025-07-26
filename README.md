# `liftctl`

A command-line interface (CLI) tool written in Go for tracking exercises, creating workout routines, and managing workout sessions. The data is stored locally using SQLite with GORM as the ORM.

> **Work in Progress:** This project is actively being developed. Some commands and features are incomplete or planned for future releases. Use accordingly.

## Features

- Add and list exercises
- Create workout routines and associate exercises (coming soon)
- Start workout routines and track progress (planned)

## Technology Stack

- Go
- [Cobra](https://github.com/spf13/cobra) for CLI commands
- [GORM](https://gorm.io/) ORM for SQLite integration
- SQLite for local data persistence

## Installation

1. Ensure you have [Go](https://golang.org/dl/) installed (version 1.18 or later recommended).

2. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/workout-tracker.git
   cd workout-tracker
   ```

3. Build the CLI binary:

   ```bash
   go build -o liftctl
   ```

4. Run the CLI:

   ```bash
   ./liftctl
   ```

## Usage

Run commands with:

```bash
./liftctl <command> [flags]
```

### Available Commands

#### Add Exercise

Add a new exercise to the database.

```bash
./liftctl add-exercise --name "Bench Press" --muscle "chest"
```

- `--name` (required): Name of the exercise  
- `--muscle` (optional): Target muscle group

#### List Exercises

List all exercises stored in the database.

```bash
./liftctl list-exercises
```

This command displays all exercises with their ID, name, and muscle group.

---

## Project Structure

```
liftctl/
├── cmd/               # CLI command definitions
├── internal/          # Internal packages (database, models)
│   ├── db/
│   └── models/
├── main.go            # Entry point
├── go.mod
└── workout.db         # SQLite database (created on first run)
```

## Future Work

- Implement routines management commands (`create-routine`, `add-to-routine`)  
- Implement starting and tracking workout sessions (`start-routine`)  
- Add update and delete commands for exercises and routines  
- Improve CLI UX and error handling

## License

MIT License © 2025 Andrew Bennett
