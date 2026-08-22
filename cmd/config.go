/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
)

// todoFilePath returns the path to the TODO markdown file, read from the
// TODO_CLI_FILE environment variable, falling back to ~/todo.md.
func todoFilePath() (string, error) {
	if path := os.Getenv("TODO_CLI_FILE"); path != "" {
		return path, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error finding home directory: %w", err)
	}
	return homeDir + "/todo.md", nil
}
