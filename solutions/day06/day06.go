package day06

import (
	"aoc-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	lines := utils.ReadInputs(6, "\n", false)
	operators := make([]string, 0)
	for _, field := range strings.Fields(lines[len(lines)-1]) {
		operators = append(operators, field)
	}

	part1 := solvePart1(lines, operators)
	part2 := solvePart2(lines, operators)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func solvePart1(lines []string, operators []string) int {
	numbers := make([][]int, 0)
	for i := 0; i < len(lines)-1; i++ {
		lineNumbers := make([]int, 0)
		for _, field := range strings.Fields(lines[i]) {
			num, _ := strconv.Atoi(field)
			lineNumbers = append(lineNumbers, num)
		}
		numbers = append(numbers, lineNumbers)
	}

	total := 0
	for i, operator := range operators {
		columnCount := applyOperator(numbers[0][i], numbers[1][i], operator)
		for j := 2; j < len(numbers); j++ {
			columnCount = applyOperator(columnCount, numbers[j][i], operator)
		}
		total += columnCount
	}

	return total
}

func solvePart2(lines []string, operators []string) int {
	matrix := make([][]rune, 0)
	numbersRowsLength := len(lines) - 1
	for i := 0; i < numbersRowsLength; i++ {
		matrix = append(matrix, []rune(lines[i]))
	}
	newMatrix := utils.Transpose(matrix)
	numbers := make([][]int, 0)
	columnNumbers := make([]int, 0)
	for _, row := range newMatrix {
		fields := strings.Fields(string(row))
		if len(fields) == 0 {
			numbers = append(numbers, columnNumbers)
			columnNumbers = make([]int, 0)
			continue
		}
		num, _ := strconv.Atoi(fields[0])
		columnNumbers = append([]int{num}, columnNumbers...)
	}
	numbers = append(numbers, columnNumbers)

	total := 0
	for i, operator := range operators {
		columnCount := applyOperator(numbers[i][0], numbers[i][1], operator)
		for j := 2; j < len(numbers[i]); j++ {
			columnCount = applyOperator(columnCount, numbers[i][j], operator)
		}
		total += columnCount
	}
	return total
}

func applyOperator(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}
