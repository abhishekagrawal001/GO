package main

import "fmt"

func SelectionSort(arr [6]int) {

	for i := 0; i < len(arr)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		{
			temp := arr[minIndex]
			arr[minIndex] = arr[i]
			arr[i] = temp
		}
	}
	fmt.Println(arr)
}

func main() {
	arr := [6]int{4, 5, 7, 1, 3, 0}
	SelectionSort(arr)
}
