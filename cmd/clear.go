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

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Wipe your entire TODO list after confirmation",
	Long: `This command will erase all TODOs from your TODO markdown file after a confirmation prompt.

Usage:
  clear

You will be asked to confirm before the file is wiped.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, err := todoFilePath()
		if err != nil {
			fmt.Println(err)
			return
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Are you sure you want to clear your entire TODO list? This cannot be undone! (y/N): ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))
		if response != "y" && response != "yes" {
			fmt.Println("Aborted. Your TODO list was not cleared.")
			return
		}

		err = os.WriteFile(filePath, []byte{}, 0644)
		if err != nil {
			fmt.Println("Error clearing TODO file:", err)
			return
		}
		fmt.Println("Your TODO list has been cleared.")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
