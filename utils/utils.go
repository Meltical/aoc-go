package utils

import (
	"fmt"
	"os"
	"strings"
)

// ReadInput reads the entire input file for a given day
func ReadInput(day int) string {
	filename := fmt.Sprintf("inputs/day%02d.txt", day)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error reading input file %s: %v", filename, err))
	}
	return string(data)
}

// ReadLines reads the input file and returns it as a slice of lines
func ReadLines(day int) []string {
	input := ReadInput(day)
	input = strings.TrimSpace(input)
	return strings.Split(input, "\n")
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Mod returns a positive modulus result, and handles negative a values correctly
func Mod(a, b int) int {
	return (a%b + b) % b
}
