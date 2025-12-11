package utils

import "os"

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
