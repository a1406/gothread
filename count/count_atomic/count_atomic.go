package count_atomic

import (
	"fmt"
	"sync/atomic"
)

type Count_atomic struct {
	counter uint64
}

func (c *Count_atomic)Init_count(num int) {
	fmt.Println("init atomic count")
}

func (c *Count_atomic)Inc_count(i int) {
	atomic.AddUint64(&c.counter, 1)
}

func (c *Count_atomic)Read_count(i int) uint64 {
	if i == 0 {
		return c.counter
	}
	return 0
}


