package lockh_list

import (
       "../node"
       "sync"
)

const MAX_HASH_SIZE int = 8

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

func (l *Hashlist)Init_list() {
	l.head_locker = new(sync.Mutex)
	l.tail_locker = new(sync.Mutex)
	for i := 0; i < MAX_HASH_SIZE; i++ {
		l.hash_data[i].locker = new(sync.Mutex)
	}
}

func  (l *Hashlist)Push_head(node *node.Node) int {
	return 0
}
func  (l *Hashlist)Pop_head() *node.Node {
	return nil
}
func  (l *Hashlist)Push_tail(node *node.Node) int {
	return 0	
}
func  (l *Hashlist)Pop_tail() *node.Node {
	return nil
}



//static int moveleft(int idx)
//{
// 	return (idx - 1) & (PDEQ_N_BKTS - 1);
//}
// 
//static int moveright(int idx)
//{
// 	return (idx + 1) & (PDEQ_N_BKTS - 1);
//}
