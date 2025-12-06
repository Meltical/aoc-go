package utils

import (
	"fmt"
	"os"
	"strings"
)

// ReadInput reads the entire input file for a given day
func ReadInput(day int, test bool) string {
	filename := ""
	if test {
		filename = fmt.Sprintf("inputs/day%02d-test.txt", day)
	} else {

		filename = fmt.Sprintf("inputs/day%02d.txt", day)
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error reading input file %s: %v", filename, err))
	}
	return string(data)
}

// ReadInputs reads the input file and returns it as a slice of lines
func ReadInputs(day int, separator string, test bool) []string {
	input := ReadInput(day, test)
	input = strings.TrimSpace(input)
	return strings.Split(input, separator)
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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Transpose[T any](A [][]T) [][]T {
	N := len(A)
	M := len(A[0])

	AT := make([][]T, M)
	for i := 0; i < M; i++ {
		AT[i] = make([]T, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			AT[j][i] = A[i][j]
		}
	}
	return AT
}
