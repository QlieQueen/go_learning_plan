package main

import "fmt"

// 第二章 Part 7 函数验证：多返回值 / 传值 / 闭包 / defer
// 对应四道题，跑起来核对预测。

func main() {
	fmt.Println("=== 题1: 一切传值 ===")
	s := []int{1, 2}
	tweak(s)
	fmt.Println("  s[0] =", s[0], ", len =", len(s)) // 数组被改，但外部 len 不变？

	m := map[string]int{}
	set(m)
	fmt.Println("  m =", m) // map 内容被改到了？

	fmt.Println("\n=== 题2: defer LIFO + 参数立即求值 ===")
	demo()

	fmt.Println("\n=== 题3: 循环闭包捕获 —— Go 1.22 起每次迭代独立 ===")
	var funcs []func()
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() { fmt.Println(i) })
	}
	for _, f := range funcs {
		f()
	}
	fmt.Println("  （0 1 2 = Go1.22+ 新语义；3 3 3 = 旧版本才有的行为）")

	fmt.Println("\n=== 题4: defer 修改命名返回值 ===")
	fmt.Println("  f() =", f()) // 6？
}

func tweak(s []int) {
	s[0] = 99          // 改底层数组 → 外部可见
	s = append(s, 100) // 改的是 len 副本 → 外部 len 不变
}

func set(m map[string]int) { m["a"] = 1 }

func demo() {
	i := 1
	defer fmt.Println("  defer1:", i)    // 立即求值，存下 i=1
	defer fmt.Println("  defer2:", i+10) // 立即求值，存下 11
	i = 99                               // 改 i 不影响已求值的 defer
	fmt.Println("  i已改成:", i)          // 函数体先执行，再跑 defer
}

func f() (n int) {
	defer func() { n++ }() // 在"交出返回值前"最后改一把 n
	return 5               // 先 n=5，defer 再 n=6，最终返回 6
}
