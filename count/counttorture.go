package main

import (
	"fmt"
	"flag"
	"time"
	"./count_nonatomic"
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

func count_read_perf_test(i int) {
	var j uint64 = 0
	var n_reads_local uint64 = 0

	for ; goflag == 0; {
		for i := uint64(0); i < COUNT_READ_RUN; i++ {
			j = j + count_nonatomic.Read_count()
		}
		n_reads_local += COUNT_READ_RUN;
	}
	thread_param[i] += n_reads_local
	threaddone[i] <- true
}

func count_update_perf_test(i int) {
	var n_update_local uint64 = 0

	for ; goflag == 0; {
		for i := uint64(0); i < COUNT_UPDATE_RUN; i++ {
			count_nonatomic.Inc_count()
		}
		n_update_local += COUNT_READ_RUN;
	}
	thread_param[i] += n_update_local
	threaddone[i] <- true	
}

func perftest(nreader, nupdater int) {
	for i := 0; i < nreader + nupdater; i++ {
		threaddone[i] = make(chan bool)		
	}
//	fmt.Printf("n_read = %d, n_update = %d, duration = %d\n", nreader, nupdater, duration)
	if nreader + nupdater > MAX_THREAD {
		fmt.Println("too much thread")
		return
	}
	for i := 0; i < nreader; i++ {
		go count_read_perf_test(i)
	}
	for i := 0; i < nupdater; i++ {
		go count_update_perf_test(i + nreader)
	}
	perftestrun(nreader, nupdater)
}

func perftestrun(nreader, nupdater int) {
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
//		float64(duration) * 1000000.0 * float64(nreader) / n_reads,
//		duration * 1000000 * nupdater / n_updates)
}

func main() {
	n_read := flag.Int("r", 1, "num of read thread")
	n_update := flag.Int("u", 1, "num of update thread")
	n_duration := flag.Int64("s", 240, "sleep time")
	flag.Parse()
	duration = *n_duration

	perftest(*n_read, *n_update)
}
