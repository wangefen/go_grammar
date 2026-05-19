package main

import "fmt"

func sum(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(sum(a, b, c))
}
