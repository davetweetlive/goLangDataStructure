package main

import (
	"fmt"
)

func main() {
    inputMenu()
    choice := 1
	switch choice {
	case 1:
		fmt.Println("Single Linked List")
	case 2:
		fmt.Println("Circular Linked List")
	case 3:
		fmt.Println("Double Linked List")
	case 4:
		fmt.Println("Circular Double Linked List")
	case 5:
		fmt.Println("Stacks using Arrays")
	case 6:
		fmt.Println("Queue using Arrays")
	case 7:
		fmt.Println("Dequeue using arrays")
	default:
		fmt.Println("Please select option which is availabe!")
	}
}


func inputMenu()  {
    fmt.Println("")
    fmt.Println("\t\t\t\t\t\t1. Single Linked List \t\t\t 2. Circular Linked List")
    fmt.Println("\t\t\t\t\t\t3. Double Linked List \t\t\t 4. Circular Double Linked List")
    fmt.Println("\t\t\t\t\t\t6. Stack using arrays \t\t\t 6. Queue using arrays")
    fmt.Println("\t\t\t\t\t\t________________________________________________________________________")
}
