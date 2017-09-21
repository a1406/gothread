package main

import (
	"fmt"
	"flag"
	"time"
	"./nonlock_table"
	"./rcu_table"
	"./lock_table"		
)

type hash_table_int interface {
	Init()
	Lookup(k int) bool
	Insert(k int)
	Delete(k int) bool
	Num() uint
}

const MAX_THREAD uint = 20
var goflag int = 0
var read_threaddone [MAX_THREAD]chan bool
var insert_threaddone [MAX_THREAD]chan bool
var delete_threaddone [MAX_THREAD]chan bool
var read [MAX_THREAD] uint
var insert_num [MAX_THREAD] uint
var delete_num [MAX_THREAD] uint
var delete_suc_num [MAX_THREAD] uint

func read_perf_test(i uint, table_int hash_table_int) {
	for ; goflag == 0; {
		for t := -1000; t <= 1000; t++ {
			table_int.Lookup(t)
			read[i]++
		}
	}
	read_threaddone[i] <- true
}
func insert_perf_test(i uint, table_int hash_table_int) {
	for ; goflag == 0; {
		for t := -1000; t <= 1000; t++ {
			table_int.Insert(t)
			insert_num[i]++
		}
	}
	insert_threaddone[i] <- true	
}
func delete_perf_test(i uint, table_int hash_table_int) {
	for ; goflag == 0; {
		for t := -1000; t <= 1000; t++ {
			if table_int.Delete(t) {
				delete_suc_num[i]++
			}
			delete_num[i]++			
		}
	}
	delete_threaddone[i] <- true	
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
		insert_threaddone[i] = make(chan bool)		
	}
	for i := uint(0); i < writer / 2; i++ {
		delete_threaddone[i] = make(chan bool)		
	}
	
	table_int.Init()
	for i := uint(0); i < reader; i++ {
		go read_perf_test(i, table_int)
	}
	for i := uint(0); i < writer; i++ {
		go insert_perf_test(i, table_int)
	}
	for i := uint(0); i < writer / 2; i++ {
		go delete_perf_test(i, table_int)
	}

	start_time := time.Now().UnixNano()
//	fmt.Println(start_time)	
	time.Sleep(time.Duration(duration) * time.Millisecond)
	goflag = 1

	for i := uint(0); i < reader; i++ {
		<- read_threaddone[i]
	}
	for i := uint(0); i < writer; i++ {
		<- insert_threaddone[i]		
	}
	for i := uint(0); i < writer / 2; i++ {
		<- delete_threaddone[i]		
	}

	end_time := time.Now().UnixNano()	
	//	fmt.Println(end_time)
	escape_time := end_time - start_time
	fmt.Println(escape_time)

	var n_reads, n_writes1, n_writes2 , check_num uint
	for i := uint(0); i < MAX_THREAD; i++ {
		n_reads += read[i]
		n_writes1 += insert_num[i]
		n_writes1 += delete_suc_num[i]
		n_writes2 += insert_num[i]
		n_writes2 += delete_num[i]

		check_num += insert_num[i]
		check_num -= delete_suc_num[i]

		// if insert_num[i] > 0 {
		// 	fmt.Printf("insert %d\n", insert_num[i])
		// }
		// if delete_suc_num[i] > 0 {
		// 	fmt.Printf("delete %d\n", delete_suc_num[i])
		// }
		
	}
	
	fmt.Printf("n_reads: %d n_writes: [%d]%d nreaders: %d nwriters: [%d]%d duration: %d\n",
		n_reads, n_writes1, n_writes2, reader, writer / 2, writer, escape_time)
	var tr float64 = float64(escape_time) * float64(reader) / float64(n_reads)
	var tu float64 = float64(escape_time) * float64(writer + writer / 2) / float64(n_writes2)
	fmt.Printf("ns/read: %f  ns/update: %f\n", tr, tu)

	final_count := table_int.Num()
	fmt.Printf("final count = %d, check num = %d\n", final_count, check_num)
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
 	case 1:
 		table_int = new(lock_table.Table_lock)
 		break
 	case 2:
		table_int = new(rcu_table.Table_rcu)
		break
 	default:
 		table_int = new(nonlock_table.Table_nonlock)
 		break
 	}

	perftest(reader, writer, duration, table_int)
}
