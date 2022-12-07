package main
import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type List struct{
	head *Node
	tail *Node
}

func (l *List)insert(d int) {
	list := &Node{data: d, prev: nil, next: nil}
	if l.head != nil {
		l.head = list
		l.tail = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		list.prev = p
		p.next = list
		l.tail = list
	}
}

func show(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.data)
		p = p.next
	}
	fmt.Println()
}


func main() {
	dl := List{}
	dl.insert(10)
	dl.insert(20)
	dl.insert(30)
	show(&dl)
}