package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type List[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (l *List[T]) insert(val T) {
	newNode := &Node[T]{value: val}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
		l.tail = newNode
	}
	l.length++
}

func (l *List[T]) print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d", current.value)
		current = current.next
	}
	fmt.Printf("\n")
}

func BuildList(line string) *List[int] {
	list := List[int]{}
	for _, val := range strings.Fields(line) {
		vInt, _ := strconv.Atoi(val)
		list.insert(vInt)
	}
	return &list
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when function exits

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	lists := []*List[int]{}
	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		list := BuildList(line)
		lists = append(lists, list)
	}

	// For each list
	var safeRows int = 0
	for i := 0; i < len(lists); i++ {
		current := lists[i].head
		lists[i].print()
		var diffSafe = true
		var isDecreasing = false
		var isIncreasing = false
		var noChange = false
		for current != nil {
			// if this is not head, check previous.
			// determine the change (isDesc/isAsc) depending on
			if current.next != nil {

				if current.value > current.next.value {
					isDecreasing = true
				}
				if current.value < current.next.value {
					isIncreasing = true
				}
				if current.value == current.next.value {
					noChange = true
				}

				if IsSafeDiff(Diff(current.value, current.next.value)) == false {
					diffSafe = false
				}
			}
			current = current.next
		}

		rowSafe := diffSafe && !noChange && (isDecreasing != isIncreasing)
		if rowSafe {
			safeRows++
		}
	}

	fmt.Printf("There are %d safe rows", safeRows)
	// Check for errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}

func IsSafeDiff(x int) bool {
	if x <= 3 {
		return true
	}
	return false
}

func Diff(x int, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	}
	return diff
}
