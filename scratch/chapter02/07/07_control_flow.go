package main

import "fmt"

// 第二章 Part 6 控制流验证：if / for / switch
// 对应四个题目，跑起来核对预测。

func main() {
	fmt.Println("=== 题1: switch 顺序匹配，case 可以是表达式 ===")
	x := 10
	switch x {
	case x / 2: // x/2 = 5，不等于 10
		fmt.Println("  case x/2 命中")
	case 10:
		fmt.Println("  case 10 命中")
	default:
		fmt.Println("  都没匹配")
	}

	fmt.Println("\n=== 题2: switch 里的 break 只退 switch，不退循环 ===")
	for i := 0; i < 3; i++ {
		switch i {
		case 1:
			break // 只退出 switch
		}
		fmt.Printf("  %d ", i)
	}
	fmt.Println()

	fmt.Println("\n=== 题3a: for 只有条件 = while 形态 ===")
	n := 0
	for n < 3 {
		n++
	}
	fmt.Println("  while 形态结束，n =", n)

	fmt.Println("\n=== 题3b: 无条件 switch = if-else 链 ===")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("  A")
	case score >= 80:
		fmt.Println("  B")
	default:
		fmt.Println("  C")
	}

	fmt.Println("\n=== 题4: fallthrough 穿到下一个 case，但【不重新求值】它的条件 ===")
	probe := func() int {
		fmt.Println("  [probe] case 2 的条件被求值了！") // 若 fallthrough 会重新求值，这行会出现
		return 2
	}
	switch 1 {
	case 1:
		fmt.Println("  one")
		fallthrough
	case probe():
		fmt.Println("  two")
	}

	fmt.Println("\n=== 补充: label 跳出双层循环 ===")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer
			}
			fmt.Printf("  (%d,%d)", i, j)
		}
	}
	fmt.Println()
}
