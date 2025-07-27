package tui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/AndrewBennettDev/liftctl/internal/db"
	"github.com/AndrewBennettDev/liftctl/internal/models"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.currentView == ViewAddRoutine || m.currentView == ViewAddExercise ||
			(m.currentView == ViewEditRoutineExercise && m.editingField != "") ||
			(m.currentView == ViewEditWorkoutSet && m.editingField != "") {
			return m.handleInput(msg)
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			if m.currentView == ViewEditRoutineExercise {
				m.currentView = ViewRoutineDetail
				m.cursor = 0
				return m, nil
			} else if m.currentView == ViewEditWorkoutSet {
				m.currentView = ViewActiveWorkout
				m.cursor = 0
				return m, nil
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < m.currentListLength()-1 {
				m.cursor++
			}
		case "enter":
			return m.handleEnter()
		case "b":
			m.currentView = ViewMainMenu
			m.cursor = 0
		case "t":
			if m.currentView == ViewActiveWorkout && m.cursor < len(m.workoutSets) {
				m.toggleSetCompletion(m.cursor)
			}
		}
	}
	return m, nil
}

func (m model) handleInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		return m.handleInputSubmit()
	case "esc", "ctrl+c":
		if m.currentView == ViewAddRoutine {
			m.inputBuffer = ""
			m.currentView = ViewRoutines
		} else if m.currentView == ViewAddExercise {
			m.inputBuffer = ""
			m.exerciseName = ""
			m.exerciseMuscle = ""
			m.exerciseStep = 0
			m.currentView = ViewExercises
		} else if m.currentView == ViewEditRoutineExercise {
			if m.editingField != "" {
				m.editingField = ""
				m.editingValue = ""
				return m, nil
			} else {
				m.currentView = ViewRoutineDetail
			}
		} else if m.currentView == ViewEditWorkoutSet {
			if m.editingField != "" {
				m.editingField = ""
				m.editingValue = ""
				return m, nil
			} else {
				m.currentView = ViewActiveWorkout
			}
		}
		m.cursor = 0
		return m, nil
	case "backspace":
		if m.currentView == ViewAddRoutine || m.currentView == ViewAddExercise {
			if len(m.inputBuffer) > 0 {
				m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
			}
		} else if m.currentView == ViewEditRoutineExercise || m.currentView == ViewEditWorkoutSet {
			if len(m.editingValue) > 0 {
				m.editingValue = m.editingValue[:len(m.editingValue)-1]
			}
		}
	default:
		if len(msg.String()) == 1 {
			if m.currentView == ViewAddRoutine || m.currentView == ViewAddExercise {
				m.inputBuffer += msg.String()
			} else if m.currentView == ViewEditRoutineExercise || m.currentView == ViewEditWorkoutSet {
				if (msg.String() >= "0" && msg.String() <= "9") || msg.String() == "." {
					m.editingValue += msg.String()
				}
			}
		}
	}
	return m, nil
}

func (m model) handleInputSubmit() (tea.Model, tea.Cmd) {
	switch m.currentView {
	case ViewAddRoutine:
		if strings.TrimSpace(m.inputBuffer) != "" {
			m.createRoutineFromInput()
			m.loadRoutines()
			m.currentView = ViewRoutines
		}
		m.inputBuffer = ""
	case ViewAddExercise:
		if strings.TrimSpace(m.inputBuffer) != "" {
			if m.exerciseStep == 0 {
				m.exerciseName = strings.TrimSpace(m.inputBuffer)
				m.exerciseStep = 1
				m.inputBuffer = ""
				return m, nil
			} else if m.exerciseStep == 1 {
				m.exerciseMuscle = strings.TrimSpace(m.inputBuffer)
				m.createExerciseFromInput()
				m.loadExercises()
				m.exerciseName = ""
				m.exerciseMuscle = ""
				m.exerciseStep = 0
				m.currentView = ViewExercises
			}
		}
		m.inputBuffer = ""
	case ViewEditRoutineExercise:
		if strings.TrimSpace(m.editingValue) != "" {
			m.updateRoutineExerciseField()
			m.loadRoutineExercises()
		}
		m.editingField = ""
		m.editingValue = ""
		return m, nil
	case ViewEditWorkoutSet:
		if strings.TrimSpace(m.editingValue) != "" {
			m.updateWorkoutSetField()
			m.loadWorkoutSets()
		}
		m.editingField = ""
		m.editingValue = ""
		return m, nil
	}

	m.cursor = 0
	return m, nil
}

