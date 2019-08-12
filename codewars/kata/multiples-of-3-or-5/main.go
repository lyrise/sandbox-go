package main

import "fmt"

func Multiple3And5(number int) int {
	var result int

	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}

	return result
}

func main() {
	fmt.Println(Multiple3And5(10))
}
