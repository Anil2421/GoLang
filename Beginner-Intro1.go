// This is the 1st section,
// also know as the package declaration section
// This line means that the below written code belongs 
// to which pkg
package main

/*
This is the 2nd section,
also known as the package import section.
This line means that the following packages will be
used in this file
*/
import "fmt"

/*
This is the 3rd section,
also known as the entry point of the program.
The program execution begins from this point onwards.
Every project must have a single main function.
*/
func main(){
	// We are using "P"rintln method from "fmt" pkg
	// Note, "P"rintln starts with uppercase
	// this means that the Println method is public as it starts with uppercase
	fmt.Println("Hello Team!!")
}