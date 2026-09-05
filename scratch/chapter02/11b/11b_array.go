package main

import "fmt"

// Part 10.5 数组补遗：值拷贝 / 长度是类型一部分 / 数组vs slice / 可比较当key

func main() {
	fmt.Println("=== 题1: 长度是类型的一部分 ===")
	a := [3]int{1, 2, 3}
	fmt.Printf("  a 的类型: %T\n", a)
	// b := [4]int{1,2,3,4}
	// a = b              // ✗ cannot use b (type [4]int) as type [3]int

	fmt.Println("\n=== 题2: 数组值拷贝 vs slice 共享 ===")
	arr := [3]int{1, 2, 3}
	cp := arr
	cp[0] = 99
	fmt.Printf("  数组: arr=%v cp=%v（互不影响，整块拷贝）\n", arr, cp)

	sl := []int{1, 2, 3}
	sl2 := sl
	sl2[0] = 99
	fmt.Printf("  slice: sl=%v sl2=%v（header 拷贝，ptr 共享底层）\n", sl, sl2)

	fmt.Println("\n=== 题3: 传参成本对照 ===")
	fmt.Println("  func f(a [100]int): 每次调用 memcpy 整个 100*8 字节")
	fmt.Println("  func f(s []int):    每次调用只拷 24 字节 header")

	fmt.Println("\n=== 题4: 数组可比较、可当 map key；slice 不行 ===")
	if [3]int{1, 2, 3} == [3]int{1, 2, 3} {
		fmt.Println("  数组可以 == 比较")
	}
	seen := map[[2]int]bool{}
	seen[[2]int{3, 4}] = true
	seen[[2]int{3, 4}] = true
	fmt.Printf("  数组当 map key, len=%d\n", len(seen)) // 1（重复 key 覆盖）

	fmt.Println("\n=== 补充: arr[:] 数组→slice, 窗口共享 ===")
	arr2 := [3]int{7, 8, 9}
	s := arr2[:]
	s[0] = 100 // 切出来的 slice 指向 arr2
	fmt.Printf("  arr2=%v s=%v（s 是 arr2 的窗口）\n", arr2, s)
}
