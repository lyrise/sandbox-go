package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	fiveCount := 0
	sevenCount := 0

	for i := 0; i < 3; i++ {
		x := nextInt()
		if x == 5 {
			fiveCount++
		} else if x == 7 {
			sevenCount++
		}
	}

	if fiveCount == 2 && sevenCount == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
