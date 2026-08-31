package main

import "fmt"

func main() {
	// 三字段
	s := make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))

	// 窗口共享
	a := []int{0, 1, 2, 3, 4, 5}
	b := a[1:4] // 1, 2, 3
	b[0] = 99
	fmt.Println(a) // 0, 99, 2, 3, 4, 5

	// append 扩容断开
	c := []int{1, 2, 3}
	d := c
	d = append(d, 4)
	d[0] = 99
	fmt.Println("扩容后 c 没变:", c)
	fmt.Println("扩容后 d:", d)

	// nil 切片可用
	var n []int
	n = append(n, 1, 2, 3)
	fmt.Println("nil 上 append:", n, len(n), cap(n))

	// 题目1
	a1 := []int{1, 2, 3, 4, 5}
	b1 := a1[1:4]
	b1[0] = 99
	b1[2] = 100
	fmt.Println(a) // 1, 99, 3, 100, 5

	// 题目2
	s1 := make([]int, 2, 5)
	for i := 0; i < 5; i++ {
		s1 = append(s1, 8)
		fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1)) // append第四个元素时，就进行扩容 cap：10
	}

	// 题目3
	x := []int{1, 2, 3}
	y := append(x, 4) // y 扩容，len：4，cap：6，
	y[0] = 100
	fmt.Println(len(y), cap(y), y[0], x[0])

	x2 := make([]int, 3, 6)
	y2 := append(x2, 4)
	fmt.Println(len(y2), cap(y2)) // 4 6
	y2[0] = 11
	fmt.Println(x2[0]) // 11

	// 总结：make([]int, x, y) 会把前x个元素全部初始化为零值，
	// 此时len为x，cap为y，当append超过y - x时，会扩容为2x
}
