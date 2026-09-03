package main

import "fmt"

// 函数传 slice：扩容 = 断开共享的分水岭
// 场景 A cap 不足 → append 扩容搬家 → 之后写入外部不可见
// 场景 B cap 充足 → append 原地扩 → 之后写入外部仍可见

func main() {
	fmt.Println("=== 场景 A: cap 不足(2/2)，append 扩容搬家 ===")
	s1 := []int{1, 2} // len=2 cap=2
	tweakA(s1)
	fmt.Printf("外部 s1 = %v (len=%d, cap=%d)\n", s1, len(s1), cap(s1))
	fmt.Println("  ↑ s1[0]=99 外部可见；s1[1]=98 外部【看不见】→ 仍是 2")

	fmt.Println("\n=== 场景 B: cap 充足(2/4)，append 原地扩 ===")
	s2 := make([]int, 2, 4) // len=2 cap=4
	tweakB(s2)
	fmt.Printf("外部 s2 = %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))
	fmt.Println("  ↑ s2[0]=99 和 s2[1]=98 外部【都看得见】→ 没搬家还在同一数组")
	fmt.Println("  但外部 len 仍是 2 —— header 是副本，长度改不了外部的")
}

func tweakA(s []int) {
	s[0] = 99          // 共享期：写同一块数组 → 外部可见
	s = append(s, 100) // cap 满(2/2) → 扩容到 4，s 搬家到新数组
	s[1] = 98          // 写的是新数组 → 外部看不见
	fmt.Printf("  tweakA 内部: s=%v (len=%d, cap=%d)\n", s, len(s), cap(s))
}

func tweakB(s []int) {
	s[0] = 99          // 共享期：写同一块数组 → 外部可见
	s = append(s, 100) // cap 够(2/4) → 原地扩，还在同一块数组
	s[1] = 98          // 仍写同一块数组 → 外部看得见
	fmt.Printf("  tweakB 内部: s=%v (len=%d, cap=%d)\n", s, len(s), cap(s))
}
