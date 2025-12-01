package main

import (
	"aoc-go/solutions/day01"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		fmt.Println("Example: go run main.go 1")
		fmt.Println("Or: go run main.go all")
		os.Exit(1)
	}

	arg := os.Args[1]

	if arg == "all" {
		runAll()
		return
	}

	day, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Invalid day: %s\n", arg)
		os.Exit(1)
	}

	runDay(day)
}

func runDay(day int) {
	fmt.Printf("=== Day %d ===\n", day)

	switch day {
	case 1:
		day01.Solve()
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
	}
}

func runAll() {
	fmt.Println("Running all implemented solutions...")
	fmt.Println()

	days := []int{1}

	for _, day := range days {
		runDay(day)
		fmt.Println()
	}
}
