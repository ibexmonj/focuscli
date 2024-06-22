package cmd

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/ibexmonj/focuscli/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

var sessionDuration int

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a focus session",
	Long:  `Start a focus session of a specified duration in minutes using the Pomodoro Technique.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("Starting a %d-minute focus session...", sessionDuration)
		time.Sleep(time.Duration(sessionDuration) * time.Minute)
		fmt.Println("Focus session ended. Time for a break!")

		err := beeep.Alert("FocusCLI", "Focus session ended. Time for a break!", "")
		if err != nil {
			logrus.Errorf("Failed to send notification: %v", err)
		}

		// Save session data
		session := utils.Session{
			Type:      "focus",
			Duration:  sessionDuration,
			Timestamp: time.Now(),
		}
		err = utils.SaveSession(session)
		if err != nil {
			logrus.Errorf("Failed to save session data: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().IntVarP(&sessionDuration, "session", "s", 25, "Duration of the focus session in minutes")
}
