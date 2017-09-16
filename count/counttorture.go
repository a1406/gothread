package main

import (
	"fmt"
	"flag"
	"time"
	"./count_nonatomic"
	"./count_atomic"
	"./count_stat"		
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const MAX_THREAD int = 200
const COUNT_READ_RUN uint64 = 1000
const COUNT_UPDATE_RUN uint64 = 1000

var threaddone [MAX_THREAD]chan bool
var thread_param [MAX_THREAD]uint64
var goflag int = 0
var duration int64

type count_int interface {
	Init_count(num int)
	Inc_count(i int)
	Read_count(i int) uint64
}

func count_read_perf_test(i int, count_int count_int) {
	var n_reads_local uint64 = 0

	for ; goflag == 0; {
		for j := uint64(0); j < COUNT_READ_RUN; j++ {
//			fmt.Printf("i = %d, countint = %p\n", i, count_int)
			(count_int).Read_count(0)
		}
		n_reads_local += COUNT_READ_RUN;
	}
	thread_param[i] += n_reads_local
	threaddone[i] <- true
}

func count_update_perf_test(i int, count_int count_int) {
	var n_update_local uint64 = 0

	for ; goflag == 0; {
		for j := uint64(0); j < COUNT_UPDATE_RUN; j++ {
			(count_int).Inc_count(i)
		}
		n_update_local += COUNT_READ_RUN;
	}
	thread_param[i] += n_update_local
	threaddone[i] <- true	
}

func perftest(nreader, nupdater int, count_int count_int) {
	if nreader + nupdater > MAX_THREAD {
		fmt.Println("too much thread")
		return
	}
	
	for i := 0; i < nreader + nupdater; i++ {
		threaddone[i] = make(chan bool)		
	}

	(count_int).Init_count(nreader + nupdater)

	for i := 0; i < nreader; i++ {
		go count_read_perf_test(i, count_int)
	}
	for i := 0; i < nupdater; i++ {
		go count_update_perf_test(i + nreader, count_int)
	}
	perftestrun(nreader, nupdater, count_int)
}

func perftestrun(nreader, nupdater int, count_int count_int) {
	fmt.Println(time.Now())	
	time.Sleep(time.Duration(duration) * time.Millisecond)
	goflag = 1

	fmt.Println(time.Now())

	var n_reads, n_updates uint64
	for i := 0; i < nreader; i++ {
		<- threaddone[i]
		n_reads += thread_param[i]
	}
	for i := 0; i < nupdater; i++ {
		<- threaddone[i + nreader]		
		n_updates += thread_param[i + nreader]
	}
	
	fmt.Printf("n_reads: %d n_updates: %d nreaders: %d nupdaters: %d duration: %d\n",
		n_reads, n_updates, nreader, nupdater, duration)
	var tr float64 = float64(duration) * 1000000 * float64(nreader) / float64(n_reads)
	var tu float64 = float64(duration) * 1000000 * float64(nupdater) / float64(n_updates)	
	fmt.Printf("ns/read: %f  ns/update: %f\n", tr, tu)

	var final_count uint64
	for i := 0; i < nreader + nupdater; i++ {
		final_count += (count_int).Read_count(i)
	}
	
	fmt.Printf("read count = %d[%f%%]\n", final_count,
		float64(final_count) / float64(n_updates) * 100)
}

func main() {
	n_read := flag.Int("r", 1, "num of read thread")
	n_update := flag.Int("u", 1, "num of update thread")
	n_duration := flag.Int64("s", 240, "sleep time")
	var runtype int
	flag.IntVar(&runtype, "t", 0, "do atomic")
	flag.Parse()
	duration = *n_duration

	var count_int count_int
	
	switch runtype {
	case 1:
		count_int = new(count_atomic.Count_atomic)
		break
	case 2:
		count_int = new(count_stat.Count_stat)
		break
	default:
		count_int = new(count_nonatomic.Count_nonatomic)
		break
	}

	perftest(*n_read, *n_update, count_int)
}
