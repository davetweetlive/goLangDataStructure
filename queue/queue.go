package queue

import (
	"fmt"
	"os"
	"os/exec"
)

// Global variables for the package stack
var (
	queueMenuOption int
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

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

// QueueMain ...
func QueueMain() {
	clearScreen()
	QueueMenu()
	newQueue := NewQueue()
	for queueMenuOption != 0 {
		if !invalidChoice {
			switch queueMenuOption {
			case 1:

				fmt.Println("Stack has been succesfully created, please enter value!")
			case 2:
				fmt.Println("Enter a number\t")
				// var num int
				// _, err := fmt.Scanf("%d", &num)
				// if err != nil {
				// 	fmt.Println("Can't take user input", err)
				// }
				// newStack.Push(&Node{num})
			case 3:
				// val := newStack.Pop()
				// fmt.Println("Popped Value is:", val)
			case 4:
				fmt.Println("Length ")
			case 5:
				fmt.Println("Display!")
				// fmt.Println((*newStack).nodes)
			default:
				fmt.Println("Invalid choice!")
			}
		} else {
			fmt.Println("Not a valid choice!")
			return
		}
		QueueMenu()
	}

}

func QueueMenu() {
	fmt.Println("")
	fmt.Println("\t\t\t\t\t\t1. Create stack        \t\t\t 2. Push value")
	fmt.Println("\t\t\t\t\t\t3. Pop value           \t\t\t 4. Length of stack")
	fmt.Println("\t\t\t\t\t\t5. Display stack ")
	fmt.Println("\t\t\t\t\t\t0. Exit from the menu")
	fmt.Println("\t\t\t\t\t\t________________________________________________________________________")
	_, err := fmt.Scanf("%d", &queueMenuOption)

	if err != nil {
		invalidChoice = true
		return
	}
}