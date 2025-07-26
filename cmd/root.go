package cmd

import (
	"fmt"
	"os"

	"github.com/AndrewBennettDev/liftctl/internal/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "liftctl",
	Short: "Workout tracker CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing database...")
		db.Init()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Execution error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
