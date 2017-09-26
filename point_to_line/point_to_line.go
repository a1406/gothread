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

func count_pos(a float64, x1, y1, dist float64) (float64, float64) {
	a2 := -1.0 / a
	c := math.Sqrt(a2 * a2 + 1.0)
	rate := dist / c

	var ret_x, ret_y float64
	tmp := x1 * a
	if tmp > y1 {
		ret_y = y1 + rate * math.Abs(a2)
	} else {
		ret_y = y1 - rate * math.Abs(a2)
	}

	tmp = y1 / a
	if tmp > x1 {
		ret_x = x1 + rate
	} else {
		ret_x = x1 - rate		
	}

	return ret_x, ret_y
}

func count_pos2(a float64, x1, y1, dist float64) (float64, float64) {
	var ret_x, ret_y float64
	//y = ax+b  b=y-ax
	a2 := -1.0 / a
//	b := math.Abs(y1 - a2 * x1)
	b := y1 - a2 * x1
	
	angle := math.Atan(a)
	//sin=t1/b   t1=sin*b
	t1 := math.Sin(angle) * b
	ret_x = math.Cos(angle) * t1
	ret_y = math.Sin(angle) * t1
	
	return ret_x, ret_y
}

func count_pos3(a float64, x1, y1, dist float64) (float64, float64) {
	a2 := -1.0 / a
	b := y1 - a2 * x1

	//y=ax y=a2x+b
	//ax-a2x=b  (a-a2)x=b x=b/(a-a2)
	var ret_x, ret_y float64
	ret_x = b / (a - a2)
	ret_y = a * ret_x
	return ret_x, ret_y
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
// 	fmt.Println(math.Atan(1))
// 	fmt.Println(math.Atan(0))
// 	fmt.Println(math.Atan(-1))
	// 	fmt.Println(math.Atan(0.5))
	var ret_x, ret_y float64
	ret_x, ret_y = count_pos(float64(a), x1, y1, dist)
	fmt.Printf("%.2f, %.2f\n", ret_x, ret_y)

	ret_x, ret_y = count_pos2(float64(a), x1, y1, dist)
	fmt.Printf("%.2f, %.2f\n", ret_x, ret_y)

	ret_x, ret_y = count_pos3(float64(a), x1, y1, dist)
	fmt.Printf("%.2f, %.2f\n", ret_x, ret_y)	
	
}

