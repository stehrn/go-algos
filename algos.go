package main

import "fmt"

func main() {
	fmt.Println("Running binary search...")
}

// BinarySearch on sorted input array 'A', searching for value 's'
// will return array index of found value, -1 otherwise
func BinarySearch(A []int, s int) int {
	start := 0
	end := len(A) - 1
	for start <= end {
		mid := start + (end-start)/2
		value := A[mid]
		if value == s {
			return mid
		}
		if value < s {
			start = mid + 1
		} else if value > s {
			end = mid - 1
		}
	}
	return -1
}
