package modules

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

var (
	Config       ConfigS
	ConfigLoaded = false
)

func mustGetEnv(key string, isPath bool) string {
	value := os.Getenv(key)

	if value == "" {
		panic("Missing environment variable: " + key)
	}

	if isPath {
		found, err := isValidPath(value)

		if !found || err != nil {
			panic("Missing or invalid environment variable: " + key)
		}

	}

	return value
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	Config.LinksFilePath = mustGetEnv("LinksFilePath", true)
	Config.ZippedDirPath = mustGetEnv("ZippedDirPath", true)
	Config.RvzDirPath = mustGetEnv("RvzDirPath", true)
	Config.IsoDirPath = mustGetEnv("IsoDirPath", true)
	Config.WBFSDirPath = mustGetEnv("WBFSDirPath", true)

	ConfigLoaded = true
}

func isValidPath(path string) (bool, error) {
	absPath, err := filepath.Abs(path) // Get the absolute path

	if err != nil {
		return false, err
	}

	_, err = os.Stat(absPath) // Check if the path exists

	if os.IsNotExist(err) {
		return false, nil // Path does not exist
	} else if err != nil {
		return false, err // Some other error (permissions, etc.)
	}

	return true, nil // Path exists and is valid
}