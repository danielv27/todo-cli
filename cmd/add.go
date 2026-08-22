/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [TODO item]",
	Short: "Add a new TODO item to your TODO file",
	Long: `Add a new unchecked TODO item to your TODO markdown file.

Usage:
  add [TODO item]

Example:
  add Buy groceries
This will append '- [ ] [id:1234] Buy groceries' to your specified TODO file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a TODO item.")
			return
		}
		todo := strings.Join(args, " ")
		filePath, err := todoFilePath()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Count the number of lines in the file to use as a unique salt
		lineCount := 0
		if f, err := os.Open(filePath); err == nil {
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				lineCount++
			}
			f.Close()
		}

		// Generate a 4-character hash-based ID from the description and line count
		h := sha1.New()
		h.Write([]byte(fmt.Sprintf("%s%d", todo, lineCount)))
		hash := hex.EncodeToString(h.Sum(nil))[:4]
		line := "- [ ] [id:" + hash + "] " + todo + "\n"

		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()

		if _, err := f.WriteString(line); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Printf("Added TODO (id:%s): %s\n", hash, todo)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
