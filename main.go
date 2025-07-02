package main

import (
	"bufio"
	"fmt"
	"os"
)

func compareUniqueLines(file1Path, file2Path string) error {
	file1, err := os.Open(file1Path)
	if err != nil {
		return fmt.Errorf("error opening file1: %v", err)
	}
	defer file1.Close()

	file2, err := os.Open(file2Path)
	if err != nil {
		return fmt.Errorf("error opening file2: %v", err)
	}
	defer file2.Close()

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	lines1 := make(map[string]int)
	lines2 := make(map[string]int)
	var order1, order2 []string

	for i := 1; scanner1.Scan(); i++ {
		line := scanner1.Text()
		lines1[line] = i
		order1 = append(order1, line)
	}
	if err := scanner1.Err(); err != nil {
		return fmt.Errorf("error reading file1: %v", err)
	}

	for i := 1; scanner2.Scan(); i++ {
		line := scanner2.Text()
		lines2[line] = i
		order2 = append(order2, line)
	}
	if err := scanner2.Err(); err != nil {
		return fmt.Errorf("error reading file2: %v", err)
	}

	hasUnique := false

	for _, line := range order1 {
		if _, exists := lines2[line]; !exists {
			hasUnique = true
			fmt.Printf("L%d: %s: %s\n", lines1[line], file1Path, line)
		}
	}

	for _, line := range order2 {
		if _, exists := lines1[line]; !exists {
			hasUnique = true
			fmt.Printf("L%d: %s: %s\n", lines2[line], file2Path, line)
		}
	}

	if !hasUnique {
		fmt.Println("No unique lines found.")
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run compare_unique_lines.go <file1> <file2>")
		os.Exit(1)
	}
	file1Path := os.Args[1]
	file2Path := os.Args[2]
	if err := compareUniqueLines(file1Path, file2Path); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
