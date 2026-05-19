package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

type cal func(x, y int) int

func call(x, y int, fc cal) int {
	return fc(x, y)
}

func do(o string) cal {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x int, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func main() {
	var o string
	fmt.Scan(&o)
	cb := do(o)
	var a, b int
	fmt.Scan(&a, &b)
	re := cb(a, b)
	fmt.Println(re)
}
