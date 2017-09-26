package main

import (
	"fmt"
	"flag"
	"math"
)

//n=3: c=Sqrt(1 - 0.5 * 0.5) h=0.5
//n=4: a=c1 b=1-h1 c=sqrt(a*a + b*b) h=sqrt(b*b+c*c*0.5*0.5) s=0.5*a*b*3
func count_area(n uint) float64 {
//	if n_line == 3 {
//		return math.Sqrt(1 - 0.5 * 0.5) * 0.5 * 3
	//	}
	c1 := math.Sqrt(1 - 0.5 * 0.5)
	h1 := 0.5
//	var a, b float64
	var s float64 = math.Sqrt(1 - 0.5 * 0.5) * 0.5 * 3
	for i := uint(1); i < n; i++ {
		a := c1
		b := 1 - h1
		c1 = math.Sqrt(0.25*a*a + b*b)
		h1 = math.Sqrt(b*b + c1*c1*0.25)
		s = s + 1.5 * a * b
		fmt.Printf("s = %.3f, a = %.3f b = %.3f, c1 = %.3f, h1 = %.3f\n", s, a, b, c1, h1)
	}
	return s
}

func adjust_line(n_line uint) (uint, uint) {
	var ret uint = 1
	var tmp uint
	for tmp = uint(3); tmp < n_line; tmp = tmp * 2 {
		ret++
	}
	return ret, tmp
}

func main() {
	var n_line uint
	flag.UintVar(&n_line, "l", 3, "num of line")
	flag.Parse()

	n, n_line := adjust_line(n_line)

	area := count_area(n)

	fmt.Printf("n_line = %d, n = %d, area = %.10f\n", n_line, n, area)
}
