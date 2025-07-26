package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/AndrewBennettDev/liftctl/internal/models"
)

var DB *gorm.DB

func Init() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("workout.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	fmt.Println("Connected to database")

	err = DB.AutoMigrate(
		&models.Exercise{},
		&models.Routine{},
		&models.RoutineExercise{},
		&models.Workout{},
		&models.WorkoutSet{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	fmt.Println("Database migrated")

	return DB
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call Init() first.")
	}
	return DB
}
