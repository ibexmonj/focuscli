package cmd

import (
	"fmt"
	"github.com/ibexmonj/focuscli/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Display session history",
	Long:  `Display the history of focus and break sessions.`,
	Run: func(cmd *cobra.Command, args []string) {
		sessions, err := utils.LoadSessions()
		if err != nil {
			logrus.Errorf("Failed to load session data: %v", err)
			return
		}

		if len(sessions) == 0 {
			fmt.Println("No session history found.")
			return
		}

		fmt.Println("Session History:")
		for _, session := range sessions {
			fmt.Printf("%s - %d minutes on %s\n", session.Type, session.Duration, session.Timestamp.Format("2006-01-02 15:04:05"))
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
