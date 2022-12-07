package main

import "fmt"

func main() {
	arr := [5]int{5,6,3,7,1}

	//Bubble sort in ascending order
	for i:=0 ; i<len(arr)-1; i++{
		for j:=0 ; j<len(arr)-i-1; j++ {
			if arr[j]>arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)

	//Bubble sort in descending order
	for i:=0; i<len(arr); i++ {
		for j:=len(arr)-1; j>=i+1; j-- {
			if arr[j]>arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
	fmt.Println(arr)
}