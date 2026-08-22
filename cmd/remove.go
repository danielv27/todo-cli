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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [TODO ID or text]",
	Short: "Remove a TODO by its ID or a unique part of its text",
	Long: `Delete a TODO item from your TODO markdown file.

You can specify the TODO to remove by either its unique hash ID or by providing a unique substring of the TODO text.

Usage:
  remove [TODO ID or text]

Examples:
  remove 1a2b
  remove groceries

If multiple TODOs match the provided text, none will be removed and you will need to use the ID instead.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide the TODO id (hash) or text to remove.")
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
		matchIdx := -1
		matchCount := 0

		// Try to match by ID
		for i, line := range lines {
			if strings.Contains(line, "[id:"+idOrText+"]") {
				matchIdx = i
				matchCount = 1
				break
			}
		}

		// If not found by ID, try to match by substring in the TODO text
		if matchCount == 0 {
			for i, line := range lines {
				if strings.Contains(line, idOrText) && (strings.Contains(line, "- [ ] ") || strings.Contains(line, "- [x] ")) {
					if matchIdx == -1 {
						matchIdx = i
					}
					matchCount++
				}
			}
		}

		if matchCount == 0 {
			fmt.Printf("No TODO found that matches: %s\n", idOrText)
			return
		} else if matchCount > 1 {
			fmt.Println("Multiple TODOs match the given text. Please be more specific or use the ID.")
			return
		}

		// Remove the matched line
		lines = append(lines[:matchIdx], lines[matchIdx+1:]...)

		output := strings.Join(lines, "\n")
		err = os.WriteFile(filePath, []byte(output), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Printf("Removed TODO '%s'.\n", idOrText)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
