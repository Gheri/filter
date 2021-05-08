package main

import "os"

func main() {
	logger := Logger{}
	filename := os.Getenv("FILTER_FILE_NAME")
	if filename == "" {
		logger.Info("Env variable FILTER_FILE_NAME is not set.")
		filename = "data.json"
	}
	users, err := NewUsers(filename)
	if err != nil {
		logger.Error(err)
	}
	invalidIds := users.GetInvalidIds()
	for _, id := range invalidIds {
		logger.Info(id)
	}
}
