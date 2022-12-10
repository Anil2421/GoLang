/*
In this prg, we will see the functionality of the waitgroup.
WaitGroup is the concurrency pattern,
This include 4 steps:
1. creating a global wg obj from sync.WaitGroup class (at line 19)
2. add the counter saying wg.Add(countVal) (at line 30)
3. Inform main thread to wait until the goroutines are finished executing by, wg.Wait() (at line 34)
4. Inform the main thread that the goroutine have finished his job by, wg.Done() (at line 26)
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = new(sync.WaitGroup)

func printName(val string) {
	for i := 0; i < 10; i++ {
		time.Sleep(5 * time.Second)
		fmt.Printf("\nHello teamMate number - %v from %v Team \n", i, val)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	fmt.Print("Starting the Main App")
	go printName("Blue")
	go printName("Red")
	wg.Wait()
	fmt.Print("Now exiting the Main App")
}
