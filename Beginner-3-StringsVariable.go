package main
import "fmt"
func main(){
	// This is var declaration with default val = ""
	var firstName string
	
	
	// This is the complete variable declaration
	var name string = "Andy"

	// This is the simple declaration,
	// where the compiler determins the data type
	// of the variable based on the val at right side
	var name2 = "Cooper"

	// This is the short hand var declaration
	// This cannot be use at the global var decl
	name3:= "Tim"

	// We must use the variables declared inside the prg
	// else will get the error like 
	//'./Beginner-3-StringsVariable.go:14:2: name3 declared but not used' 
	fmt.Printf("FirstName = %v, name=%v, name2=%v,  name3=%v\n",firstName,name,name2,name3)
}