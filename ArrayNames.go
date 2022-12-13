// Create an array with the number 0 to 10
package main

import "fmt"

func main() {
	var intArray [10]int
	for i := 0; i < 10; i++ {
		intArray[i] = i + 1
	}
	fmt.Println(intArray)
}
