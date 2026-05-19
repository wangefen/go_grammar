package main

import (
	"fmt"
)

func main() {
	namelist := make([]map[string]string, 3, 3)
	if namelist[0] == nil {
		namelist[0] = map[string]string{}
		namelist[0]["wyf"] = "18"
		namelist[0]["www"] = "123"
	}
	if namelist[1] == nil {
		namelist[1] = make(map[string]string)
		namelist[1]["wxf"] = "22"
	}
	if namelist[2] == nil {
		namelist[2] = make(map[string]string)
		namelist[2]["wf"] = "33"
	}
	for i := 0; i < len(namelist); i++ {
		for k, v := range namelist[i] {
			fmt.Println(k, v)
		}
	}
	fmt.Println(namelist)
}
