package cmd

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
	"time"
)

var breakDuration int

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Start a break session",
	Long:  `Start a break session of a specified duration in minutes.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting a %d-minute break...\n", breakDuration)
		time.Sleep(time.Duration(breakDuration) * time.Minute)
		fmt.Println("Break session ended. Ready to focus again!")
		err := beeep.Alert("FocusCLI", "Break session ended. Ready to focus again!", "")
		if err != nil {
			fmt.Println("Failed to send notification:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(breakCmd)

	breakCmd.Flags().IntVarP(&breakDuration, "duration", "d", 5, "Duration of the break in minutes")
}
