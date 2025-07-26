package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))
	createItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("36")).Bold(true)
)

func (m model) View() string {
	switch m.currentView {
	case ViewMainMenu:
		return m.viewMainMenu()
	case ViewExercises:
		return m.viewExerciseList()
	case ViewRoutines:
		return m.viewRoutineList()
	case ViewWorkouts:
		return m.viewWorkoutList()
	case ViewAddRoutine:
		return m.viewAddRoutine()
	case ViewAddExercise:
		return m.viewAddExercise()
	case ViewSelectRoutineForWorkout:
		return m.viewSelectRoutineForWorkout()
	case ViewRoutineDetail:
		return m.viewRoutineDetail()
	case ViewAddExerciseToRoutine:
		return m.viewAddExerciseToRoutine()
	case ViewActiveWorkout:
		return m.viewActiveWorkout()
	case ViewEditRoutineExercise:
		return m.viewEditRoutineExercise()
	case ViewEditWorkoutSet:
		return m.viewEditWorkoutSet()
	default:
		return "Unknown view"
	}
}

func (m model) viewMainMenu() string {
	options := []string{
		"Exercises",
		"Routines",
		"Workouts",
	}

	s := "Main Menu (Use arrow keys, Enter to select, q to quit)\n\n"
	for i, option := range options {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("> ")
		}
		s += fmt.Sprintf("%s%s\n", cursor, option)
	}
	return s
}

func (m model) viewExerciseList() string {
	if len(m.exercises) == 0 {
		return "No exercises found.\n\nPress b to go back."
	}

	s := "Exercises (b to go back)\n\n"
	for i, ex := range m.exercises {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("> ")
		}

		if i == 0 {
			s += fmt.Sprintf("%s%s\n", cursor, createItemStyle.Render(ex.Name))
		} else {
			s += fmt.Sprintf("%s[%d] %s (%s)\n", cursor, ex.ID, ex.Name, ex.Muscle)
		}
	}
	return s
}

func (m model) viewRoutineList() string {
	if len(m.routines) == 0 {
		return "No routines found.\n\nPress b to go back."
	}

	s := "Routines (b to go back)\n\n"
	for i, r := range m.routines {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("> ")
		}

		if i == 0 {
			s += fmt.Sprintf("%s%s\n", cursor, createItemStyle.Render(r.Name))
		} else {
			s += fmt.Sprintf("%s[%d] %s\n", cursor, r.ID, r.Name)
		}
	}
	return s
}

func (m model) viewWorkoutList() string {
	if len(m.workouts) == 0 {
		return "No workouts found.\n\nPress b to go back."
	}

	s := "Workouts (b to go back)\n\n"
	for i, w := range m.workouts {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("> ")
		}

		if i == 0 {
			s += fmt.Sprintf("%s%s\n", cursor, createItemStyle.Render(w.StartTime))
		} else {
			endTime := w.EndTime
			if endTime == "" {
				endTime = "In Progress"
			}
			s += fmt.Sprintf("%s[%d] Routine ID: %d Started: %s Ended: %s\n",
				cursor, w.ID, w.RoutineID, w.StartTime, endTime)
		}
	}
	return s
}

func (m model) viewAddRoutine() string {
	s := "Create New Routine\n\n"
	s += "Enter routine name (ESC to cancel):\n"
	s += fmt.Sprintf("> %s", m.inputBuffer)
	if len(m.inputBuffer)%2 == 0 {
		s += "█"
	}
	return s
}

func (m model) viewAddExercise() string {
	s := "Create New Exercise\n\n"
	s += "This feature is not yet implemented.\n"
	s += "Press ESC to go back."
	return s
}

func (m model) viewSelectRoutineForWorkout() string {
	if len(m.routines) == 0 {
		return "No routines available.\n\nPress b to go back."
	}

	s := "Select Routine for Workout (b to go back)\n\n"
	for i, r := range m.routines {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("> ")
		}
		s += fmt.Sprintf("%s[%d] %s\n", cursor, r.ID, r.Name)
	}
	return s
}

