package utils

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

const dataFile = "sessions.json"

type Session struct {
	Type      string    `json:"type"`
	Duration  int       `json:"duration"`
	Timestamp time.Time `json:"timestamp"`
}

func SaveSession(session Session) error {
	var sessions []Session

	// Check if the data file exists
	if _, err := os.Stat(dataFile); err == nil {
		file, err := os.Open(dataFile)
		if err != nil {
			return err
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &sessions)
		if err != nil {
			return err
		}
	}

	// Append the new session
	sessions = append(sessions, session)

	// Save back to file
	data, err := json.Marshal(sessions)
	if err != nil {
		return err
	}

	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadSessions() ([]Session, error) {
	var sessions []Session

	// Check if the data file exists
	if _, err := os.Stat(dataFile); err == nil {
		file, err := os.Open(dataFile)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &sessions)
		if err != nil {
			return nil, err
		}
	}

	return sessions, nil
}
