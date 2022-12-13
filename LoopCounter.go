//Make a program that counts from 1 to 10.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var counterString string
	var sum int = 0
	fmt.Scanf("%s", &counterString)
	counter, _ := strconv.Atoi(counterString)

	for i := 1; i <= counter; i++ {
		sum += i
	}
	fmt.Println("Sum = ", sum)
}
