package main

import (
	"fmt"
)

type hatsunemiku int

var x hatsunemiku

func main() {
	fmt.Printf("%v, %T", x, x)
	x = 42
	fmt.Printf("%v, %T", x, x)
}
