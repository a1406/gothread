package main

import (
	"fmt"
	"flag"
	//	"time"
	"./node"
	"./nonlock_list"
)

type list_int interface {
	Push_head(node *node.Node) int
	Pop_head() *node.Node
	Push_tail(node *node.Node) int
	Pop_tail() *node.Node
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func test_normal(l list_int) {
	var tnode *node.Node
	
	for i := 0; i < 1000; i++ {
		tnode = new(node.Node)
		tnode.Data = i
		l.Push_head(tnode)
	}
	for i := 0; i < 1000; i++ {
		tnode = l.Pop_tail()
		if tnode.Data != i {
			err := fmt.Sprintf("err pop tail, i = %d", i)
			panic(err)
		}
	}
	tnode = l.Pop_head()
	if tnode != nil {
		panic("err pop head, not nil")		
	}
	tnode = l.Pop_tail()
	if tnode != nil {
		panic("err pop tail, not nil")		
	}

	for i := 0; i < 1000; i++ {
		tnode = new(node.Node)
		tnode.Data = i
		l.Push_tail(tnode)
	}
	for i := 0; i < 1000; i++ {
		tnode = l.Pop_head()
		if tnode.Data != i {
			err := fmt.Sprintf("err pop head, i = %d", i)
			panic(err)
		}
	}
	tnode = l.Pop_head()
	if tnode != nil {
		panic("err pop head, not nil")		
	}
	tnode = l.Pop_tail()
	if tnode != nil {
		panic("err pop tail, not nil")		
	}
}


func main() {
//	n_read := flag.Int("r", 1, "num of read thread")
//	n_update := flag.Int("u", 1, "num of update thread")
//	n_duration := flag.Int64("s", 240, "sleep time")
	var runtype int
	flag.IntVar(&runtype, "t", 0, "do atomic")
	flag.Parse()
	//	duration = *n_duration

	var list list_int
	
	switch runtype {
	case 1:
		list = new(nonlock_list.List)
		break
//	case 2:
//		count_int = new(count_stat.Count_stat)
//		break
	default:
		list = new(nonlock_list.List)
		break
	}

	test_normal(list)
//	perftest(*n_read, *n_update, count_int)
}
