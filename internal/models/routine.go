package models

import (
	"gorm.io/gorm"
)

type Routine struct {
	gorm.Model
	Name      string
	Exercises []Exercise `gorm:"many2many:routine_exercises;"`
}
