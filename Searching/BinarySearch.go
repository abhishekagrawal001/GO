package main

import "fmt"

func search(array []int, k int) int {
	low, high := 0, len(array)-1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] == k {
			return mid
		} else if array[mid] > k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 1
	result := search(arr, k)
	if result != -1 {
		fmt.Printf("Element %v was found at index %v", k, result)
	} else {
		fmt.Printf("Element was not found")
	}
}
