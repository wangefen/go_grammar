package main

import (
	"fmt"
	"package_demo/calc"

	"github.com/shopspring/decimal"
)

func main() {
	calc.Speech()
	x := calc.Add(3, 4)
	fmt.Println(x)
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)
}
