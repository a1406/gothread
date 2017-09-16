package count_stat

import (
	"fmt"	
)

type Count_stat struct {
	counter []uint64	
}



func (c *Count_stat)Init_count(num int) {
	c.counter = make([]uint64, num)
	fmt.Printf("init stat count, num = %d, len = %d, c = %p, c.counter = %p\n",
		num, len(c.counter), &c, &c.counter)	
}

func (c *Count_stat)Inc_count(i int) {
// 	if i >= len(c.counter) {
// 		fmt.Printf("inc: i = %d, len = %d, c = %p, c.count = %p\n", i, len(c.counter), &c, &c.counter)
// 		return
// 	}
	c.counter[i]++
}

func (c *Count_stat)Read_count(i int) uint64 {
// 	if i >= len(c.counter) {
// 		fmt.Printf("read: i = %d, len = %d, c = %p, c.count = %p\n", i, len(c.counter), &c, &c.counter)
// 		return 0
// 	}
	return c.counter[i]
}

