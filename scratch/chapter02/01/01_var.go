package main

import "fmt"

func main() {
	// 零值
	var i int
	var s string
	var b bool
	var p *int
	fmt.Printf("int=%#v, string=%#v, bool=%#v, ptr=%#v\n", i, s, b, p)

	// 类型后置的好处
	var f func(int) bool
	fmt.Printf("func 零值 nil? %v\n", f == nil)

	// 短声明 + 类型推断
	x := 42
	fmt.Printf("x 的类型是 %T\n", x)

	// 题目1
	var sl []int
	var m map[string]int
	// sl=[]int(nil), m=map[string]int(nil)
	fmt.Printf("sl=%#v, m=%#v\n", sl, m)
}
