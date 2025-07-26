package cmd

import (
	"fmt"
	"time"

	"github.com/AndrewBennettDev/liftctl/internal/db"
	"github.com/AndrewBennettDev/liftctl/internal/models"
	"github.com/spf13/cobra"
)

var (
	startRoutineID uint
	logWorkoutID   uint
	logExerciseID  uint
	logWeight      float64
	logReps        int
	listWorkoutID  uint
	endWorkoutID   uint
)

var startWorkoutCmd = &cobra.Command{
	Use:   "start-workout",
	Short: "Start a new workout session based on a routine",
	Run: func(cmd *cobra.Command, args []string) {
		if startRoutineID == 0 {
			fmt.Println("Please provide a valid routine ID using --routine-id")
			return
		}

		database := db.GetDB()

		var routine models.Routine
		if err := database.First(&routine, startRoutineID).Error; err != nil {
			fmt.Printf("Routine with ID %d not found.\n", startRoutineID)
			return
		}

		workout := models.Workout{
			RoutineID: startRoutineID,
			StartTime: time.Now(),
		}

		if err := database.Create(&workout).Error; err != nil {
			fmt.Println("Failed to start workout:", err)
			return
		}

		fmt.Printf("Workout started for routine '%s' with Workout ID %d\n", routine.Name, workout.ID)
	},
}

var logSetCmd = &cobra.Command{
	Use:   "log-set",
	Short: "Log a set for an exercise during a workout",
	Run: func(cmd *cobra.Command, args []string) {
		if logWorkoutID == 0 || logExerciseID == 0 || logReps <= 0 {
			fmt.Println("You must provide --workout-id, --exercise-id, and --reps (must be > 0)")
			return
		}

		db := db.GetDB()

		set := models.WorkoutSet{
			WorkoutID:  logWorkoutID,
			ExerciseID: logExerciseID,
			Weight:     logWeight,
			Reps:       logReps,
			Timestamp:  time.Now(),
		}

		if err := db.Create(&set).Error; err != nil {
			fmt.Println("Failed to log set:", err)
			return
		}

		fmt.Printf("Logged %d reps @ %.2fkg for exercise ID %d in workout ID %d\n",
			set.Reps, set.Weight, set.ExerciseID, set.WorkoutID)
	},
}

var listSetsCmd = &cobra.Command{
	Use:   "list-sets",
	Short: "List all sets in a workout",
	Run: func(cmd *cobra.Command, args []string) {
		if listWorkoutID == 0 {
			fmt.Println("You must provide --workout-id")
			return
		}

		db := db.GetDB()

		var sets []models.WorkoutSet
		err := db.Preload("Exercise").Where("workout_id = ?", listWorkoutID).Order("timestamp asc").Find(&sets).Error
		if err != nil {
			fmt.Println("Failed to retrieve sets:", err)
			return
		}

		if len(sets) == 0 {
			fmt.Printf("No sets found for workout ID %d.\n", listWorkoutID)
			return
		}

		fmt.Printf("Sets for workout ID %d:\n", listWorkoutID)
		for _, set := range sets {
			fmt.Printf("[%s] %s: %d reps @ %.2fkg\n",
				set.Timestamp.Format("2006-01-02 15:04:05"),
				set.Exercise.Name,
				set.Reps,
				set.Weight,
			)
		}
	},
}

var endWorkoutCmd = &cobra.Command{
	Use:   "end-workout",
	Short: "Mark a workout as finished",
	Run: func(cmd *cobra.Command, args []string) {
		if endWorkoutID == 0 {
			fmt.Println("Please provide --workout-id")
			return
		}

		db := db.GetDB()

		var workout models.Workout
		if err := db.First(&workout, endWorkoutID).Error; err != nil {
			fmt.Printf("Workout with ID %d not found.\n", endWorkoutID)
			return
		}

		if workout.FinishedAt != nil {
			fmt.Printf("Workout ID %d was already completed at %s\n", endWorkoutID, workout.FinishedAt.Format("2006-01-02 15:04:05"))
			return
		}

		now := time.Now()
		workout.FinishedAt = &now

		if err := db.Save(&workout).Error; err != nil {
			fmt.Println("Failed to mark workout as completed:", err)
			return
		}

		fmt.Printf("Workout ID %d completed at %s\n", workout.ID, now.Format("2006-01-02 15:04:05"))
	},
}

func init() {
	rootCmd.AddCommand(startWorkoutCmd)
	rootCmd.AddCommand(logSetCmd)
	rootCmd.AddCommand(listSetsCmd)
	rootCmd.AddCommand(endWorkoutCmd)

	startWorkoutCmd.Flags().UintVarP(&startRoutineID, "routine-id", "i", 0, "ID of the routine to start")
	startWorkoutCmd.MarkFlagRequired("routine-id")

	logSetCmd.Flags().UintVarP(&logWorkoutID, "workout-id", "i", 0, "Workout ID")
	logSetCmd.Flags().UintVarP(&logExerciseID, "exercise-id", "x", 0, "Exercise ID")
	logSetCmd.Flags().Float64VarP(&logWeight, "weight", "w", 0, "Weight used")
	logSetCmd.Flags().IntVarP(&logReps, "reps", "r", 0, "Number of reps")

	listSetsCmd.Flags().UintVarP(&listWorkoutID, "workout-id", "i", 0, "Workout ID to list sets for")

	endWorkoutCmd.Flags().UintVarP(&endWorkoutID, "workout-id", "i", 0, "Workout ID to complete")
}
