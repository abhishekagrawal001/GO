package main

import "fmt"

func Insertionsort(arr [6]int) {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		//shifting elements to make space for key element
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	fmt.Println(arr)
}

func main() {
	arr := [6]int{4, 5, 7, 1, 3, 0}
	Insertionsort(arr)
}
