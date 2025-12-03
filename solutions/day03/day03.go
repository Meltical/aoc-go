package day03

import (
	"aoc-go/utils"
	"fmt"
	"math"
)

func Solve() {
	lines := utils.ReadInputs(3, "\n", false)

	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func solvePart1(lines []string) int {
	part1 := 0

	for _, line := range lines {
		length := len(line)
		leftMax := 0
		rightMax := 0

		for i := 0; i < length-1; i++ {
			val := int(line[i] - '0')
			nextVal := int(line[i+1] - '0')

			if val > leftMax {
				leftMax = val
				rightMax = 0
			}

			if nextVal > rightMax {
				rightMax = nextVal
			}
		}

		part1 += leftMax*10 + rightMax
	}

	return part1
}

const joltLength = 12

func solvePart2(lines []string) int {
	part2 := 0

	for _, line := range lines {
		length := len(line)
		jolts := make([]int, joltLength)
		leftMaxIdx := 0

		for joltIdx := 0; joltIdx < joltLength; joltIdx++ {
			lastPossibleJoltIdx := length - (joltLength - joltIdx)

			for i := leftMaxIdx; i <= lastPossibleJoltIdx; i++ {
				val := int(line[i] - '0')

				if val > jolts[joltIdx] {
					jolts[joltIdx] = val
					leftMaxIdx = i + 1
				}
			}
		}

		for i, _ := range jolts {
			part2 += int(math.Pow10(i)) * jolts[joltLength-i-1]
		}
	}

	return part2
}
