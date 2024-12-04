package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SortArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

func Diff(x int, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	}
	return diff
}

func Distances(left []int, right []int) []int {
	var diffs []int
	for i := 0; i < len(left); i++ {
		diffs = append(diffs, Diff(left[i], right[i]))
	}
	return diffs
}

func Sum(diffs []int) int {
	var sum = 0
	for _, diff := range diffs {
		sum += diff
	}
	return sum
}

func SimilarityScore(left []int, right []int) []int {
	var scores []int
	for i := 0; i < len(left); i++ {
		count := 0
		for j := 0; j < len(right); j++ {
			if right[j] == left[i] {
				count++
			}
		}
		scores = append(scores, count*left[i])
	}
	return scores
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

	var list1 []int
	var list2 []int

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line) // Print each line
		entries := strings.Fields(line)
		leftVal, _ := strconv.Atoi(entries[0])
		rightVal, _ := strconv.Atoi(entries[1])
		list1 = append(list1, leftVal)
		list2 = append(list2, rightVal)
	}

	// Check for errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sortedList1 := SortArray(list1)
	sortedList2 := SortArray(list2)
	fmt.Printf("sl1 %d\n", sortedList1)
	fmt.Printf("sl2 %d\n", sortedList2)
	fmt.Printf("Distances %d\n", Distances(sortedList1, sortedList2))
	fmt.Printf("Sum %d\n", Sum(Distances(sortedList1, sortedList2)))
	fmt.Printf("Similarity Scores %d\n", SimilarityScore(sortedList1, sortedList2))
	fmt.Printf("Similarity Scores Sum %d\n", Sum(SimilarityScore(sortedList1, sortedList2)))
}
