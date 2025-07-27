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

	err = DB.AutoMigrate(
		&models.Exercise{},
		&models.Routine{},
		&models.RoutineExercise{},
		&models.Workout{},
		&models.WorkoutSet{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	seedDefaultExercises()

	return DB
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call Init() first.")
	}
	return DB
}

func seedDefaultExercises() {
	var count int64
	DB.Model(&models.Exercise{}).Count(&count)

	if count > 0 {
		return
	}

	defaultExercises := []models.Exercise{
		{Name: "Bench Press", Muscle: "chest"},
		{Name: "Incline Bench Press", Muscle: "chest"},
		{Name: "Dumbbell Flyes", Muscle: "chest"},
		{Name: "Push-ups", Muscle: "chest"},
		{Name: "Dips", Muscle: "chest"},

		{Name: "Squats", Muscle: "legs"},
		{Name: "Deadlifts", Muscle: "legs"},
		{Name: "Lunges", Muscle: "legs"},
		{Name: "Leg Press", Muscle: "legs"},
		{Name: "Leg Curls", Muscle: "legs"},
		{Name: "Leg Extensions", Muscle: "legs"},
		{Name: "Calf Raises", Muscle: "legs"},

		{Name: "Pull-ups", Muscle: "back"},
		{Name: "Chin-ups", Muscle: "back"},
		{Name: "Bent-over Rows", Muscle: "back"},
		{Name: "Lat Pulldowns", Muscle: "back"},
		{Name: "Seated Cable Rows", Muscle: "back"},
		{Name: "T-Bar Rows", Muscle: "back"},

		{Name: "Overhead Press", Muscle: "shoulders"},
		{Name: "Lateral Raises", Muscle: "shoulders"},
		{Name: "Front Raises", Muscle: "shoulders"},
		{Name: "Rear Delt Flyes", Muscle: "shoulders"},
		{Name: "Arnold Press", Muscle: "shoulders"},
		{Name: "Upright Rows", Muscle: "shoulders"},

		{Name: "Bicep Curls", Muscle: "arms"},
		{Name: "Hammer Curls", Muscle: "arms"},
		{Name: "Preacher Curls", Muscle: "arms"},
		{Name: "Tricep Dips", Muscle: "arms"},
		{Name: "Tricep Extensions", Muscle: "arms"},
		{Name: "Close-grip Bench Press", Muscle: "arms"},

		{Name: "Planks", Muscle: "core"},
		{Name: "Crunches", Muscle: "core"},
		{Name: "Russian Twists", Muscle: "core"},
		{Name: "Mountain Climbers", Muscle: "core"},
		{Name: "Leg Raises", Muscle: "core"},
		{Name: "Dead Bug", Muscle: "core"},
	}

	result := DB.Create(&defaultExercises)
	if result.Error != nil {
		fmt.Printf("Warning: Could not seed default exercises: %v\n", result.Error)
		return
	}
}
