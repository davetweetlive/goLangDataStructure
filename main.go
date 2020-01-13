package main

import (
	"fmt"
	"reflect"
)

var mainMenuChoice int

func main() {
	inputMenu()

	for mainMenuChoice != 0  {
		fmt.Println(reflect.TypeOf(mainMenuChoice))
		if reflect.TypeOf(mainMenuChoice) == reflect.TypeOf(1) {
			// switch mainMenuChoice {
			// case 1:
			// 	fmt.Println("Single Linked List")
			// case 2:
			// 	fmt.Println("Circular Linked List")
			// case 3:
			// 	fmt.Println("Double Linked List")
			// case 4:
			// 	fmt.Println("Circular Double Linked List")
			// case 5:
			// 	fmt.Println("Stacks using Arrays")
			// case 6:
			// 	fmt.Println("Queue using Arrays")
			// case 7:
			// 	fmt.Println("Dequeue using arrays")
			// default:
			// 	fmt.Println("Please select option which is availabe!")
			// }
			fmt.Println("The value is int")

		} else {
			fmt.Println("That's a wrong input, you need to pick a number from the menu!")
		}
		inputMenu()
	}
}

func inputMenu() {
	fmt.Println("")
	fmt.Println("\t\t\t\t\t\t1. Single Linked List \t\t\t 2. Circular Linked List")
	fmt.Println("\t\t\t\t\t\t3. Double Linked List \t\t\t 4. Circular Double Linked List")
	fmt.Println("\t\t\t\t\t\t6. Stack using arrays \t\t\t 6. Queue using arrays")
	fmt.Println("\t\t\t\t\t\t________________________________________________________________________")
	fmt.Scanf("%d", &mainMenuChoice)
}
