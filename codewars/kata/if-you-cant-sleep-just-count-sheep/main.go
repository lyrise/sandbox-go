package main

import (
	"fmt"
	"strings"
)

func countSheep(num int) string {
	var builder strings.Builder
	for i := 1; i <= num; i++ {
		fmt.Fprintf(&builder, "%d sheep...", i)
	}
	return builder.String()
}

func main() {
	fmt.Println(countSheep(0))
	fmt.Println(countSheep(1))
	fmt.Println(countSheep(3))
}
