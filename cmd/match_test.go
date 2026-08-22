/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import "testing"

func TestFindByID(t *testing.T) {
	lines := []string{
		"- [ ] [id:aaaa] Buy groceries",
		"- [x] [id:bbbb] Walk the dog",
	}

	if idx := findByID(lines, "bbbb"); idx != 1 {
		t.Errorf("findByID(bbbb) = %d, want 1", idx)
	}
	if idx := findByID(lines, "cccc"); idx != -1 {
		t.Errorf("findByID(cccc) = %d, want -1", idx)
	}
}

func TestFindBySubstring(t *testing.T) {
	lines := []string{
		"- [ ] [id:aaaa] Buy groceries",
		"- [ ] [id:bbbb] Buy stamps",
		"- [x] [id:cccc] Buy milk",
	}

	// Unique match.
	idx, count := findBySubstring(lines, "groceries", "- [ ] ")
	if idx != 0 || count != 1 {
		t.Errorf("findBySubstring(groceries) = (%d, %d), want (0, 1)", idx, count)
	}

	// Ambiguous match: two pending lines contain "Buy".
	idx, count = findBySubstring(lines, "Buy", "- [ ] ")
	if idx != -1 || count != 2 {
		t.Errorf("findBySubstring(Buy) = (%d, %d), want (-1, 2)", idx, count)
	}

	// No match: marker restricts to pending, but "milk" is only on a done line.
	idx, count = findBySubstring(lines, "milk", "- [ ] ")
	if idx != -1 || count != 0 {
		t.Errorf("findBySubstring(milk) = (%d, %d), want (-1, 0)", idx, count)
	}

	// Multiple markers: matches a line with either marker.
	idx, count = findBySubstring(lines, "milk", "- [ ] ", "- [x] ")
	if idx != 2 || count != 1 {
		t.Errorf("findBySubstring(milk, both markers) = (%d, %d), want (2, 1)", idx, count)
	}
}
