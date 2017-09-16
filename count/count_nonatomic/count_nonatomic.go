package count_nonatomic

import (
	"fmt"	
)

type Count_nonatomic struct {
	counter uint64
}

func (c *Count_nonatomic)Init_count(num int) {
	fmt.Println("init nonatomic count")
}

func (c *Count_nonatomic)Inc_count(i int) {
	c.counter++
}

func (c *Count_nonatomic)Read_count(i int) uint64 {
	if i == 0 {
		return c.counter
	}
	return 0
}