func (m model) viewRoutineDetail() string {
	s := "Routine Exercises (b to go back)\n\n"

	if len(m.routineExercises) == 0 {
		s += "No exercises in this routine yet.\n\n"
	} else {
		for i, re := range m.routineExercises {
			cursor := "  "
			if m.cursor == i {
				cursor = cursorStyle.Render("> ")
			}
			s += fmt.Sprintf("%s%s - %d sets x %d reps @ %.1f lbs\n",
				cursor, re.ExerciseName, re.PlannedSets, re.PlannedReps, re.PlannedWeight)
		}
		s += "\n"
	}

	cursor := "  "
	if m.cursor == len(m.routineExercises) {
		cursor = cursorStyle.Render("> ")
	}
	s += fmt.Sprintf("%s%s\n", cursor, createItemStyle.Render("➕ Add Exercise"))

	return s
}

func (m model) viewAddExerciseToRoutine() string {
	if len(m.exercises) == 0 {
		return "No exercises available.\n\nPress b to go back."
	}

	s := "Add Exercise to Routine (b to go back)\n\n"
	for i, ex := range m.exercises {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("> ")
		}
		s += fmt.Sprintf("%s[%d] %s (%s)\n", cursor, ex.ID, ex.Name, ex.Muscle)
	}
	return s
}

func (m model) viewActiveWorkout() string {
	s := "Active Workout (b to go back)\n\n"

	for i, set := range m.workoutSets {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("> ")
		}

		checkmark := "⬜"
		if set.Completed {
			checkmark = "✅"
		}

		s += fmt.Sprintf("%s%s %s - Set %d: %.1f lbs x %d reps\n",
			cursor, checkmark, set.ExerciseName, set.SetNumber, set.Weight, set.Reps)
	}

	s += "\n"
	cursor := "  "
	if m.cursor == len(m.workoutSets) {
		cursor = cursorStyle.Render("> ")
	}
	s += fmt.Sprintf("%s%s\n", cursor, createItemStyle.Render("✅ Complete Workout"))

	cursor = "  "
	if m.cursor == len(m.workoutSets)+1 {
		cursor = cursorStyle.Render("> ")
	}
	s += fmt.Sprintf("%s%s\n", cursor, createItemStyle.Render("🗑️  Delete Workout"))

	return s
}

func (m model) viewEditRoutineExercise() string {
	if m.editingIndex >= len(m.routineExercises) {
		return "Error: Invalid exercise index"
	}

	re := m.routineExercises[m.editingIndex]
	s := fmt.Sprintf("Edit %s (ESC to cancel)\n\n", re.ExerciseName)

	fields := []string{"Sets", "Reps", "Weight"}
	values := []string{
		fmt.Sprintf("%d", re.PlannedSets),
		fmt.Sprintf("%d", re.PlannedReps),
		fmt.Sprintf("%.1f lbs", re.PlannedWeight),
	}

	if m.editingField != "" {
		s += fmt.Sprintf("Enter new %s: %s", m.editingField, m.editingValue)
		if len(m.editingValue)%2 == 0 {
			s += "█"
		}
		return s
	}

	for i, field := range fields {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("> ")
		}
		s += fmt.Sprintf("%s%s: %s\n", cursor, field, values[i])
	}

	return s
}

func (m model) viewEditWorkoutSet() string {
	if m.editingIndex >= len(m.workoutSets) {
		return "Error: Invalid set index"
	}

	ws := m.workoutSets[m.editingIndex]
	s := fmt.Sprintf("Edit %s - Set %d (ESC to cancel)\n\n", ws.ExerciseName, ws.SetNumber)

	fields := []string{"Reps", "Weight"}
	values := []string{
		fmt.Sprintf("%d", ws.Reps),
		fmt.Sprintf("%.1f lbs", ws.Weight),
	}

	if m.editingField != "" {
		s += fmt.Sprintf("Enter new %s: %s", m.editingField, m.editingValue)
		if len(m.editingValue)%2 == 0 {
			s += "█"
		}
		return s
	}

	for i, field := range fields {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("> ")
		}
		s += fmt.Sprintf("%s%s: %s\n", cursor, field, values[i])
	}

	s += "\nPress 't' to toggle completion, or select a field to edit."

	return s
}
