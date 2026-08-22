/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple CLI for managing TODOs in a markdown file",
	Long: `todo-cli manages TODO items in a plain markdown file, using
GitHub-style checkbox syntax ('- [ ]' / '- [x]').

By default it reads and writes ~/todo.md. Set the TODO_CLI_FILE
environment variable to point it at a different file.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
