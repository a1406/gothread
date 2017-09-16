package count_atomic

import (
	"sync/atomic"
)

var counter uint64 = 0

func Inc_count() {
	atomic.AddUint64(&counter, 1)
//	counter++
}

func Read_count() uint64 {
	return counter
}


