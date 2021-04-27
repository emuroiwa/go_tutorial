package main

import "fmt"

func main() {
	var fruit string
	string = "apple"
	var fruit string = "apple"
	var fruit = "apple"
	//skip the var keyword, define a name followed by := and value and let the compiler infer its type.
	fruit := "apple"

	fmt.Println(fruit)

	//Go Multiple Variable Declaration
	found, answer := true, "yes"
	var name, age = "Steve", 35
	fmt.Println(found, answer, name, age)

	var firstName string = "Abe"
	var lastName string = "Lincoln"
	
	// prints "Abe Lincoln"
	fmt.Println(firstName + " " + lastName)
}
