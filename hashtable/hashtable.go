package main

import (
	"fmt"
	"flag"
	"time"
	"./nonlock_table"
//	"./rcu_table"
//	"./lock_table"		
)

type hash_table_int interface {
	Init()
	Lookup(k int) bool
	Insert(k int)
	Delete(k int)
}

const MAX_THREAD uint = 20
var goflag int = 0
var read_threaddone [MAX_THREAD]chan bool
var write_threaddone [MAX_THREAD]chan bool

func read_perf_test(i uint, table_int hash_table_int) {
	for ; goflag == 0; {
	}
	read_threaddone[i] <- true
}
func write_perf_test(i uint, table_int hash_table_int) {
	for ; goflag == 0; {
	}
	write_threaddone[i] <- true	
}

func perftest(reader, writer, duration uint, table_int hash_table_int) {
	if reader > MAX_THREAD || writer > MAX_THREAD {
		fmt.Println("too much thread")
		return
	}

	for i := uint(0); i < reader; i++ {
		read_threaddone[i] = make(chan bool)		
	}
	for i := uint(0); i < writer; i++ {
		write_threaddone[i] = make(chan bool)		
	}
	
	table_int.Init()
	for i := uint(0); i < reader; i++ {
		go read_perf_test(i, table_int)
	}
	for i := uint(0); i < writer; i++ {
		go write_perf_test(i, table_int)
	}

	fmt.Println(time.Now())	
	time.Sleep(time.Duration(duration) * time.Millisecond)
	goflag = 1
	fmt.Println(time.Now())

	for i := uint(0); i < reader; i++ {
		<- read_threaddone[i]
	}
	for i := uint(0); i < writer; i++ {
		<- write_threaddone[i]		
	}
}

func main() {
	var reader, writer, duration uint
	var runtype int
	flag.IntVar(&runtype, "t", 0, "runtype")
	flag.UintVar(&reader, "r", 2, "read thread num")
	flag.UintVar(&writer, "w", 2, "write thread num")
	flag.UintVar(&duration, "s", 240, "sleep time")	
	flag.Parse()

	var table_int hash_table_int
 	switch runtype {
// 	case 1:
// 		table_int = new(count_atomic.Count_atomic)
// 		break
// 	case 2:
// 		table_int = new(count_stat.Count_stat)
// 		break
 	default:
 		table_int = new(nonlock_table.Table_nonlock)
 		break
 	}

	perftest(reader, writer, duration, table_int)
}
