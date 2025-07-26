package tui

import tea "github.com/charmbracelet/bubbletea"

type View int

const (
	ViewMainMenu View = iota
	ViewExercises
	ViewRoutines
	ViewWorkouts
	ViewAddExercise
	ViewAddRoutine
	ViewSelectRoutineForWorkout
	ViewRoutineDetail
	ViewAddExerciseToRoutine
	ViewActiveWorkout
	ViewWorkoutComplete
	ViewEditRoutineExercise
	ViewEditWorkoutSet
)

type model struct {
	currentView View

	exercises []Exercise
	routines  []Routine
	workouts  []Workout

	cursor int

	inputBuffer string

	selectedRoutineID uint
	routineExercises  []RoutineExercise

	activeWorkoutID uint
	workoutSets     []WorkoutSet

	editingIndex int
	editingField string
	editingValue string
}

type Exercise struct {
	ID     uint
	Name   string
	Muscle string
}

type Routine struct {
	ID   uint
	Name string
}

type Workout struct {
	ID        uint
	RoutineID uint
	StartTime string
	EndTime   string
}

type RoutineExercise struct {
	ID            uint
	RoutineID     uint
	ExerciseID    uint
	ExerciseName  string
	Order         int
	PlannedSets   int
	PlannedReps   int
	PlannedWeight float64
}

type WorkoutSet struct {
	ID           uint
	WorkoutID    uint
	ExerciseID   uint
	ExerciseName string
	SetNumber    int
	Weight       float64
	Reps         int
	Completed    bool
}

func InitialModel() model {
	return model{
		currentView: ViewMainMenu,
		cursor:      0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) currentListLength() int {
	switch m.currentView {
	case ViewMainMenu:
		return 3
	case ViewExercises:
		return len(m.exercises)
	case ViewRoutines:
		return len(m.routines)
	case ViewWorkouts:
		return len(m.workouts)
	case ViewSelectRoutineForWorkout:
		return len(m.routines)
	case ViewRoutineDetail:
		return len(m.routineExercises) + 1
	case ViewAddExerciseToRoutine:
		return len(m.exercises)
	case ViewActiveWorkout:
		return len(m.workoutSets) + 2
	case ViewEditRoutineExercise:
		return 3
	case ViewEditWorkoutSet:
		return 2
	}
	return 0
}
