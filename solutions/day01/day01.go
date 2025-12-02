package day01

import (
	"aoc-go/utils"
	"fmt"
	"strconv"
)

func Solve() {
	lines := utils.ReadInputs(1, "\n", false)
	inputs := make([]int, 0, len(lines))

	for _, val := range lines {
		dir := 1
		if val[0] == 'L' {
			dir = -1
		}
		rotation, _ := strconv.Atoi(val[1:])
		inputs = append(inputs, rotation*dir)
	}

	part1, part2 := solveAll(inputs)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func solveAll(inputs []int) (int, int) {
	passedZeroCount := 0
	onZeroCount := 0
	value := 50

	for _, rot := range inputs {
		d := value + rot
		passedZeroCount += utils.Abs(d) / 100
		if value != 0 && d <= 0 {
			passedZeroCount++
		}
		if value == 0 {
			onZeroCount++
		}
		value = utils.Mod(d, 100)
	}

	return onZeroCount, passedZeroCount
}
