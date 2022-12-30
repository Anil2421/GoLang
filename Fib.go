// Create a fib. method
package main

import "fmt"

func fib(v uint64) uint64 {
	if v == 1 {
		return 1
	} else {
		v = v * fib(v-1)
	}
	return v
}
func main() {
	fmt.Println("Enter the no, to calculate the fibnacci: ")
	var val uint64
	fmt.Scanf("%d\n", &val)
	a := fib(val)
	println(a)
}
