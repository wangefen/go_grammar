package main

import (
	"fmt"
	"sort"
)

type Point struct {
	X, Y int
}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	numList := []int{6, 3, 1, 34, 12, 2, 4}
	for i := 0; i < len(numList)-1; i++ {
		for j := 0; j < len(numList)-i-1; j++ {
			if numList[j] > numList[j+1] {
				swap(numList, j, j+1)
			}
		}
	}
	fmt.Println(numList)

	num := []int{8, 7, 6, 5, 4, 3, 2, 1}
	sort.Ints(num)
	fmt.Println(num)
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)

}
