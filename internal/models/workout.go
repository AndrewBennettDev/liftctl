package models

import "time"

type Workout struct {
	ID         uint `gorm:"primaryKey"`
	RoutineID  uint
	StartTime  time.Time
	EndTime    *time.Time
	FinishedAt *time.Time `gorm:"default:null"`

	Routine Routine
	Sets    []WorkoutSet
}

type WorkoutSet struct {
	ID         uint `gorm:"primaryKey"`
	WorkoutID  uint
	ExerciseID uint
	SetNumber  int
	Weight     float64
	Reps       int
	Completed  bool `gorm:"default:false"`
	Timestamp  time.Time

	Workout  Workout  `gorm:"foreignKey:WorkoutID"`
	Exercise Exercise `gorm:"foreignKey:ExerciseID"`
}
