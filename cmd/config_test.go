/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTodoFilePathUsesEnvVar(t *testing.T) {
	t.Setenv("TODO_CLI_FILE", "/tmp/custom-todo.md")

	path, err := todoFilePath()
	if err != nil {
		t.Fatalf("todoFilePath() returned error: %v", err)
	}
	if path != "/tmp/custom-todo.md" {
		t.Errorf("todoFilePath() = %q, want %q", path, "/tmp/custom-todo.md")
	}
}

func TestTodoFilePathDefaultsToHome(t *testing.T) {
	t.Setenv("TODO_CLI_FILE", "")

	path, err := todoFilePath()
	if err != nil {
		t.Fatalf("todoFilePath() returned error: %v", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("os.UserHomeDir() returned error: %v", err)
	}
	want := filepath.Join(homeDir, "todo.md")
	if path != want {
		t.Errorf("todoFilePath() = %q, want %q", path, want)
	}
}
