package main

import "fmt"

type shuaibi struct { //父类
	face string
}

func (sb shuaibi) say() {
	fmt.Println("长着一张", sb.face)
}

func (sb shuaibi) do() {
	fmt.Println("我要打十个！")
}

type wyf struct {
	name string
	shuaibi
}

func (w wyf) ssay() {
	fmt.Print(w.name)
}

func main() {
	wxf := wyf{
		name:    "王小凡",
		shuaibi: shuaibi{face: "惨绝人寰的帅脸"},
	}
	wxf.ssay()
	wxf.say()
	wxf.do()
}