func (m *model) createRoutineFromInput() {
	rt := models.Routine{
		Name: strings.TrimSpace(m.inputBuffer),
	}

	db := db.GetDB()
	if err := db.Create(&rt).Error; err != nil {
	}
}

func (m *model) createExerciseFromInput() {
	ex := models.Exercise{
		Name:   m.exerciseName,
		Muscle: m.exerciseMuscle,
	}

	db := db.GetDB()
	if err := db.Create(&ex).Error; err != nil {
	}
}

func (m model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.currentView {
	case ViewMainMenu:
		switch m.cursor {
		case 0:
			m.loadExercises()
			m.currentView = ViewExercises
		case 1:
			m.loadRoutines()
			m.currentView = ViewRoutines
		case 2:
			m.loadWorkouts()
			m.currentView = ViewWorkouts
		}
	case ViewExercises:
		if m.cursor == 0 {
			m.currentView = ViewAddExercise
			m.inputBuffer = ""
			m.exerciseName = ""
			m.exerciseMuscle = ""
			m.exerciseStep = 0
		}
	case ViewRoutines:
		if m.cursor == 0 {
			m.currentView = ViewAddRoutine
			m.inputBuffer = ""
		} else if m.cursor > 0 && m.cursor < len(m.routines) {
			m.selectedRoutineID = m.routines[m.cursor].ID
			m.loadRoutineExercises()
			m.currentView = ViewRoutineDetail
		}
	case ViewWorkouts:
		if m.cursor == 0 {
			m.loadRoutines()
			m.currentView = ViewSelectRoutineForWorkout
		} else if m.cursor > 0 && m.cursor < len(m.workouts) {
			workout := m.workouts[m.cursor]
			if workout.EndTime == "" || workout.EndTime == "In Progress" {
				m.activeWorkoutID = workout.ID
				m.loadWorkoutSets()
				m.currentView = ViewActiveWorkout
			}
		}
	case ViewSelectRoutineForWorkout:
		if m.cursor < len(m.routines) {
			if m.startWorkoutWithRoutine(m.routines[m.cursor].ID) {
				return m, nil
			}
			m.loadWorkouts()
			m.currentView = ViewWorkouts
		}
	case ViewRoutineDetail:
		if m.cursor == len(m.routineExercises) {
			m.loadExercisesForSelection()
			m.currentView = ViewAddExerciseToRoutine
		} else if m.cursor < len(m.routineExercises) {
			m.editingIndex = m.cursor
			m.currentView = ViewEditRoutineExercise
		}
	case ViewAddExerciseToRoutine:
		if m.cursor < len(m.exercises) {
			m.addExerciseToRoutine(m.exercises[m.cursor].ID)
			m.loadRoutineExercises()
			m.currentView = ViewRoutineDetail
		}
	case ViewActiveWorkout:
		if m.cursor < len(m.workoutSets) {
			m.editingIndex = m.cursor
			m.currentView = ViewEditWorkoutSet
		} else if m.cursor == len(m.workoutSets) {
			m.completeWorkout()
			m.loadWorkouts()
			m.currentView = ViewWorkouts
		} else if m.cursor == len(m.workoutSets)+1 {
			m.deleteWorkout()
			m.loadWorkouts()
			m.currentView = ViewWorkouts
		}
	case ViewEditRoutineExercise:
		m.editingField = []string{"sets", "reps", "weight"}[m.cursor]
		re := m.routineExercises[m.editingIndex]
		switch m.editingField {
		case "sets":
			m.editingValue = fmt.Sprintf("%d", re.PlannedSets)
		case "reps":
			m.editingValue = fmt.Sprintf("%d", re.PlannedReps)
		case "weight":
			m.editingValue = fmt.Sprintf("%.1f", re.PlannedWeight)
		}
		return m, nil
	case ViewEditWorkoutSet:
		m.editingField = []string{"reps", "weight"}[m.cursor]
		ws := m.workoutSets[m.editingIndex]
		switch m.editingField {
		case "reps":
			m.editingValue = fmt.Sprintf("%d", ws.Reps)
		case "weight":
			m.editingValue = fmt.Sprintf("%.1f", ws.Weight)
		}
		return m, nil
	}

	m.cursor = 0
	return m, nil
}

