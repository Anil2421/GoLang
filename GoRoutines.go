/*
This method shows the flow of methods,
1. All the methods which are prefixed as go, are termed as goroutines.
2. These acts as a parallel threads which executes seperately
3. At line 25, we said main() method to wait and let other goroutines do their jobs
4. If line 25 is commented, then the goroutines will not be able to finish their tasks
and the prg will terminate as soon as the main terminates
*/
package main

import (
	"fmt"
	"time"
)

func returnSquare(val int16) int16 {
	fmt.Println(val)
	time.Sleep(100 * time.Second)
	return val * val
}
func main() {
	var return1 int16
	fmt.Println("Hello, Showing demo of Go routines")
	go returnSquare(12)
	go returnSquare(199)
	return1 = returnSquare(900)
	time.Sleep(1 * time.Second)
	fmt.Println(return1)
}
