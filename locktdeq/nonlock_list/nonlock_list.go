package nonlock_list

import (
	"../node"
)

type List struct {
	head *node.Node
	tail *node.Node
}

func (l *List)Init_list() {
}

func (l *List)Push_head(n *node.Node) int {
	n.Pre = nil
	n.Next = l.head
	if l.head != nil {
		l.head.Pre = n
	}
	if l.tail == nil {
		l.tail = n
	}
	
	l.head = n
	return 0
}
func (l *List)Pop_head() *node.Node {
	if l.head == nil {
		return nil
	}
	var ret *node.Node
	ret = l.head
	l.head = l.head.Next

	if l.head != nil {
		l.head.Pre = nil
	} else {
		l.tail = nil
	}
	return ret
}
func (l *List)Push_tail(node *node.Node) int {
	node.Next = nil
	node.Pre = l.tail
	if l.tail != nil {
		l.tail.Next = node
	}
	if l.head == nil {
		l.head = node
	}
	
	l.tail = node
	return 0
}
func (l *List)Pop_tail() *node.Node {
	if l.tail == nil {
		return nil
	}
	var ret *node.Node
	ret = l.tail
	l.tail = l.tail.Pre

	if l.tail != nil {
		l.tail.Next = nil
	} else {
		l.head = nil
	}
	return ret
}

func (l *List)Debug_print() {
}
