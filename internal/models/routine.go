package models

import (
	"gorm.io/gorm"
)

type Routine struct {
	gorm.Model
	Name             string
	RoutineExercises []RoutineExercise
}

type RoutineExercise struct {
	ID            uint `gorm:"primaryKey"`
	RoutineID     uint
	ExerciseID    uint
	Order         int
	PlannedSets   int
	PlannedReps   int
	PlannedWeight float64

	Routine  Routine  `gorm:"foreignKey:RoutineID"`
	Exercise Exercise `gorm:"foreignKey:ExerciseID"`
}
