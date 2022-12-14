package main

import (
	"container/list"
	"fmt"
)

func main() {
	//new linkedlist
	queue := list.New()

	//append to queue
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	//dequeue
	front := queue.Front()
	fmt.Println(front.Value)

	//remove allocated memory
	queue.Remove(front)

}
