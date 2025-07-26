package cmd

import (
	"fmt"

	"github.com/AndrewBennettDev/liftctl/internal/db"
	"github.com/AndrewBennettDev/liftctl/internal/models"
	"github.com/spf13/cobra"
)

var (
	routineName      string
	routineID        uint
	newName          string
	exerciseIDs      []uint
	routineIDForList uint
	removeRoutineID  uint
	removeExerciseID uint
)

var createRoutineCmd = &cobra.Command{
	Use:   "create-routine",
	Short: "Create new workout routine",
	Run: func(cmd *cobra.Command, args []string) {
		if routineName == "" {
			fmt.Println("Routine must include name (--name)")
			return
		}

		routine := models.Routine{
			Name: routineName,
		}

		result := db.DB.Create(&routine)
		if result.Error != nil {
			fmt.Println("Failed to create routine: ", result.Error)
			return
		}

		fmt.Printf("Routine %s create with ID %d\n", routine.Name, routine.ID)
	},
}

var listRoutinesCmd = &cobra.Command{
	Use:   "list-routines",
	Short: "List all routines",
	Run: func(cmd *cobra.Command, args []string) {
		var routines []models.Routine
		result := db.DB.Find(&routines)
		if result.Error != nil {
			fmt.Println("Error fetching routines:, ", result.Error)
			return
		}
		for _, routine := range routines {
			fmt.Printf("[%d] %s\n", routine.ID, routine.Name)
		}
	},
}

var editRoutineCmd = &cobra.Command{
	Use:   "edit-routine",
	Short: "Edit an existing routine",
	Run: func(cmd *cobra.Command, args []string) {
		if routineID == 0 {
			fmt.Println("--id is required")
			return
		}
		if newName == "" {
			fmt.Println("--name is required to update the routine")
			return
		}

		db := db.GetDB()
		var routine models.Routine
		if err := db.First(&routine, routineID).Error; err != nil {
			fmt.Printf("Routine with ID %d not found.\n", routineID)
			return
		}

		routine.Name = newName
		if err := db.Save(&routine).Error; err != nil {
			fmt.Println("Failed to update routine:", err)
			return
		}

		fmt.Printf("Routine ID %d updated to '%s'.\n", routineID, routine.Name)
	},
}

var deleteRoutineCmd = &cobra.Command{
	Use:   "delete-routine",
	Short: "Delete an existing routine",
	Run: func(cmd *cobra.Command, args []string) {
		if routineID == 0 {
			fmt.Println("--id is required")
			return
		}

		db := db.GetDB()
		var routine models.Routine
		if err := db.First(&routine, routineID).Error; err != nil {
			fmt.Printf("Routine with ID %d not found.\n", routineID)
			return
		}

		if err := db.Delete(&routine).Error; err != nil {
			fmt.Println("Failed to delete routine:", err)
			return
		}

		fmt.Printf("Routine with ID %d deleted successfully.\n", routineID)
	},
}

var addExerciseToRoutineCmd = &cobra.Command{
	Use:   "add-exercise-to-routine",
	Short: "Add an exercise to a routine",
	Run: func(cmd *cobra.Command, args []string) {
		if routineID == 0 {
			fmt.Println("--routine-id is required")
			return
		}
		if len(exerciseIDs) == 0 {
			fmt.Println("--exercise-ids is required (comma-separated list of exercise IDs)")
			return
		}

		db := db.GetDB()

		var routine models.Routine
		if err := db.Preload("Exercises").First(&routine, routineID).Error; err != nil {
			fmt.Printf("Routine with ID %d not found.\n", routineID)
			return
		}

		var exercises []models.Exercise
		if err := db.Find(&exercises, exerciseIDs).Error; err != nil {
			fmt.Println("Error retrieving exercises:", err)
			return
		}
		if len(exercises) == 0 {
			fmt.Println("No valid exercises found with the given IDs.")
			return
		}

		if err := db.Model(&routine).Association("Exercises").Append(exercises); err != nil {
			fmt.Println("Failed to add exercises to routine:", err)
			return
		}

		fmt.Printf("Added %d exercise(s) to routine '%s'.\n", len(exercises), routine.Name)
	},
}

