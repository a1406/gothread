package lock_list

import (
	"../node"
	"sync/mutex"
)

type Locklist stuct {
	head list
	tail list
}

type list struct {
	locker *sync.Mutex
	head *node.Node
	tail *node.Node
}

func (l *Locklist)InitList() {
	l.head.locker = new(sync.Mutex)
	l.tail.locker = new(sync.Mutex)
}

func (l *Locklist)Push_head(n *node.Node) int {
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
func (l *Locklist)Pop_head() *node.Node {
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
func (l *Locklist)Push_tail(node *node.Node) int {
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
func (l *Locklist)Pop_tail() *node.Node {
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