func (m *model) startWorkoutWithRoutine(routineID uint) bool {
	workout := models.Workout{
		RoutineID: routineID,
		StartTime: time.Now(),
	}

	db := db.GetDB()
	if err := db.Create(&workout).Error; err != nil {
		return false
	}

	var routineExercises []models.RoutineExercise
	if err := db.Preload("Exercise").Where("routine_id = ?", routineID).Order("`order`").Find(&routineExercises).Error; err != nil {
		return false
	}

	for _, re := range routineExercises {
		for setNum := 1; setNum <= re.PlannedSets; setNum++ {
			workoutSet := models.WorkoutSet{
				WorkoutID:  workout.ID,
				ExerciseID: re.ExerciseID,
				SetNumber:  setNum,
				Weight:     re.PlannedWeight,
				Reps:       re.PlannedReps,
				Completed:  false,
				Timestamp:  time.Now(),
			}
			db.Create(&workoutSet)
		}
	}

	m.activeWorkoutID = workout.ID
	m.loadWorkoutSets()
	m.currentView = ViewActiveWorkout
	return true
}

func (m *model) loadWorkoutSets() {
	db := db.GetDB()
	var workoutSets []models.WorkoutSet
	if err := db.Preload("Exercise").Where("workout_id = ?", m.activeWorkoutID).Order("exercise_id, set_number").Find(&workoutSets).Error; err != nil {
		fmt.Println("Error loading workout sets:", err)
		m.workoutSets = []WorkoutSet{}
		return
	}
	m.workoutSets = convertWorkoutSets(workoutSets)
}

func (m *model) toggleSetCompletion(setIndex int) {
	if setIndex >= len(m.workoutSets) {
		return
	}

	set := &m.workoutSets[setIndex]
	set.Completed = !set.Completed

	db := db.GetDB()
	db.Model(&models.WorkoutSet{}).Where("id = ?", set.ID).Update("completed", set.Completed)
}

func (m *model) completeWorkout() {
	now := time.Now()
	db := db.GetDB()
	db.Model(&models.Workout{}).Where("id = ?", m.activeWorkoutID).Updates(map[string]interface{}{
		"end_time":    now,
		"finished_at": now,
	})
}

func (m *model) deleteWorkout() {
	db := db.GetDB()
	db.Where("workout_id = ?", m.activeWorkoutID).Delete(&models.WorkoutSet{})
	db.Delete(&models.Workout{}, m.activeWorkoutID)
}

func (m *model) updateRoutineExerciseField() {
	if m.editingIndex >= len(m.routineExercises) {
		return
	}

	re := &m.routineExercises[m.editingIndex]
	db := db.GetDB()

	switch m.editingField {
	case "sets":
		if val, err := strconv.Atoi(m.editingValue); err == nil && val > 0 {
			re.PlannedSets = val
			db.Model(&models.RoutineExercise{}).Where("id = ?", re.ID).Update("planned_sets", val)
		}
	case "reps":
		if val, err := strconv.Atoi(m.editingValue); err == nil && val > 0 {
			re.PlannedReps = val
			db.Model(&models.RoutineExercise{}).Where("id = ?", re.ID).Update("planned_reps", val)
		}
	case "weight":
		if val, err := strconv.ParseFloat(m.editingValue, 64); err == nil && val >= 0 {
			re.PlannedWeight = val
			db.Model(&models.RoutineExercise{}).Where("id = ?", re.ID).Update("planned_weight", val)
		}
	}
}

func (m *model) updateWorkoutSetField() {
	if m.editingIndex >= len(m.workoutSets) {
		return
	}

	ws := &m.workoutSets[m.editingIndex]
	db := db.GetDB()

	switch m.editingField {
	case "reps":
		if val, err := strconv.Atoi(m.editingValue); err == nil && val > 0 {
			ws.Reps = val
			db.Model(&models.WorkoutSet{}).Where("id = ?", ws.ID).Update("reps", val)
		}
	case "weight":
		if val, err := strconv.ParseFloat(m.editingValue, 64); err == nil && val >= 0 {
			ws.Weight = val
			db.Model(&models.WorkoutSet{}).Where("id = ?", ws.ID).Update("weight", val)
		}
	}
}

// -- Loaders --

