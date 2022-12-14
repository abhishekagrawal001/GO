package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type linkedList struct {
	head *Node
}

func (l *linkedList) insert(d int) {
	list := &Node{data: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}


func main() {
	sl := linkedList{}
	sl.insert(1)
	sl.insert(2)
	sl.insert(3)
	sl.insert(2)
	sl.insert(1)
	// show(&sl)
}
