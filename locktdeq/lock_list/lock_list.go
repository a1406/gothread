package lock_list

import (
	"../node"
	"sync"
)

type Locklist struct {
	head list
	tail list
}

type list struct {
	locker *sync.Mutex
	head *node.Node
	tail *node.Node
}

func (l *list)pop_head() *node.Node{
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
func (l *list)pop_tail() *node.Node{
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


func (l *Locklist)Init_list() {
	l.head.locker = new(sync.Mutex)
	l.tail.locker = new(sync.Mutex)
}

func (l *Locklist)Push_head(n *node.Node) int {
	l.head.locker.Lock()
	
	n.Pre = nil
	n.Next = l.head.head
	if l.head.head != nil {
		l.head.head.Pre = n
	}
	if l.head.tail == nil {
		l.head.tail = n
	}
	
	l.head.head = n
	l.head.locker.Unlock()
	return 0
}

func (l *Locklist)tail_to_head() *node.Node {
	l.tail.locker.Lock()

	if l.tail.head == nil {
		l.tail.locker.Unlock()
		l.head.locker.Unlock()	
		return nil
	}

	var ret *node.Node
	ret = l.tail.pop_head()
	
	l.tail.locker.Unlock()
	l.head.locker.Unlock()	
	return ret
}
func (l *Locklist)head_to_tail() *node.Node {
	l.tail.locker.Unlock()

	l.head.locker.Lock()	
	l.tail.locker.Lock()

	if l.tail.tail != nil {
		var ret *node.Node
		ret = l.tail.pop_tail()
		l.tail.locker.Unlock()
		l.head.locker.Unlock()
		return ret
	}

	if l.head.tail == nil {
		l.tail.locker.Unlock()
		l.head.locker.Unlock()	
		return nil
	}

	var ret *node.Node
	ret = l.head.pop_tail()
	l.tail.locker.Unlock()
	l.head.locker.Unlock()
	return ret
}

func (l *Locklist)Pop_head() *node.Node {
	l.head.locker.Lock()
	
	if l.head.head == nil {
		return l.tail_to_head()
	}
	var ret *node.Node
	ret = l.head.pop_head()
	l.head.locker.Unlock()	
	return ret
}
func (l *Locklist)Push_tail(node *node.Node) int {
	l.tail.locker.Lock()			
	node.Next = nil
	node.Pre = l.tail.tail
	if l.tail.tail != nil {
		l.tail.tail.Next = node
	}
	if l.tail.head == nil {
		l.tail.head = node
	}
	
	l.tail.tail = node
	l.tail.locker.Unlock()		
	return 0
}
func (l *Locklist)Pop_tail() *node.Node {
	l.tail.locker.Lock()				
	if l.tail.tail == nil {
		return l.head_to_tail()
	}
	var ret *node.Node
	ret = l.tail.pop_tail()
	l.tail.locker.Unlock()			
	return ret
}

