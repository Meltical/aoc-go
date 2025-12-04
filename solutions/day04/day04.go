package day04

import . "aoc-go/utils"
import "fmt"

func Solve() {
	lines := ReadInputs(4, "\n", false)
	matrix := Parse(lines)

	part1 := solvePart1(Copy(matrix))
	part2 := solvePart2(matrix)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func removeRolls(m *Matrix[rune]) int {
	removedIdx := make([]int, 0, len(m.Data))

	for i, data := range m.Data {
		if data != '@' {
			continue
		}

		adjacentRollsCount := 0
		for _, delta := range Deltas8 {
			x, y := ToXY(m, i)
			adjX, adjY := x+delta[1], y+delta[0]
			adjData, err := Get(m, adjY, adjX)
			if !err && adjData == '@' {
				adjacentRollsCount++
			}

			if adjacentRollsCount == 4 {
				break
			}
		}

		if adjacentRollsCount < 4 {
			removedIdx = append(removedIdx, i)
		}
	}

	for _, idx := range removedIdx {
		m.Data[idx] = '.'
	}
	return len(removedIdx)
}

func solvePart1(m *Matrix[rune]) int {
	return removeRolls(m)
}

func solvePart2(m *Matrix[rune]) int {
	sum := 0
	for removedCount := removeRolls(m); removedCount > 0; removedCount = removeRolls(m) {
		sum += removedCount
	}
	return sum
}
