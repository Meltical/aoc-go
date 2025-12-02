package day02

import (
	"aoc-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	ranges := utils.ReadInputs(2, ",", false)
	boundsInt := make([][2]int, len(ranges))

	for i, input := range ranges {
		bounds := strings.Split(input, "-")
		for j, bound := range bounds {
			val, _ := strconv.Atoi(bound)
			boundsInt[i][j] = val
		}
	}

	part1, part2 := solveAll(boundsInt)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func isInvalidID(strId string, partCount int) bool {
	if len(strId)%partCount != 0 {
		return false
	}

	partLen := len(strId) / partCount
	partToMatch := strId[:partLen]

	for i := 1; i < partCount; i++ {
		start := i * partLen
		end := start + partLen
		if strId[start:end] != partToMatch {
			return false
		}
	}

	return true
}

func solveAll(ints [][2]int) (int, int) {
	part1 := 0
	part2 := 0

	for i := 0; i < len(ints); i++ {
		intPair := ints[i]

		for id := intPair[0]; id <= intPair[1]; id++ {
			strId := strconv.Itoa(id)

			if isInvalidID(strId, 2) {
				part1 += id
				part2 += id
				continue
			}

			for partCount := 3; partCount <= len(strId); partCount++ {
				if isInvalidID(strId, partCount) {
					part2 += id
					break
				}
			}
		}
	}

	return part1, part2
}
