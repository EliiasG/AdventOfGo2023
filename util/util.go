package util

import (
	"os"
)

func GetInput(name string) string {
	res, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
