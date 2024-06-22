package cmd

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/ibexmonj/focuscli/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

var breakDuration int

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Start a break session",
	Long:  `Start a break session of a specified duration in minutes.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("Starting a %d-minute break...", breakDuration)
		time.Sleep(time.Duration(breakDuration) * time.Minute)
		fmt.Println("Break session ended. Ready to focus again!")

		err := beeep.Alert("FocusCLI", "Break session ended. Ready to focus again!", "")
		if err != nil {
			logrus.Errorf("Failed to send notification: %v", err)
		}

		// Save session data
		session := utils.Session{
			Type:      "break",
			Duration:  breakDuration,
			Timestamp: time.Now(),
		}
		err = utils.SaveSession(session)
		if err != nil {
			logrus.Errorf("Failed to save session data: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(breakCmd)

	breakCmd.Flags().IntVarP(&breakDuration, "duration", "d", 5, "Duration of the break in minutes")
}
