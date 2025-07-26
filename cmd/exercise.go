package cmd

import (
	"fmt"

	"github.com/AndrewBennettDev/liftctl/internal/db"
	"github.com/AndrewBennettDev/liftctl/internal/models"
	"github.com/spf13/cobra"
)

var (
	exerciseName       string
	exerciseMuscle     string
	editExerciseID     uint
	editExerciseName   string
	editExerciseMuscle string
	deleteExerciseID   uint
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
		if editExerciseID == 0 {
			fmt.Println("--id is required and must be a positive integer")
			return
		}

		if editExerciseName == "" && editExerciseMuscle == "" {
			fmt.Println("At least one of --name or --muscle must be provided to update")
			return
		}

		db := db.GetDB()
		var exercise models.Exercise

		if err := db.First(&exercise, editExerciseID).Error; err != nil {
			fmt.Printf("Exercise with ID %d not found.\n", editExerciseID)
			return
		}

		if editExerciseName != "" {
			exercise.Name = editExerciseName
		}
		if editExerciseMuscle != "" {
			exercise.Muscle = editExerciseMuscle
		}

		if err := db.Save(&exercise).Error; err != nil {
			fmt.Println("Failed to update exercise:", err)
			return
		}

		fmt.Printf("Exercise ID %d updated successfully!\n", exercise.ID)

	},
}

var deleteExerciseCmd = &cobra.Command{
	Use:   "delete-exercise",
	Short: "Delete an existing exercise",
	Run: func(cmd *cobra.Command, args []string) {
		if deleteExerciseID == 0 {
			fmt.Println("--id is required and must be a valid positive integer")
			return
		}

		db := db.GetDB()
		var exercise models.Exercise
		if err := db.First(&exercise, deleteExerciseID).Error; err != nil {
			fmt.Printf("Exercise with ID %d not found.\n", deleteExerciseID)
			return
		}

		if err := db.Delete(&exercise).Error; err != nil {
			fmt.Println("Failed to delete exercise:", err)
			return
		}

		fmt.Printf("Exercise with ID %d deleted successfully.\n", deleteExerciseID)
	},
}

func init() {
	rootCmd.AddCommand(addExerciseCmd)
	rootCmd.AddCommand(listExercisesCmd)
	rootCmd.AddCommand(editExerciseCmd)
	rootCmd.AddCommand(deleteExerciseCmd)

	addExerciseCmd.Flags().StringVarP(&exerciseName, "name", "n", "", "Name of the exercise")
	addExerciseCmd.Flags().StringVarP(&exerciseMuscle, "muscle", "m", "", "Target muscle group (optional)")

	editExerciseCmd.Flags().UintVarP(&editExerciseID, "id", "i", 0, "ID of the exercise to edit")
	editExerciseCmd.Flags().StringVarP(&editExerciseName, "name", "n", "", "New name for the exercise")
	editExerciseCmd.Flags().StringVarP(&editExerciseMuscle, "muscle", "m", "", "New muscle group for the exercise")
	editExerciseCmd.MarkFlagRequired("id")

	deleteExerciseCmd.Flags().UintVarP(&deleteExerciseID, "id", "i", 0, "ID of the exercise to delete")
	deleteExerciseCmd.MarkFlagRequired("id")
}
