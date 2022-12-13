// Get a number from the console and check if it’s between 1 and 10.

/* Caution
In order to run this prg, we have to:
- The pkg slices needed to be installed, to get it installed we have to
		- create new go.mod file--> cmd---> go init mod <ModuleName>
		- go mod tidy
- Read the user input  in str format
- convert it using strconv.Atio method, return type of this is int not int8
-Check the user input lies in between the slice using slices.Contains method

*/

package main

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"
)

func CheckNo1to10() {
	var val string
	fmt.Scanf("%s", &val)
	validValues := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	conVal, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("Unable to convert str to int")
	}
	if slices.Contains(validValues, int8(conVal)) {
		//if val in validValues{
		fmt.Println("Value lies in-between 1-10")
	} else {
		fmt.Println("Value doesn't lies in-between 1-10")
	}
}
func main() {
	CheckNo1to10()
}
