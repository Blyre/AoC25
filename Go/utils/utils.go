package utils

import (
	"os"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func OpenAndReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	CheckErr(err)
	return string(data)
}

func StringtToMatrix(content string) [][]rune {
	lines := strings.Split(content, "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
}
