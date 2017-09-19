package lockh_list

import (
       "../node"
       "sync"
)

const MAX_HASH_SIZE int = 8
const MAX_HASH_MASK int = MAX_HASH_SIZE - 1

type Hashlist struct {
	head_idx int
	head_locker *sync.Mutex	
	tail_idx int
	tail_locker *sync.Mutex		
	hash_data [MAX_HASH_SIZE]list
}

type list struct {
	locker *sync.Mutex
	head *node.Node
	tail *node.Node
}

func (l *list)empty() bool {
	return l.head == nil
}

func (l *list)pop_head() *node.Node {
	var ret *node.Node
	l.locker.Lock()
	if l.empty() {
		l.locker.Unlock()		
		return nil
	}
	
	ret = l.head
	l.head = l.head.Next

	if l.head != nil {
		l.head.Pre = nil
	} else {
		l.tail = nil
	}
	l.locker.Unlock()
	return ret
}
func (l *list)push_head(n *node.Node) {
	l.locker.Lock()

	n.Pre = nil
	n.Next = l.head
	if l.head != nil {
		l.head.Pre = n
	} else {
		l.tail = n
	}
	
	l.head = n
	l.locker.Unlock()
	return
}
func (l *list)pop_tail() *node.Node {
	var ret *node.Node
	l.locker.Lock()	
	ret = l.tail
	l.tail = l.tail.Pre

	if l.tail != nil {
		l.tail.Next = nil
	} else {
		l.head = nil
	}
	l.locker.Unlock()	
	return ret
}
func (l *list)push_tail(n *node.Node) {
	l.locker.Lock()

	n.Next = nil
	n.Pre = l.tail
	if l.tail != nil {
		l.tail.Next = n
	} else {
		l.head = n
	}
	
	l.tail = n
	l.locker.Unlock()
	return
}

func (l *Hashlist)Init_list() {
	l.head_locker = new(sync.Mutex)
	l.tail_locker = new(sync.Mutex)
	for i := 0; i < MAX_HASH_SIZE; i++ {
		l.hash_data[i].locker = new(sync.Mutex)
	}
}

func (l *Hashlist)Push_head(node *node.Node) int {
	l.head_locker.Lock()
	
	l.move_head()
	l.head_locker.Unlock()	
	return 0
}
func (l *Hashlist)Pop_head() *node.Node {
	return nil
}
func (l *Hashlist)Push_tail(node *node.Node) int {
	return 0	
}
func (l *Hashlist)Pop_tail() *node.Node {
	return nil
}

func (l *Hashlist)move_head() {
	l.head_idx = l.head_idx - 1 & MAX_HASH_MASK
}
func (l *Hashlist)move_tail() {
	l.tail_idx = l.tail_idx + 1 & MAX_HASH_MASK
}