var listRoutineExercisesCmd = &cobra.Command{
	Use:   "list-routine-exercises",
	Short: "List all exercises in a specific routine",
	Run: func(cmd *cobra.Command, args []string) {
		if routineIDForList == 0 {
			fmt.Println("--routine-id is required")
			return
		}

		db := db.GetDB()
		var routine models.Routine

		if err := db.Preload("Exercises").First(&routine, routineIDForList).Error; err != nil {
			fmt.Printf("Routine with ID %d not found.\n", routineIDForList)
			return
		}

		var routineExercises []models.RoutineExercise
		if err := db.Preload("Exercise").Where("routine_id = ?", routine.ID).Order("`order`").Find(&routineExercises).Error; err != nil {
			fmt.Printf("Error loading exercises for routine: %v\n", err)
			return
		}

		if len(routineExercises) == 0 {
			fmt.Printf("Routine '%s' has no exercises.\n", routine.Name)
			return
		}

		fmt.Printf("Exercises in routine '%s':\n", routine.Name)
		for _, re := range routineExercises {
			fmt.Printf("  [%d] %s (%s) - %d sets x %d reps @ %.1f lbs\n",
				re.Exercise.ID, re.Exercise.Name, re.Exercise.Muscle,
				re.PlannedSets, re.PlannedReps, re.PlannedWeight)
		}
	},
}

var removeExerciseFromRoutineCmd = &cobra.Command{
	Use:   "remove-exercise-from-routine",
	Short: "Remove an exercise from a routine",
	Run: func(cmd *cobra.Command, args []string) {
		if removeRoutineID == 0 || removeExerciseID == 0 {
			fmt.Println("--routine-id and --exercise-id are required")
			return
		}

		db := db.GetDB()

		var routine models.Routine
		if err := db.First(&routine, removeRoutineID).Error; err != nil {
			fmt.Printf("Routine with ID %d not found.\n", removeRoutineID)
			return
		}

		var exercise models.Exercise
		if err := db.First(&exercise, removeExerciseID).Error; err != nil {
			fmt.Printf("Exercise with ID %d not found.\n", removeExerciseID)
			return
		}

		if err := db.Model(&routine).Association("Exercises").Delete(&exercise); err != nil {
			fmt.Println("Failed to remove exercise from routine:", err)
			return
		}

		fmt.Printf("Exercise '%s' removed from routine '%s'.\n", exercise.Name, routine.Name)
	},
}

func init() {
	rootCmd.AddCommand(createRoutineCmd)
	rootCmd.AddCommand(listRoutinesCmd)
	rootCmd.AddCommand(editRoutineCmd)
	rootCmd.AddCommand(deleteRoutineCmd)
	rootCmd.AddCommand(addExerciseToRoutineCmd)
	rootCmd.AddCommand(listRoutineExercisesCmd)
	rootCmd.AddCommand(removeExerciseFromRoutineCmd)

	createRoutineCmd.Flags().StringVarP(&routineName, "name", "n", "", "Name of the routine")
	createRoutineCmd.MarkFlagRequired("name")

	editRoutineCmd.Flags().UintVarP(&routineID, "id", "i", 0, "ID of the routine to edit")
	editRoutineCmd.Flags().StringVarP(&newName, "name", "n", "", "New name for the routine")
	editRoutineCmd.MarkFlagRequired("id")
	editRoutineCmd.MarkFlagRequired("name")

	deleteRoutineCmd.Flags().UintVarP(&routineID, "id", "i", 0, "ID of the routine to delete")
	deleteRoutineCmd.MarkFlagRequired("id")

	addExerciseToRoutineCmd.Flags().UintVarP(&routineID, "routine-id", "r", 0, "ID of the routine")
	addExerciseToRoutineCmd.Flags().UintSliceVarP(&exerciseIDs, "exercise-ids", "e", []uint{}, "Comma-separated exercise IDs (e.g. --exercise-ids=1,2,3)")
	addExerciseToRoutineCmd.MarkFlagRequired("routine-id")
	addExerciseToRoutineCmd.MarkFlagRequired("exercise-ids")

	listRoutineExercisesCmd.Flags().UintVarP(&routineIDForList, "routine-id", "r", 0, "ID of the routine to list exercises for")
	listRoutineExercisesCmd.MarkFlagRequired("routine-id")

	removeExerciseFromRoutineCmd.Flags().UintVarP(&removeRoutineID, "routine-id", "r", 0, "ID of the routine")
	removeExerciseFromRoutineCmd.Flags().UintVarP(&removeExerciseID, "exercise-id", "e", 0, "ID of the exercise to remove")
	removeExerciseFromRoutineCmd.MarkFlagRequired("routine-id")
	removeExerciseFromRoutineCmd.MarkFlagRequired("exercise-id")
}
