package stack

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Global variables for the package stack
var (
	stackMenuOption int
	invalidChoice   bool
)

// Item does somethings which is unknow for now
type Item generic.Type

type itemStack struct {
	items []Item
	lock  sync.RWMutex
}

// Clear screen function clears the terminal screen, currently working only for linux operating systems
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// StackMain ...
func StackMain() {

	clearScreen()
	stackMenu()

	for stackMenuOption != 0 {
		if !invalidChoice {
			switch stackMenuOption {
			case 1:
				fmt.Println("Enter a value")
			case 2:
				fmt.Println("Push Value !")
			case 3:
				fmt.Println("Pop Value!")
			case 4:
				fmt.Println("Length ")
			case 5:
				fmt.Println("Display!")
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

// CreateStack Functions written below are self explainatory by their names

// New creates a new itemStack
func (s *itemStack) New() *itemStack {
	s.items = []Item{}
	return s
}

// Push adds an Item to the top of the stack
func (s *itemStack) Push(t Item) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Pop removes an Item from the top of the stack
func (s *itemStack) Pop() *Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item
}
