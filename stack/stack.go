package stack

import (
	"fmt"
	"os"
	"os/exec"
	// "sync"
	// "github.com/cheekybits/genny/generic"
)

// Global variables for the package stack
var (
	stackMenuOption int
	invalidChoice   bool
)

// Clear screen function clears the terminal screen, currently working only for linux operating systems
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type Node struct {
	Value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// type newStack *Stack

// StackMain ...
func StackMain() {
	clearScreen()
	stackMenu()
	newStack := NewStack()
	for stackMenuOption != 0 {
		if !invalidChoice {
			switch stackMenuOption {
			case 1:

				fmt.Println("Stack has been succesfully created, please enter value!")
			case 2:
				fmt.Println("Enter an integer value!\t")
				var newElement int
				fmt.Scan(&newElement)
				newStack.Push(&Node{newElement})
			case 3:
				fmt.Println("Pop Value!")
				newStack.Pop()
			case 4:
				fmt.Println("Length ")
			case 5:
				fmt.Println("Display!")
				fmt.Println(*(&newStack.nodes))
			default:
				fmt.Println("Invalid choice!")
			}
		} else {
			fmt.Println("Not a valid choice!")
			return
		}
		stackMenu()
	}

}

// stackMenu provides menu feature for stacks
func stackMenu() {
	fmt.Println("")
	fmt.Println("\t\t\t\t\t\t1. Create stack        \t\t\t 2. Push value")
	fmt.Println("\t\t\t\t\t\t3. Pop value           \t\t\t 4. Length of stack")
	fmt.Println("\t\t\t\t\t\t5. Display stack ")
	fmt.Println("\t\t\t\t\t\t0. Exit from the menu")
	fmt.Println("\t\t\t\t\t\t________________________________________________________________________")
	_, err := fmt.Scanf("%d", &stackMenuOption)

	if err != nil {
		invalidChoice = true
		return
	}
}
