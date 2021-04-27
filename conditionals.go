package main
import "fmt"

func main() {
	sick := true
	healthy := false
	if (healthy) {
		fmt.Println("Work.")
	  }
	if sick {
		fmt.Println("Stay home.")
	}	

	//Go Comparison Operators
	same := 3 == 3 
	notsame := "ABC" != "abc" 
	lessthan := 5 <= -5 
	fmt.Println(same + " - " + notsame + " - " + lessthan)

	//Go Short Variable Declaration

	// A short variable declaration can be made within the scope of an if or switch statement before the condition is specified but after 
	// the if or switch keyword. A semicolon, ;, is appended to the declaration to separate it from the condition
	if age := 55; age >= 55 {
		fmt.Println("You are retiring!")
	}
	  switch season := "spring"; season {
		case "spring":
			fmt.Println("Plant some bulbs.")
		case "summer":
			fmt.Println("xxxxxxxxxxx")
		}
}