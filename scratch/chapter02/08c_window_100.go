package main

import "fmt"

// 问题：cap 充足原地扩时，append 的 100 外部能不能看到？

func tweak(s []int) {
	s[0] = 99
	s = append(s, 100) // 100 写进共享数组的 index 2
	s[1] = 98
}

func main() {
	fmt.Println("=== 场景 B: cap 充足(2/4)，append 原地扩 ===")
	s2 := make([]int, 2, 4) // len=2 cap=4
	tweak(s2)

	fmt.Println("s2     =", s2)        // 只显示窗口内 2 个元素
	fmt.Println("s2[:3] =", s2[:3])    // 把窗口拉长到 3 → 能看到 100 吗？
	fmt.Printf("(s2 len=%d, cap=%d)\n", len(s2), cap(s2))

	fmt.Println("\n=== 场景 A: cap 不足(2/2)，append 扩容搬家 ===")
	s1 := []int{1, 2}
	tweak(s1)
	fmt.Println("s1     =", s1)
	fmt.Printf("(s1 len=%d, cap=%d)\n", len(s1), cap(s1))
	// s1[:3] 会 panic —— 外部窗口最长只能到 cap=2，永远够不着新数组
}
