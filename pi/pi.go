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
	c1 := math.Sqrt(1 - 0.5 * 0.5) * 2
	h1 := 0.5
//	var a, b float64
	var s float64 = math.Sqrt(1 - 0.5 * 0.5) * 0.5 * 3
	var n_add float64 = 3
	for i := uint(1); i < n; i++ {
		a := c1
		b := 1 - h1
		c1 = math.Sqrt(0.25*a*a + b*b)
		h1 = math.Sqrt(1 - c1*c1*0.25)
		s = s + 0.5 * a * b * n_add
		n_add = n_add * 2
		fmt.Printf("s = %.10f, a = %.10f b = %.10f, c1 = %.10f, h1 = %.10f, n_add = %.0f\n", s, a, b, c1, h1, n_add)
	}
	return s
}

func count_circumference(n uint) float64 {
	c1 := math.Sqrt(1 - 0.5 * 0.5) * 2
	h1 := 0.5
	var n_add float64 = 3
	for i := uint(1); i < n; i++ {
		a := c1
		b := 1 - h1
		c1 = math.Sqrt(0.25*a*a + b*b)
		h1 = math.Sqrt(1 - c1*c1*0.25)
//		s = s + 0.5 * a * b * n_add
		n_add = n_add * 2
		fmt.Printf("a = %.10f b = %.10f, c1 = %.10f, h1 = %.10f, n_add = %.0f\n", a, b, c1, h1, n_add)
	}
	return c1 * n_add
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
	var n_type uint
	flag.UintVar(&n_line, "l", 3, "num of line")
	flag.UintVar(&n_type, "t", 1, "type: 1 area, 2 circumference")	
	flag.Parse()

	n, n_line := adjust_line(n_line)

	var area float64
	if n_type == 1 {
		area = count_area(n)
	} else {
		area = count_circumference(n)
	}

	fmt.Printf("n_line = %d, n = %d, area = %.10f\n", n_line, n, area)
}
