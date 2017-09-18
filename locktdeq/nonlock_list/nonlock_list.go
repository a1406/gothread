package nonlock_list

import (
	"../node"
)

type List struct {
	Head *node.Node
	Tail *node.Node
}

func (l *List)InitList() {
	l.Head = nil
	l.Tail = nil
}

func (l *List)Push_head(n *node.Node) int {
	n.Pre = nil
	n.Next = l.Head
	if l.Head != nil {
		l.Head.Pre = n
	}
	if l.Tail == nil {
		l.Tail = n
	}
	
	l.Head = n
	return 0
}
func (l *List)Pop_head() *node.Node {
	if l.Head == nil {
		return nil
	}
	var ret *node.Node
	ret = l.Head
	l.Head = l.Head.Next

	if l.Head != nil {
		l.Head.Pre = nil
	} else {
		l.Tail = nil
	}
	return ret
}
func (l *List)Push_tail(node *node.Node) int {
	node.Next = nil
	node.Pre = l.Tail
	if l.Tail != nil {
		l.Tail.Next = node
	}
	if l.Head == nil {
		l.Head = node
	}
	
	l.Tail = node
	return 0
}
func (l *List)Pop_tail() *node.Node {
	if l.Tail == nil {
		return nil
	}
	var ret *node.Node
	ret = l.Tail
	l.Tail = l.Tail.Pre

	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
	return ret
}

