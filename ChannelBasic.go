package main

import (
	"fmt"
	"time"
)

var names [4]string = [4]string{"Lisa", "Mona", "Lully", "Tram"}

// Populating these channels 2 times
func Intake(n chan string) {
	n <- names[3]
	time.Sleep(20 * time.Millisecond)
	n <- names[2]
}

// Populating these channels 2 times
func InIndex(i chan int) {
	i <- 3
	time.Sleep(20 * time.Millisecond)
	i <- 2
}
func main() {
	fmt.Print("Channel Testing")
	var name = make(chan string)
	var index = make(chan int)
	// These below methosa must be executed as go routines. Calling them as a simple methos won't work
	go Intake(name)
	go InIndex(index)
	// Consuming these channel values 2 times
	fmt.Printf("\nIntake chan strinng = %s and its index = %d", <-name, <-index)

	fmt.Printf("\nIntake chan strinng = %s and its index = %d", <-name, <-index)

	// If tried to get the val from the channel 3rd time, it will show msg as "fatal error: all goroutines are asleep - deadlock!""
}
