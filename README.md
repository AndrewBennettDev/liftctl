# `liftctl`

A terminal-based workout tracking application written in Go. Features a modern TUI for managing exercises, creating workout routines, and tracking workout sessions. All data is stored locally using SQLite with GORM as the ORM.

## Features

- **Interactive TUI**: Modern terminal interface built with Bubble Tea
- **Exercise Management**: Add and organize exercises by muscle groups
- **Routine Creation**: Build custom workout routines with exercises, sets, reps, and weights
- **Live Workout Tracking**: Start workouts from routines and track progress in real-time
- **Set Management**: Edit reps and weights during workouts, mark sets as complete
- **Workout History**: Complete and manage workout sessions
- **Local Storage**: All data stored locally in SQLite database

## Technology Stack

- **Go** - Core application language
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - TUI framework for interactive terminal interface
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)** - Terminal styling and layout
- **[Cobra](https://github.com/spf13/cobra)** - CLI command structure
- **[GORM](https://gorm.io/)** - ORM for database operations
- **SQLite** - Local data persistence

## Installation

1. Ensure you have [Go](https://golang.org/dl/) installed (version 1.18 or later recommended).

2. Clone this repo:

   ```bash
   git clone https://github.com/AndrewBennettDev/liftctl.git
   cd liftctl
   ```

3. Build the application:

   ```bash
   go build -o liftctl
   ```

4. Run the TUI:

   ```bash
   ./liftctl tui
   ```

## Usage

### TUI Mode

Launch the TUI:

```bash
./liftctl tui
```

#### Navigation
- **Arrow Keys** or **j/k**: Navigate up/down through options
- **Enter**: Select/confirm
- **ESC**: Go back or cancel
- **q**: Quit application
- **b**: Back to main menu
- **t**: Toggle set completion (in active workouts)

#### Workflow
1. **Create Exercises**: Add exercises with muscle group targeting
2. **Build Routines**: Create workout routines by adding exercises with planned sets, reps, and weights
3. **Start Workouts**: Begin a workout session from any routine
4. **Track Progress**: Edit weights/reps during workouts, mark sets complete
5. **Manage Sessions**: Complete or delete workout sessions

### CLI Commands

Basic CLI commands are also available:

#### Add Exercise

Add a new exercise to the database:

```bash
./liftctl add-exercise --name "Bench Press" --muscle "chest"
```

#### List Exercises

List all exercises in the database:

```bash
./liftctl list-exercises
```

For a full list of commands simply run the `./liftctl` command!


## Development

To contribute or modify the application:

```bash
# Install dependencies
go mod download

# Run in development
go run main.go tui

# Build for production
go build -o liftctl
```

## License

MIT License © 2025 Andrew Bennett
