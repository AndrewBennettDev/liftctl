package cmd

import (
	"fmt"

	"github.com/AndrewBennettDev/liftctl/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Start the TUI (Text User Interface)",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tui.Start(); err != nil {
			fmt.Println("Failed to launch TUI:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
