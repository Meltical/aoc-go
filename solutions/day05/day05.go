package day05

import (
	"aoc-go/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	groups := utils.ReadInputs(5, "\n\n", false)
	ranges := make([][2]int, 0)
	for _, line := range strings.Split(groups[0], "\n") {
		ids := [2]int{}
		for i, id := range strings.Split(line, "-") {
			ids[i], _ = strconv.Atoi(id)
		}
		ranges = append(ranges, ids)
	}

	ingredientIds := make([]int, 0)
	for _, line := range strings.Split(groups[1], "\n") {
		id, _ := strconv.Atoi(line)
		ingredientIds = append(ingredientIds, id)
	}

	part1 := solvePart1(ranges, ingredientIds)
	part2 := solvePart2(ranges)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func isInRange(rng [2]int, id int) bool {
	return id >= rng[0] && id <= rng[1]
}

func isInAnyRange(ranges [][2]int, id int) bool {
	for _, rng := range ranges {
		if isInRange(rng, id) {
			return true
		}
	}
	return false
}

func solvePart1(ranges [][2]int, ingredientIds []int) int {
	count := 0
	for _, id := range ingredientIds {
		if isInAnyRange(ranges, id) {
			count++
		}
	}
	return count
}

func countFreshIds(ranges [][2]int) int {
	count := 0
	for _, rng := range ranges {
		d := rng[1] - rng[0] + 1
		count += d
	}
	return count
}

func solvePart2(ranges [][2]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})
	for i := 0; ; i++ {
		if i == len(ranges)-1 {
			break
		}
		if ranges[i][1] >= ranges[i+1][0] {
			ranges[i][1] = utils.Max(ranges[i][1], ranges[i+1][1])
			ranges = append(ranges[:i+1], ranges[i+2:]...)
			i--
		}
	}
	return countFreshIds(ranges)
}
