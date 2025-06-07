package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Gets the user's $XDG_DATA_HOME dir.
//
// Fallsback to the default data dir if the env var does not exist.
func UserDataDir() string {
	var dataDir string

	if dataDir = os.Getenv("XDG_DATA_HOME"); dataDir != "" {
		return dataDir
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not get the home directory")
	}

	switch runtime.GOOS {
	case "windows":
		if dataDir = os.Getenv("LOCALAPPDATA"); dataDir != "" {
			dataDir = filepath.Join(home, "AppData", "Local")
		}
	default:
		dataDir = filepath.Join(home, ".local", "share")
	}

	return dataDir
}

func ID(formID, name, component, element string) string {
	return fmt.Sprintf("%s-%s-%s-%s", formID, name, component, element)
}
