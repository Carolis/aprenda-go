package main

import (
	"fmt"
)

type hatsunemiku int

var x hatsunemiku

var y int

func main() {
	fmt.Printf("%v, %T", x, x)
	x = 42
	fmt.Printf("%v, %T", x, x)
	y = int(x)
	fmt.Printf("%v, %T", y, y)
}
