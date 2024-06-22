/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ibexmonj/focuscli/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize the logger
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)

	logrus.Info("Starting FocusCLI")

	cmd.Execute()
}
