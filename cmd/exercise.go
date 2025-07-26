package cmd

import (
	"fmt"

	"github.com/AndrewBennettDev/liftctl/internal/db"
	"github.com/AndrewBennettDev/liftctl/internal/models"
	"github.com/spf13/cobra"
)

var (
	exerciseName   string
	exerciseMuscle string
)

var addExerciseCmd = &cobra.Command{
	Use:   "add-exercise",
	Short: "Add a new exercise",
	Run: func(cmd *cobra.Command, args []string) {
		if exerciseName == "" {
			fmt.Println("Exercise name is required (--name)")
			return
		}

		exercise := models.Exercise{
			Name:   exerciseName,
			Muscle: exerciseMuscle,
		}

		result := db.DB.Create(&exercise)
		if result.Error != nil {
			fmt.Println("Failed to create exercise:", result.Error)
			return
		}

		fmt.Printf("Exercise '%s' added successfully!\n", exercise.Name)
	},
}

var listExercisesCmd = &cobra.Command{
	Use:   "list-exercises",
	Short: "List all exercises",
	Run: func(cmd *cobra.Command, args []string) {
		var exercises []models.Exercise
		result := db.DB.Find(&exercises)
		if result.Error != nil {
			fmt.Println("Error fetching exercises:", result.Error)
			return
		}

		if len(exercises) == 0 {
			fmt.Println("No exercises found.")
			return
		}

		fmt.Println("Exercises:")
		for _, ex := range exercises {
			fmt.Printf("[%d] %s (%s)\n", ex.ID, ex.Name, ex.Muscle)
		}
	},
}

var editExerciseCmd = &cobra.Command{
	Use:   "edit-exercise",
	Short: "Edit an existing exercise",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Edit exercise functionality is not implemented yet.")
	},
}

var deleteExerciseCmd = &cobra.Command{
	Use:   "delete-exercise",
	Short: "Delete an existing exercise",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete exercise functionality is not implemented yet.")
	},
}

func init() {
	rootCmd.AddCommand(addExerciseCmd)
	rootCmd.AddCommand(listExercisesCmd)
	rootCmd.AddCommand(editExerciseCmd)
	rootCmd.AddCommand(deleteExerciseCmd)

	addExerciseCmd.Flags().StringVar(&exerciseName, "name", "", "Name of the exercise")
	addExerciseCmd.Flags().StringVar(&exerciseMuscle, "muscle", "", "Target muscle group (optional)")
}
