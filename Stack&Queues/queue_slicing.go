package main

import "fmt"

func enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []int) (int, []int) {
	element := queue[0]
	if len(queue) == 1 {
		var temp = []int{}
		return element, temp
	}
	return element, queue[1:]
}

func main() {
	var queue []int

	queue = enqueue(queue, 10)
	queue = enqueue(queue, 20)
	queue = enqueue(queue, 30)

	fmt.Println("Queue: ", queue)

	_, queue = dequeue(queue)
	fmt.Println("Queue: ", queue)

	_, queue = dequeue(queue)

	queue = enqueue(queue, 40)
	fmt.Println("Queue: ", queue)
}
