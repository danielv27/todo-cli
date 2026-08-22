/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all TODOs with color highlighting for done and not done items",
	Long: `Display all TODO items from your TODO markdown file.

Done items are shown in green, and not done items are shown in yellow for easy distinction.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, err := todoFilePath()
		if err != nil {
			fmt.Println(err)
			return
		}

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening TODO file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		count := 0
		for scanner.Scan() {
			line := scanner.Text()
			if strings.TrimSpace(line) == "" {
				continue
			}
			count++
			if strings.Contains(line, "- [x] ") {
				// Green check mark for done
				content := strings.SplitN(line, "] ", 2)
				if len(content) == 2 {
					fmt.Printf("\033[32m✔ %s\033[0m\n", content[1])
				} else {
					fmt.Printf("\033[32m✔ %s\033[0m\n", line)
				}
			} else if strings.Contains(line, "- [ ] ") {
				// Red X for not done
				content := strings.SplitN(line, "] ", 2)
				if len(content) == 2 {
					fmt.Printf("\033[31m✘ %s\033[0m\n", content[1])
				} else {
					fmt.Printf("\033[31m✘ %s\033[0m\n", line)
				}
			} else {
				fmt.Println(line)
			}
		}
		if count == 0 {
			fmt.Println("No TODOs found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
