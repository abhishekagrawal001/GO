package main
import "fmt"

func search(array[]int, k int) int {
	for i:=0; i<len(array); i++ {
		if array[i] == k {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int {5,2,7,4,9}
	k := 4
	result := search(arr,k)
	if result != -1{
		fmt.Printf("Element %v was found at index %v",k,result)
	}else {
		fmt.Printf("Element was not found")
	}
}