package utils

import (
	"strings"
)

type Matrix[T any] struct {
	Rows int
	Cols int
	Data []T
}

var Deltas4 = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var Deltas8 = [8][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

func Copy[T any](m *Matrix[T]) *Matrix[T] {
	dataCopy := make([]T, len(m.Data))
	copy(dataCopy, m.Data)
	return &Matrix[T]{
		Rows: m.Rows,
		Cols: m.Cols,
		Data: dataCopy,
	}
}

func Parse(input []string) *Matrix[rune] {
	rows := len(input)
	cols := len(input[0])
	data := make([]rune, 0, rows*cols)
	for _, line := range input {
		for _, char := range line {
			data = append(data, char)
		}
	}
	return &Matrix[rune]{
		Rows: rows,
		Cols: cols,
		Data: data,
	}
}

func Log(m *Matrix[rune]) {
	var sb strings.Builder
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			char, _ := Get(m, r, c)
			sb.WriteRune(char)
		}
		sb.WriteString("\n")
	}
	println(sb.String())
}

func Get[T any](m *Matrix[T], row, col int) (T, bool) {
	if !IsValid(m, row, col) {
		var zero T
		return zero, true
	}
	return m.Data[row*m.Cols+col], false
}

func IsValid[T any](m *Matrix[T], row, col int) bool {
	return row >= 0 && row < m.Rows && col >= 0 && col < m.Cols
}

func ToXY[T any](m *Matrix[T], index int) (int, int) {
	return index % m.Cols, index / m.Cols
}
