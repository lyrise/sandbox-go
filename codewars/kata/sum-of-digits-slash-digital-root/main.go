package main

import (
	"fmt"
	"strconv"
)

func DigitalRoot(n int) int {
	text := strconv.Itoa(n)

	for {
		result := 0
		for _, w := range text {
			d, _ := strconv.Atoi(string(w))
			result += d
		}

		text = strconv.Itoa(result)
		if len(text) == 1 {
			return result
		}
	}
}

func main() {
	fmt.Println(DigitalRoot(23))
}
