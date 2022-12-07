package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}

func (l *linkedList) append (n *node){
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 30}
	node4 := &node{data: 40}
	node5 := &node{data: 50}
	myList.append(node1)
	myList.append(node2)
	myList.append(node3)
	myList.append(node4)
	myList.append(node5)
	myList.printListData()
}