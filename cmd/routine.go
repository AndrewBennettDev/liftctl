package cmd

import (
	"fmt"

	"github.com/AndrewBennettDev/liftctl/internal/db"
	"github.com/AndrewBennettDev/liftctl/internal/models"
	"github.com/spf13/cobra"
)

var routineName string

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
		fmt.Println("Edit routine functionality is not implemented yet.")
	},
}

var deleteRoutineCmd = &cobra.Command{
	Use:   "delete-routine",
	Short: "Delete an existing routine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete routine functionality is not implemented yet.")
	},
}

var addExerciseToRoutineCmd = &cobra.Command{
	Use:   "add-exercise-to-routine",
	Short: "Add an exercise to a routine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add exercise to routine functionality is not implemented yet.")
	},
}

func init() {
	rootCmd.AddCommand(createRoutineCmd)
	rootCmd.AddCommand(listRoutinesCmd)
	rootCmd.AddCommand(editRoutineCmd)
	rootCmd.AddCommand(deleteRoutineCmd)
	rootCmd.AddCommand(addExerciseToRoutineCmd)

	createRoutineCmd.Flags().StringVar(&routineName, "name", "", "Name of the routine")
}
