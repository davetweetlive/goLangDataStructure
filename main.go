package main

import (
	"fmt"
	"github.com/goLangDataStructure/stack"
)

// Global variables
var (
	mainMenuChoice int
	invalidChoice  bool
)

func main() {
	inputMenu()

	for mainMenuChoice != 0 {

		if !invalidChoice {
			switch mainMenuChoice {
			case 1:
				fmt.Println("Stack using array")
				stack.StackControl()
			case 2:
				fmt.Println("Queue using array")
			case 3:
				fmt.Println("Dequeue using array")
			case 4:
				fmt.Println("Priority queue using array")
			case 5:
				fmt.Println("Single linked list")
			case 6:
				fmt.Println("Circular linked list")
			case 7:
				fmt.Println("Double linked list")
			case 8:
				fmt.Println("Circular double linked list")
			case 9:
				fmt.Println("Tree")
			case 10:
				fmt.Println("Graph")
			default:
				fmt.Println("Please select option which is availabe!")
			}

		} else {
			fmt.Println("Well, that's not an integer value!")
			fmt.Println("Please re-run the program again!")
			return
		}
		inputMenu()
	}
}

/*
 	The inputMenu function does nothing but provides number based menu and also gets user input ans stores in
	mainMenuChoice variable, if the variable is noninteger variable the sets the invalidChoice as true and exits from
	the function
*/
func inputMenu() {
	fmt.Println("")
	fmt.Println("\t\t\t\t\t\t1. Stack using array   \t\t\t 2. Queue using array")
	fmt.Println("\t\t\t\t\t\t3. Dequeue using array \t\t\t 4. Priority queue using array")
	fmt.Println("\t\t\t\t\t\t5. Single linked list  \t\t\t 6. Circular linked list")
	fmt.Println("\t\t\t\t\t\t7. Double linked list  \t\t\t 8. Circular double linked list")
	fmt.Println("\t\t\t\t\t\t9. Tree                \t\t\t 10. Graph")
	fmt.Println("\t\t\t\t\t\t0. Exit from the menu")
	fmt.Println("\t\t\t\t\t\t________________________________________________________________________")
	_, err := fmt.Scanf("%d", &mainMenuChoice)

	if err != nil {
		invalidChoice = true
		return
	}
}
