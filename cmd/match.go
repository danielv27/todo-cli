/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import "strings"

// findByID returns the index of the line tagged with the given TODO id,
// or -1 if no line has that id.
func findByID(lines []string, id string) int {
	for i, line := range lines {
		if strings.Contains(line, "[id:"+id+"]") {
			return i
		}
	}
	return -1
}

// findBySubstring returns the index of the single line that contains substr
// and has at least one of the given markers (e.g. "- [ ] "), along with how
// many lines matched. If zero or more than one line matched, idx is -1.
func findBySubstring(lines []string, substr string, markers ...string) (idx int, count int) {
	idx = -1
	for i, line := range lines {
		if !strings.Contains(line, substr) {
			continue
		}
		for _, marker := range markers {
			if strings.Contains(line, marker) {
				if idx == -1 {
					idx = i
				}
				count++
				break
			}
		}
	}
	if count != 1 {
		idx = -1
	}
	return idx, count
}
