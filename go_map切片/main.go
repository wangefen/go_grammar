package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["wq"] = "rrr"
	fmt.Println(m["wq"])
	// m3 := map[string]string{
	// 	"name": "Gemini",
	// 	"type": "AI",
	// }
	a := []int{}
	a = append(a, 1)
	fmt.Println(a[0])
}
