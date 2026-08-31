package main

import (
	"fmt"
)

func Add(m map[string]int) {
	m["hello"] = 1
}

func main() {
	m := map[string]int{"a": 1}
	m["b"] = 2

	v := m["不存在"]
	fmt.Println("读不存在：", v)

	v2, ok := m["不存在"]
	fmt.Printf("存在吗? ok=%v v=%d\n", ok, v2)

	delete(m, "a")
	fmt.Println(m)

	m["c"] = 3
	for k, v := range m {
		fmt.Printf("%s=%d ", k, v)
	}
	fmt.Println()

	// 题目1
	m1 := map[string]int{}
	val1 := m1["A"]
	val2 := m1["B"]
	val3 := m1["C"]
	fmt.Printf("val1:%v, val2:%v, val3:%v\n", val1, val2, val3)

	val4, ok := m1["A"]
	fmt.Printf("ok=%v, val4=%v\n", ok, val4)

	val5, ok := m1["A"]
	fmt.Printf("ok=%v, val5=%v\n", ok, val5)

	val6, ok := m1["A"]
	fmt.Printf("ok=%v, val6=%v\n", ok, val6)

	// 题目2
	m2 := map[string]int{}
	for i := 0; i < 20; i++ {
		m2[fmt.Sprintf("%-.2d", i)] = i
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// 题目3
	m3 := map[string]int{}
	Add(m3)
	val7, ok := m3["hello"]
	fmt.Printf("ok=%v, val7=%v\n", ok, val7)

	// 题目4
	var m4 map[string]int
	m4["a"] = 1

}
