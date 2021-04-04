package main

import (
	"fmt"
)

var (
	x int = 420
	y string = "Hatsune Miku"
	z bool = true
)

func main() {
	fmt.Print(x, y, z)
	s := fmt.Sprintf("\n%v\n%v\n%v\n", x,y,z)
	fmt.Print("\nfinalmente", s)
}
