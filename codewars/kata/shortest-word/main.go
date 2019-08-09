package main

import (
	"math"
	"strings"
)

func FindShort(s string) int {
	count := math.MaxInt64
	slice := strings.Split(s, " ")
	for _, str := range slice {
		tmp := len(str)
		if count > tmp {
			count = tmp
		}
	}

	return count
}

func main() {
}
