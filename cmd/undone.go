/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone [TODO ID or text]",
	Short: "Mark a TODO as not done by its ID or a unique part of its text",
	Long: `Mark a TODO item as not done in your TODO markdown file.

You can specify the TODO to mark as not done by either its unique ID or by providing a unique substring of the TODO text.

Usage:
  undone [TODO ID or text]

Examples:
  undone 1a2b
  undone groceries

If multiple TODOs match the provided text, none will be marked as not done and you will need to use the ID or a more specific substring instead.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide the TODO id (hash) or text to mark as not done.")
			return
		}
		idOrText := args[0]
		filePath, err := todoFilePath()
		if err != nil {
			fmt.Println(err)
			return
		}

		input, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		lines := strings.Split(string(input), "\n")
		found := false
		alreadyUndone := false
		matchIdx := -1
		matchCount := 0

		// First, try to match by ID
		for i, line := range lines {
			if strings.Contains(line, "[id:"+idOrText+"]") {
				if strings.Contains(line, "- [x] ") {
					lines[i] = strings.Replace(line, "- [x] ", "- [ ] ", 1)
					found = true
					break
				} else if strings.Contains(line, "- [ ] ") {
					alreadyUndone = true
					break
				}
			}
		}

		// If not found by ID, try to match by substring in the TODO text
		if !found && !alreadyUndone {
			for i, line := range lines {
				if strings.Contains(line, "- [x] ") && strings.Contains(line, idOrText) {
					matchIdx = i
					matchCount++
				}
			}
			if matchCount == 1 {
				lines[matchIdx] = strings.Replace(lines[matchIdx], "- [x] ", "- [ ] ", 1)
				found = true
			} else if matchCount > 1 {
				fmt.Println("Multiple TODOs match the given text. Please be more specific or use the ID.")
				return
			}
		}

		if alreadyUndone {
			fmt.Println("TODO already set to not done.")
			return
		}

		if !found {
			fmt.Printf("No TODO found that matches: %s\n", idOrText)
			return
		}

		output := strings.Join(lines, "\n")
		err = os.WriteFile(filePath, []byte(output), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Printf("Marked TODO '%s' as not done.\n", idOrText)
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)
}
