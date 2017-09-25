package main

import (
	"fmt"
	"flag"
)

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
}

