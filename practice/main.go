package main

import "fmt"

// Variables

// Constant or global variable
const myConst = 100

func main() {
	var aString string = "This is go!"
	fmt.Println(aString)
	fmt.Printf("The variable type is %T\n", aString)

	var anInteger int = 47
	fmt.Printf("The number variable type is %T\n", anInteger)

	// A number (i.e. int) will default to 0
	var defaultInt int
	fmt.Println(defaultInt)

	// The compiler figures the variable type
	var anotherString = "This is another string."
	fmt.Println(anotherString)
	fmt.Printf("Variable type %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("Variable type %T\n", anotherInt)

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("Variable type %T\n", myString)

	fmt.Println(myConst)
	fmt.Printf("Variable type %T\n", myConst)
}
