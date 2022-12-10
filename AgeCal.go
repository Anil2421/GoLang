// Create a program that calculates the average age of 5 people
package main

import "fmt"

func CalculateAverageAge() {
	ageSlice := []int32{24, 26, 25, 22, 30}
	var avg float32 = 0
	for _, val := range ageSlice {
		//avg += float32(val)
		avg = avg + float32(val)
	}
	fmt.Println("Average = ", avg/float32(len(ageSlice)))
}
func main() {
	CalculateAverageAge()
}