func (m *model) loadExercises() {
	db := db.GetDB()
	var exs []models.Exercise
	if err := db.Find(&exs).Error; err != nil {
		fmt.Println("Error loading exercises:", err)
		m.exercises = []Exercise{}
		return
	}
	m.exercises = append([]Exercise{{Name: "➕ Create New Exercise"}}, convertExercises(exs)...)
}

func (m *model) loadExercisesForSelection() {
	db := db.GetDB()
	var exs []models.Exercise
	if err := db.Find(&exs).Error; err != nil {
		fmt.Println("Error loading exercises:", err)
		m.exercises = []Exercise{}
		return
	}
	m.exercises = convertExercises(exs)
}

func (m *model) loadRoutines() {
	db := db.GetDB()
	var rts []models.Routine
	if err := db.Find(&rts).Error; err != nil {
		fmt.Println("Error loading routines:", err)
		m.routines = []Routine{}
		return
	}
	m.routines = append([]Routine{{Name: "➕ Create New Routine"}}, convertRoutines(rts)...)
}

func (m *model) loadWorkouts() {
	db := db.GetDB()
	var wks []models.Workout
	if err := db.Find(&wks).Error; err != nil {
		fmt.Println("Error loading workouts:", err)
		m.workouts = []Workout{}
		return
	}
	m.workouts = append([]Workout{{ID: 0, RoutineID: 0, StartTime: "➕ Start New Workout"}}, convertWorkouts(wks)...)
}

func (m *model) loadRoutineExercises() {
	db := db.GetDB()
	var routineExercises []models.RoutineExercise
	if err := db.Preload("Exercise").Where("routine_id = ?", m.selectedRoutineID).Order("`order`").Find(&routineExercises).Error; err != nil {
		fmt.Println("Error loading routine exercises:", err)
		m.routineExercises = []RoutineExercise{}
		return
	}
	m.routineExercises = convertRoutineExercises(routineExercises)
}

func (m *model) addExerciseToRoutine(exerciseID uint) {
	nextOrder := len(m.routineExercises) + 1

	routineExercise := models.RoutineExercise{
		RoutineID:     m.selectedRoutineID,
		ExerciseID:    exerciseID,
		Order:         nextOrder,
		PlannedSets:   3,
		PlannedReps:   10,
		PlannedWeight: 0,
	}

	db := db.GetDB()
	if err := db.Create(&routineExercise).Error; err != nil {
		fmt.Println("Error adding exercise to routine:", err)
	}
}

func convertExercises(input []models.Exercise) []Exercise {
	out := make([]Exercise, len(input))
	for i, e := range input {
		out[i] = Exercise{ID: e.ID, Name: e.Name, Muscle: e.Muscle}
	}
	return out
}

func convertRoutines(input []models.Routine) []Routine {
	out := make([]Routine, len(input))
	for i, r := range input {
		out[i] = Routine{ID: r.ID, Name: r.Name}
	}
	return out
}

func convertWorkouts(input []models.Workout) []Workout {
	out := make([]Workout, len(input))
	for i, w := range input {
		endTime := ""
		if w.FinishedAt != nil {
			endTime = w.FinishedAt.Format("2006-01-02 15:04:05")
		}
		out[i] = Workout{
			ID:        w.ID,
			RoutineID: w.RoutineID,
			StartTime: w.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:   endTime,
		}
	}
	return out
}

func convertRoutineExercises(input []models.RoutineExercise) []RoutineExercise {
	out := make([]RoutineExercise, len(input))
	for i, re := range input {
		out[i] = RoutineExercise{
			ID:            re.ID,
			RoutineID:     re.RoutineID,
			ExerciseID:    re.ExerciseID,
			ExerciseName:  re.Exercise.Name,
			Order:         re.Order,
			PlannedSets:   re.PlannedSets,
			PlannedReps:   re.PlannedReps,
			PlannedWeight: re.PlannedWeight,
		}
	}
	return out
}

func convertWorkoutSets(input []models.WorkoutSet) []WorkoutSet {
	out := make([]WorkoutSet, len(input))
	for i, ws := range input {
		out[i] = WorkoutSet{
			ID:           ws.ID,
			WorkoutID:    ws.WorkoutID,
			ExerciseID:   ws.ExerciseID,
			ExerciseName: ws.Exercise.Name,
			SetNumber:    ws.SetNumber,
			Weight:       ws.Weight,
			Reps:         ws.Reps,
			Completed:    ws.Completed,
		}
	}
	return out
}
