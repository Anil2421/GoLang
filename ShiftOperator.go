package main

import (
	"fmt"
)

func main() {
	var num uint8
	var result uint8
	var i uint8
	fmt.Print("Enter a decimal number\t:\t")
	fmt.Scan(&num)
	for i = 31; i >= 0; i-- {
		result = num << i
		if result&1 == 1 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}
