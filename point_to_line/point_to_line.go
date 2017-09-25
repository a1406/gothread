package main

import (
	"fmt"
	"flag"
	"math"
)

//直线(一般式):Ax+By+C=0坐标(Xo，Yo)，，那么这点到这直线的距离就为:(AXo+BYo+C)的绝对值除以根号下(A的平方加上B的平方
//a=a b=-1 c=0
func count_distance(a float64, x1, y1 float64) float64 {
	//	(a * x1 - y1)
	//	math.Sqrt(
	t1 := math.Abs(a * x1 - y1)
	t2 := math.Sqrt(a * a + 1)
	return t1 / t2
}

//line: y = ax
//point: x1, y1
func main() {
	var x1, y1 float64
	var a int
	flag.IntVar(&a, "a", 1, "y = ax")
	flag.Float64Var(&x1, "x", 1, "x1")
	flag.Float64Var(&y1, "y", 1, "y1")	
	flag.Parse()
	fmt.Printf("line: y = %dx\n", a)
	fmt.Printf("point: [%.1f][%.1f]\n", x1, y1)
	dist := count_distance(float64(a), x1, y1)
	fmt.Printf("distance = %.3f\n", dist)
}

