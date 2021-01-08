package arrays

import "fmt"

func RunBinarySearch() {
	fmt.Println("---Binary Search---")
	//input 1
	input1 := []int{10, 34, 56, 67, 89, 100, 1235, 10002}
	fmt.Printf("Searching for %v in %v. Output: %v\n", 1235, input1, binarySearch(input1, 1235))
	fmt.Printf("Searching for %v in %v. Output: %v\n", 345, input1, binarySearch(input1, 345))
}
