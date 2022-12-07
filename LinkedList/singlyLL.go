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

func show(l *linkedList) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.data)
		p = p.next
	}
}

func main() {
	sl := linkedList{}
	sl.insert(10)
	sl.insert(20)
	sl.insert(30)
	sl.insert(40)
	sl.insert(50)
	show(&sl)
}
