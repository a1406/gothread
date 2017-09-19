package main

import (
	"fmt"
	"flag"
	"time"
	"./node"
	"./nonlock_list"
)

const MAX_THREAD int = 200
const COUNT_THREAD_RUN int = 1000

type thread_param_struct struct {
	pushnum uint64
	popnum uint64
}

var threaddone [MAX_THREAD]chan bool
var thread_param [MAX_THREAD]thread_param_struct
var goflag int = 0
var duration int64

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

func pushhead_func(l list_int, chan_index int) {
	for ; goflag == 0; {
	}
	threaddone[chan_index] <- true		
}
func pushtail_func(l list_int, chan_index int) {
	for ; goflag == 0; {
	}
	threaddone[chan_index] <- true		
}
func pophead_func(l list_int, chan_index int) {
	for ; goflag == 0; {
	}
	threaddone[chan_index] <- true			
}
func poptail_func(l list_int, chan_index int) {
	for ; goflag == 0; {
	}
	threaddone[chan_index] <- true			
}

func perftestrun(l list_int, nthread int) {
	fmt.Println(time.Now())	
	time.Sleep(time.Duration(duration) * time.Millisecond)
	goflag = 1

	fmt.Println(time.Now())

	var n_push, n_pop uint64
	for i := 0; i < nthread; i++ {
		<- threaddone[i]
		n_push += thread_param[i].pushnum
		n_pop += thread_param[i].popnum
	}
	
	fmt.Printf("n_push: %d n_pop: %d nthread: %d duration: %d\n",
		n_push, n_pop, nthread, duration)
	var tr float64 = float64(duration) * 1000000 / float64(n_push)
	var tu float64 = float64(duration) * 1000000 / float64(n_pop)	
	fmt.Printf("ns/push: %f  ns/pop: %f\n", tr, tu)

	var l_len uint64
	for {
		node := l.Pop_head()
		if node == nil {
			break
		}
		l_len++
	}

	if l_len + n_pop != n_push {
		fmt.Printf("err, l_len = %d, n_pop = %d, n_push = %d\n", l_len, n_pop, n_push)
	}

// 	var final_count uint64
// 	for i := 0; i < nreader + nupdater; i++ {
// 		final_count += (count_int).Read_count(i)
// 	}
// 	
// 	fmt.Printf("read count = %d[%f%%]\n", final_count,
// 		float64(final_count) / float64(n_updates) * 100)
}

func test_pushpop(l list_int, pushhead int, pushtail int, pophead int, poptail int) {

	for i := 0; i < pushhead + pushtail + pophead + poptail; i++ {
		threaddone[i] = make(chan bool)		
	}

	var chan_index int
	for i := 0; i < pushhead; i++ {
		go pushhead_func(l, chan_index)
		chan_index++
	}
	for i := 0; i < pushtail; i++ {
		go pushtail_func(l, chan_index)
		chan_index++		
	}
	for i := 0; i < pophead; i++ {
		go pophead_func(l, chan_index)
		chan_index++		
	}
	for i := 0; i < poptail; i++ {
		go poptail_func(l, chan_index)
		chan_index++		
	}
	perftestrun(l, pushhead + pushtail + pophead + poptail)
}


func main() {
//	n_read := flag.Int("r", 1, "num of read thread")
//	n_update := flag.Int("u", 1, "num of update thread")
//	n_duration := flag.Int64("s", 240, "sleep time")
	var runtype int
	flag.IntVar(&runtype, "t", 0, "do atomic")
	flag.Int64Var(&duration, "s", 240, "sleep time")	

	var pushhead, pushtail, pophead, poptail int
	flag.IntVar(&pushhead, "pushhead", 1, "")
	flag.IntVar(&pushtail, "pushtail", 1, "")
	flag.IntVar(&pophead, "pophead", 1, "")
	flag.IntVar(&poptail, "poptail", 1, "")	
	
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
	test_pushpop(list, pushhead, pushtail, pophead, poptail)
//	perftest(*n_read, *n_update, count_int)
}
